#!/bin/bash

echo "Export GOPRIVATE"
export GOPRIVATE=github.com/es-hs/*

echo "Fetch latest osmodels"
go get -u github.com/es-hs/osmodels@main

echo "Clean go mod"
go mod tidy

echo "Update go vendor"
go mod vendor
