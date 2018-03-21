#!/bin/bash

docker-compose up -d test-redis

docker-compose run --rm \
        -e REDIS_URL=test-redis:6379 \
        api-build \
        go test

docker-compose rm test-redis
