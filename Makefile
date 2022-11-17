DB_USER="user"
DB_PASS="capstone"
DB_HOST="34.69.193.163"
DB_PORT="3306"
DB_NAME="staging"

DSN="${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?charset=utf8&parseTime=True&loc=Local"

migrateup:
	goose -dir migration mysql ${DSN} up 

migratedown:
	goose -dir migration mysql ${DSN} down
	
.PHONY: migrateup migratedown