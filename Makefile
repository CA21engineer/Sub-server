# 起動
up:
	sh ./run-local.sh

# 停止
down:
	docker-compose down

# docker内のサーバにssh
app:
	docker exec -it subs-server_api-server /bin/bash

# docker内のmysqlにssh
mysql:
	docker exec -it subs-server_mysql_1 /bin/bash

# docker内のサーバlogを出力
logs:
	docker logs -f subs-server_api-server_1
