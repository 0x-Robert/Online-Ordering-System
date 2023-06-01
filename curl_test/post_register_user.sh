#!/bin/bash


curl --location --request POST 'http://localhost:8080/v01/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":"yongari",
    "password":"1234"
}'
