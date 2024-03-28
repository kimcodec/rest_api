#!/bin/sh
echo "Waiting db to init...";
sleep 5;
./bin/goose -dir ./internal/migrations/ postgres "host=postgres user=postgres database=rest_api_db password=postgres sslmode=disable" up;
/val;