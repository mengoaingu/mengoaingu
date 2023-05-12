#!/bin/sh

export $(cat .env) > /dev/null 2>&1;

docker stack deploy -c docker-stack.yaml mengoaingu