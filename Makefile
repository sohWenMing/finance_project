.PHONY: test-print
test-print: 
	@echo "hello from the make file!"


################################ For Server Image ########################################
.PHONY: build-server-image
build-server-image:
	touch go.sum
	docker build -t finance-server:latest -f cmd/server_cmd/Dockerfile .
#Usage:  docker buildx build [OPTIONS] PATH | URL | -
# in this example: -t and -f are the options and . takes up the part that is necesary for <PATH | URL | -> 
# <PATH | URL | - > is for the build context, so it can be either a path, a url like a github repo or something from stdin

.PHONY: run-server-image
run-server-image:
	@docker stop finance-server 2>/dev/null || true
	@make build-server-image
	@docker run --network finance_server_network -p 8080:8080 --name finance-server --rm  finance-server
# runs image not in detached mode, so that the server can be up for manual testing

.PHONY: run-server-image-background
run-server-image-background:
	@docker stop finance-server 2>/dev/null || true
	@make build-server-image
	@docker run -d --network finance_server_network -p 8080:8080 --name finance-server --rm  finance-server
# runs server in detached mode, so that for testing purposes the container can be running and present and automated
# tests can run against it

# Usage:  docker run [OPTIONS] IMAGE [COMMAND] [ARG...]
# when running a container , options all have to come before the final IMAGE which should be in the repository

.PHONY: stop-server-image
stop-server-image:
	@docker stop finance-server 2>/dev/null || true



############################# For PostgresDB Image ##########################################
.PHONY: start-postgres-image
start-postgres-image: 
	@docker run --name finance_postgres --network finance_server_network \
	-e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=finance_db \
	-v finance_db_data:/var/lib/postgresql/data \
	-d postgres:15.12  

.PHONY: stop-postgres-image
stop-postgres-image:
	@docker stop finance_postgres
	@docker container prune -f