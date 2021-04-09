#!/bin/bash

PLATFORMS_FILE_ARGUMENT=$1
PLATFORMS_FILE="${PLATFORMS_FILE_ARGUMENT:-platforms.txt}"
WORKING_DIR="${PWD}"
FULL_PATH_PLATFORMS_FILE=$WORKING_DIR/$PLATFORMS_FILE
APP_NAME=go-ddns

# check file platforms.txt
if [ ! -f $FULL_PATH_PLATFORMS_FILE ]; then
    echo "Platforms file: $FULL_PATH_PLATFORMS_FILE does not exist"
    exit 1
fi

# remove build folder
echo "Clean build folder"
rm -rf ./build

# create build folder
echo "Create build folder"
mkdir -p build

# read platforms file
while read line; do
    temp=(${line//\// })

    GOOS=${temp[0]}
    GOARCH=${temp[1]}

    output_name='go-ddns'

    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    mkdir -p build/$GOOS/$GOARCH
    output_path="$WORKING_DIR/build/$GOOS/$GOARCH/$output_name"

    echo "Building for OS=$GOOS Architecture=$GOARCH"
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_path
    if [ ! -f $output_path ]; then
        echo "Failed when build for OS=$GOOS Architecture=$GOARCH"
        exit 1
    else
        echo "Done with output file: $output_path"
    fi

    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done < $PLATFORMS_FILE

echo "DONE"