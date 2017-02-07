#!/bin/bash
set -e

cd $GINPATH

/go/bin/gin --appPort $PORT -port 3000 --path $GINPATH run
