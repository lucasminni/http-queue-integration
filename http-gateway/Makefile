BUILD_CMD = docker build -t http-gateway-app .
START_CMD = docker run -d --rm -p 8080:8080 --name http-gateway-app http-gateway-app 
STOP_CMD = docker stop http-gateway-app
REMOVE_CMD = docker rm http-gateway-app


build:
	-$(REMOVE_CMD)
	$(BUILD_CMD)

start:
	$(START_CMD)
	docker ps -a

rebuild:
	-$(STOP_CMD)
	-$(REMOVE_CMD)
	$(BUILD_CMD)
	$(START_CMD)
	docker ps -a

localstack:
	@echo "Iniciando Localstack e configurando AWS CLI"
	localstack start -d
	-aws configure --profile minni perfilkey: teste secret: teste token: secret region: sa-east-1 format: json
	aws configure list
	@echo Criando tabela orders no DynamoDB Local
	aws --no-cli-pager --endpoint-url=https://localhost.localstack.cloud:4566 --profile minni dynamodb create-table --table-name orders --attribute-definitions AttributeName=id,AttributeType=S --key-schema AttributeName=id,KeyType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5
