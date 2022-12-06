#!/bin/bash
set -e

export XLSX_MODEL=make  ## export || make
echo "-------------------------- begin ${XLSX_MODEL} xlsx data  --------------------------------"

go run ./src/cmd/main.go
