up: dev-env dev-air ## Startup / Spinup Docker Compose and air
down: docker-stop  ## Stop Docker

dev-env: 
	@ docker-compose -f docker-compose-pg-only.yml up -d

docker-stop:
	@ docker-compose -f docker-compose-pg-only.yml down

dev-air: $(AIR) ## Starts AIR ( Continuous Development app).
	air

upd-cov: 
	@ go test ./... -coverprofile coverage.out