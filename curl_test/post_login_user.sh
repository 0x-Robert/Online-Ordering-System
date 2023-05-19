#!/bin/bash

curl --location --request POST 'http://localhost:8080/admin/v01/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":"yongari12",
    "password":"12354"
}'
