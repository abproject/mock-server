#!/bin/bash

DIR="`dirname \"$0\"`"
# Main
markdown-pp $DIR/src/main.md -o $DIR/../README.md
# Rest
## API
markdown-pp $DIR/src/rest/rest-api.md -o $DIR/REST-API.md
## YAML
markdown-pp $DIR/src/rest/rest-yaml.md -o $DIR/REST-YAML.md
## YAML
markdown-pp $DIR/src/rest/rest-json.md -o $DIR/REST-JSON.md