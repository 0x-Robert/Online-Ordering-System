#!/bin/bash

curl --location --request POST 'http://localhost:8080/v01/admin/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":"yongari12",
    "password":"12354"
}'
