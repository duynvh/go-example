include .env

createnetwork: 
	@docker network create ${DOCKER_NETWORK}

rundb:
	@docker run -d \
	--network ${DOCKER_NETWORK} \
	--name mysql \
	--privileged=true \
	-p 3306:3306 \
	-e MYSQL_ROOT_PASSWORD=${DB_ROOT_PASSWORD} \
	-e MYSQL_USER=${DB_USER} \
	-e MYSQL_PASSWORD=${DB_PASSWORD} \
	-e MYSQL_DATABASE=${DB_NAME} \
	-v /home/pcbackend04/Documents/code/go/golang-scalable-backend/docker:/bitnami \
	bitnami/mysql:8.0

startdb:
	@docker start mysql

migrateup:
	@docker build -t migrator ./migrator && \
	docker rm -f migrator && \
	docker run \
	--name migrator \
	--network ${DOCKER_NETWORK} \
	migrator \
	-path="/migrations/" \
	-database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(mysql:3306)/${DB_NAME}?charset=utf8mb4&parseTime=True&loc=Local" \
	up

.PHONY: rundb startdb migrateup start