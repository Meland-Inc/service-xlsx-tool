#!/bin/bash
set -e


# git checkout .
# git checkout main
# git pull
# git submodule update --init --recursive


export XLSX_MODEL=export  ## export || make
echo "-------------------------- begin ${XLSX_MODEL} xlsx data  --------------------------------"

echo "------------  set xlsx tool environment variable --------------------------------"
## ---------------melandConfig----------------
export MELAND_CONFIG_DB_HOST=192.168.50.15
export MELAND_CONFIG_DB_USER=root
export MELAND_CONFIG_DB_PASS=root
export MELAND_CONFIG_DB_PORT=3306
export MELAND_CONFIG_DB_DATABASE=meland_cnf_dev


## ---------------xlsx data dir-------------------
export XLSX_DIR= ##your xlsx dir sdf


echo "---------------------------  export xlsx data  --------------------------------"
go run ./src/cmd/main.go


echo "---------------------------  export CSV data  --------------------------------"
XLSX_TO_CSV_TOOL=xlsxToCsv_mac_arm64          ## mac m1  
# XLSX_TO_CSV_TOOL=xlsxToCsv_mac_x86          ## mac x86
# XLSX_TO_CSV_TOOL=xlsxToCsv_windows.exe      ## windows
echo "xlsx to csv file: ${XLSX_DIR}/${XLSX_TO_CSV_TOOL}"
${XLSX_DIR}/${XLSX_TO_CSV_TOOL}