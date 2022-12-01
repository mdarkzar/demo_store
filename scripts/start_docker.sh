#!/usr/bin/env bash

ROOT=..

source variables.sh
cd $ROOT/docker
docker compose build
docker compose -p store up --force-recreate --remove-orphans