#!/usr/bin/env bash

PB_DIR=$1
OUT_DIR=$2

if [ ! -d $OUT_DIR ]; then
    mkdir $OUT_DIR
fi

for dir in $(ls $PB_DIR); do
	protoc -I$PB_DIR --go_out=plugins=onerpc:$OUT_DIR $PB_DIR/$dir/*.proto
done