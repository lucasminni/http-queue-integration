build:
	-docker rm http-gateway-app
	docker build -t http-gateway-app .

start:
	-docker network create gateway-queue-network
	docker run -d --rm -p 8080:8080 --name http-gateway-app --network gateway-queue-network http-gateway-app 
	docker ps -a

rebuild:
	-docker stop http-gateway-app
	-docker rm http-gateway-app
	docker build -t http-gateway-app .
	docker run -d --rm -p 8080:8080 --name http-gateway-app --network gateway-queue-network http-gateway-app 
	docker ps -a

aws:
	-docker stop localstack
	-docker rm localstack
	docker run -d --rm --name localstack --network gateway-queue-network -p 4566:4566 -e SERVICES=dynamodb localstack/localstack
	-aws configure --profile minni perfilkey: teste secret: teste token: secret region: sa-east-1 format: json
	aws configure list
	aws --no-cli-pager --endpoint-url=https://localhost.localstack.cloud:4566 --profile minni dynamodb create-table --table-name orders --attribute-definitions AttributeName=id,AttributeType=S --key-schema AttributeName=id,KeyType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5
