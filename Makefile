up:
	sh ./run-local.sh

down:
	docker-compose down

build:
	docker-compose up --build

app:
	docker exec -it subs-server_api-server /bin/bash

mysql:
	docker exec -it subs-server_mysql_1 /bin/bash
