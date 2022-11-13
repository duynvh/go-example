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
	-v /home/duynvh/Documents/code/go/go-example/.docker:/bitnami \
	bitnami/mysql:8.0

runnats:
	@docker run -d --name nats --network ${DOCKER_NETWORK} --rm -p 4222:4222 -p 8222:8222 nats --http_port 8222

runredis:
	@docker run -d --name redis --network ${DOCKER_NETWORK} --rm -p 6379:6379  redis

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

runapp:
	@docker run -d \
	--network ${DOCKER_NETWORK} \
	--name go-example \
	--privileged=true \
	-p 3000:3000 \
	-e MYSQL_GORM_DB_URI=${MYSQL_GORM_DB_URI} \
	-e MYSQL_GORM_DB_TYPE=${MYSQL_GORM_DB_TYPE} \
	-e SECRET=${SECRET} \
	-e USER_SERVICE_URL=${USER_SERVICE_URL} \
	nguyenvohoangduy/go-example

bufgenerate:
	buf generate

.PHONY: rundb startdb migrateup start runapp runnats runredis bufgenerate