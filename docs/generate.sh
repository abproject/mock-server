#!/bin/bash

DIR="`dirname \"$0\"`"
# Main
markdown-pp $DIR/src/main.md -o $DIR/../README.md