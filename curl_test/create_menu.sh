#!/bin/bash

curl --location --request POST 'http://localhost:8080/admin/v01/menu/create' \
--header 'Content-Type: application/json' \
--data-raw '{   
    "menu_id" : 1,
    "image_url":"https://dimg04.c-ctrip.com/images/0M71x120009orv5n7598E_Q60.jpg_.webp",
    "name":"MeatSashimi",
    "quantity" : 100,
    "price":10000,
    "recommendation":true ,
    "admin" : "yongari"
}'

