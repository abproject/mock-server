#!/bin/bash

docker run -p 8000:8000 -v ${PWD}/example:/example abezpalov/mock-server:latest -file=/example/crud/config.yml