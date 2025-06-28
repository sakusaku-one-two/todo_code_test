#bin/bash


docker-compose up --build
# docker-compose exec api_server cd /usr/src/app || go run cmd/migrate/up/main.go
docker-compose down
echo DONE!!