#!/bin/bash

echo "Export GOPRIVATE"
export GOPRIVATE=github.com/es-hs/*

echo "Fetch latest authzclient"
go get -u github.com/es-hs/authzclient@main

echo "Clean go mod"
go mod tidy

echo "Update go vendor"
go mod vendor
