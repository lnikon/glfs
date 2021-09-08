#!/bin/bash

if [[ -z ${GL_PORT} ]]; then
  GL_PORT=8090
else
  GL_PORT="${GL_PORT}"
fi


if [[ -z ${GL_HOST} ]]; then
  GL_HOST=localhost
else
  GL_HOST="${GL_HOST}"
fi

echo "Testing endpoints with port=${GL_HOST}"

curl -XPOST -d'{}' ${GL_HOST}:${GL_PORT}/algorithm 
curl -XGET -d'{}' ${GL_HOST}:${GL_PORT}/computation 
curl -XPOST -d'{}' ${GL_HOST}:${GL_PORT}/computation 
