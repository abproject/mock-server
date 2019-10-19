#!/bin/bash

DIR="`dirname \"$0\"`"
# Main
markdown-pp $DIR/src/main.md -o $DIR/../README.md

# Examples
## Files
markdown-pp $DIR/src/examples/files/files.md -o $DIR/FILES_EXAMPLE.md
## CRUD
markdown-pp $DIR/src/examples/crud/crud.md -o $DIR/CRUD_EXAMPLE.md
## Entities
markdown-pp $DIR/src/examples/entities/entities.md -o $DIR/ENTITIES_EXAMPLE.md
