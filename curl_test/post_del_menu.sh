#!/bin/bash


# curl --location --request POST 'http://localhost:8080/admin/v01/menu/delete' \
# --header 'Content-Type: application/json' \
# --data-raw '{   
#     "menuid" : 1,
#     "imageurl":"https://dimg04.c-ctrip.com/images/0M71x120009orv5n7598E_Q60.jpg_.webp",
#     "name":"MeatSashimi",
#     "quantity" : 100,
#     "price":10000,
#     "recommendation":true ,
#     "admin" : "yongari"
# }'
curl --location --request POST 'http://localhost:8080/admin/v01/admin/menu/delete' \
--header 'Content-Type: application/json' \
--data-raw '{   
    "menuid" : 3,
    "imageurl":"https://dimg04.c-ctrip.com/images/0M71x120009orv5n7598E_Q60.jpg_.webp",
    "name":"MeatSashimi",
    "quantity" : 100,
    "price":10000,
    "recommendation":true ,
    "admin" : "yongari", 
    "score" : 1, 
    "review" : "맛있는 육회 최고!"
}'

curl --location --request POST 'http://localhost:8080/admin/v01/admin/menu/delete' \
--header 'Content-Type: application/json' \
--data-raw '{   
    "menuid" : 3,
}'