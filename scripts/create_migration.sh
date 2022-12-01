#!/bin/bash

ROOT=../docker

if [ -z "$1" ]
  then
    echo "Укажите название миграции"
    exit 1
fi

NAME=$1

migrate create -ext sql -dir $ROOT/migrate -tz=Asia/Tashkent -format=020120061504  $NAME