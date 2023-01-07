package misc

import (
	"fmt"
	"log"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbName = "test"
	passwd = "test"
)

func SetupGormWithDocker() (*gorm.DB, func()) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	runDockerOpt := &dockertest.RunOptions{
		Repository: "postgres",    // image
		Tag:        "14.1-alpine", // version
		Env:        []string{"POSTGRES_PASSWORD=" + passwd, "POSTGRES_DB=" + dbName},
	}

	fnConfig := func(config *docker.HostConfig) {
		config.AutoRemove = true                     // set AutoRemove to true so that stopped container goes away by itself
		config.RestartPolicy = docker.NeverRestart() // don't restart container
	}

	resource, err := pool.RunWithOptions(runDockerOpt, fnConfig)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// call clean up function to release resource
	fnCleanup := func() {
		err := resource.Close()
		if err != nil {
			log.Fatalf("Could not close resource: %s", err)
		}
	}

	conStr := fmt.Sprintf("host=localhost port=%s user=postgres dbname=%s password=%s sslmode=disable",
		resource.GetPort("5432/tcp"), // get port of localhost
		dbName,
		passwd,
	)

	fmt.Println("conStr:", conStr)
	var gdb *gorm.DB
	// retry until db server is ready
	err = pool.Retry(func() error {
		gdb, err = gorm.Open(postgres.Open(conStr))
		if err != nil {
			fmt.Println("gorm open error:", err)
			return err
		}
		db, err := gdb.DB()
		if err != nil {
			fmt.Println("gorm db error:", err)
			return err
		}
		return db.Ping()
	})
	chk(err)

	// container is ready, return *gorm.Db for testing
	return gdb, fnCleanup
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
