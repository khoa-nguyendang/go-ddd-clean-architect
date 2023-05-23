build:
	docker-compose build

up:
	docker-compose up -d
down:
	docker-compose down

prepare:
	sudo mkdir -p ./_volumes/mysqldb
	sudo mkdir -p ./_volumes/tidb
	sudo mkdir -p ./_volumes/zookeeper
	sudo mkdir -p ./_volumes/kafka
	sudo mkdir -p ./_volumes/elasticsearch
	sudo mkdir -p ./_volumes/logstash/pipeline
	sudo mkdir -p ./_volumes/logstash/config/queries

	sudo chown 999:999 ./_volumes/mysqldb
	sudo chown 999:999 ./_volumes/tidb
	sudo chown 999:999 ./_volumes/zookeeper
	sudo chown 999:999 ./_volumes/kafka
	sudo chown 999:999 ./_volumes/elasticsearch
	sudo chown 999:999 ./_volumes/logstash
	sudo chown 999:999 ./docker/mysqlseed
	
	sudo chmod -R 777 ./_volumes/mysqldb
	sudo chmod -R 777 ./_volumes/tidb
	sudo chmod -R 777 ./_volumes/zookeeper
	sudo chmod -R 777 ./_volumes/kafka
	sudo chmod -R 777 ./_volumes/elasticsearch
	sudo chmod -R 777 ./_volumes/
	sudo chmod -R 777 ./docker/*