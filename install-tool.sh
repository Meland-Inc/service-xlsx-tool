#!/bin/bash
set -e

echo "------------------ install golang version manager 'gvm' --------------"
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

echo "------------------ install golang version 1.17.6 --------------"
gvm install go1.17.6
gvm diff go1.17.6
go version
go env -w GO111MODULE=auto
go env -w GOPROXY=https://goproxy.cn,direct 