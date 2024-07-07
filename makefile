.PHONY: write_db_link

DB_NAME = your_database_name
PG_USER = your_postgres_username
PG_PASSWORD = your_postgres_password
PG_HOST = your_postgres_host
PG_PORT = your_postgres_port

DB_CONNECTION_STRING = postgres://$(PG_USER):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(DB_NAME)?sslmode=disable

write_db_link:
	echo $(DB_CONNECTION_STRING) > db_link.txt