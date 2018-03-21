#!/bin/bash

docker-compose up api-build && \
docker-compose up -e REDIS_URL=redis:6379 api
