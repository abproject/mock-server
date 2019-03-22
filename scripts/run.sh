#!/bin/bash

docker run -p 8000:8000 -v ${PWD}/example:/example mock-server -file=/example/file/config.yml
