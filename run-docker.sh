#!/bin/bash

make build
docker build --no-cache -t guestbook .
docker run --rm -p 8080:8080 guestbook