#!/bin/bash

set -e 
set -o pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Create build directory
mkdir -p $DIR/build/static

# Copy schema to build directory
#cp $DIR/schema.sql $DIR/build

echo "Compile UI"
pushd $DIR
	npm install
	./node_modules/.bin/gulp browserify
popd

go-bindata -prefix $DIR/build/static -pkg bindata -o $DIR/server/bindata/bindata.go $DIR/build/static

echo "Compile executable"
pushd $DIR/build
	govendor build ../server/mtti-board
popd

cp $DIR/client/index.html $DIR/build/static
