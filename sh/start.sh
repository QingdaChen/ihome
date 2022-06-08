#!/bin/sh

if [ "$1" = "web" ] ; then
   cd ../web
   go run main.go
elif [ "$1" = "captcha" ] ; then
   cd ../service/captcha
   go run main.go handler.go
elif [ "$1" = "house" ] ; then
    cd ../service/house
    go run main.go handler.go
elif [ "$1" = "user" ] ; then
    cd ../service/user
    go run main.go handler.go
fi