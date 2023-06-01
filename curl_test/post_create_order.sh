#!/bin/bash


curl --location --request POST 'http://localhost:8080/v01/user/order' \
--header 'Content-Type: application/json' \
--data-raw '{
    "menuname" :"육회사시미",
    "customer" :"용가리",
    "phonenumber":"010-1234-5678",
    "address" : "인천시 북항 ",
    "quantity" : 4,
    "paymentinformation":"현금"
}'

