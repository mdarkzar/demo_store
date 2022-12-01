#!/bin/bash

ROOT=../docker
export DB_NAME=store
export DB_PASSWORD=store

migrate -path $ROOT/migrate -database postgres://$DB_NAME:$DB_PASSWORD@localhost:5432/$DB_NAME?sslmode=disable down 1