#!/usr/bin/env bash

ROOT=../backend

mockgen -destination=$ROOT/internal/repository/mocks.go -package=repository -source=$ROOT/internal/repository/interface.go
mockgen -destination=$ROOT/internal/bridge/mocks.go -package=bridge -source=$ROOT/internal/bridge/interface.go
mockgen -destination=$ROOT/internal/transaction/mocks.go -package=transaction -source=$ROOT/internal/transaction/interface.go
