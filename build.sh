#!/usr/bin/env bash


export APP_NAME="IMDB-Rating"
export PORT=8080
export ATLAS_URI="mongodb://user:pass@127.0.0.1:27017/?authMechanism=SCRAM-SHA-256"

# source dev.env

go install
if [ $? != 0 ]; then
  echo "## Build Failed ##"
  exit
fi


echo "Doing some cleaning ..."
go clean
echo "Done."

echo "Running goimport ..."
goimports -w=true .
echo "Done."

echo "Running go format ..."
gofmt -w .
echo "Done."

echo "Running go build ..."
go build -race
if [ $? != 0 ]; then
  echo "## Build Failed ##"
  exit
fi
echo "Done."

echo "Running unit test ..."
# go test -p=1 ./services/... ./utils/...
if [ $? == 0 ]; then
    echo "Done."
	echo "## Starting service ##"
    ./movie-suggestions-api
fi

./movie-suggestions-api
