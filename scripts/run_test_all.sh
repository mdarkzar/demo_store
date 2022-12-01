#!/bin/bash
source variables.sh
export DEBUG=true
export ROOT=../backend

go clean -testcache
cd $ROOT/internal/
echo '-- тест usecase --'
export CONF_PATH=../../../../config/conf.yaml
for s in $(go list ./usecase/test/...); do if ! go test -failfast -p 1 $s; then break; fi; done 2>&1 
echo '-- тест repository --'
export CONF_PATH=../../../../../config/conf.yaml
for s in $(go list ./repository/postgresql/test/...); do if ! go test -failfast -p 1 $s; then break; fi; done 2>&1 