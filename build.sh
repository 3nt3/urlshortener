#!/bin/sh

docker build -t 'urlshortener-api' .
docker-compose up $1