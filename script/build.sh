#!/usr/bin/env bash

set -euo pipefail

VERSION=${3:-$(go run main.go --version | awk '{print $2}')}
export GOOS=${1:-darwin}
export GOARCH=${2:-amd64}
ARCHIVE_PATH="../dist/${GOOS}-${GOARCH}-v${VERSION}.zip"

echo "Building for ${GOOS}/${GOARCH}"
go build -o tmp/yamlc
cd tmp
[[ -f $ARCHIVE_PATH ]] && rm -f $ARCHIVE_PATH
zip $ARCHIVE_PATH yamlc

rm -f yamlc

exit
