#!/bin/bash

export DB_NAME=store
export DB_PASSWORD=store

migrate -path ../docker/migrate -database postgres://$DB_NAME:$DB_PASSWORD@localhost:5432/$DB_NAME?sslmode=disable up 1