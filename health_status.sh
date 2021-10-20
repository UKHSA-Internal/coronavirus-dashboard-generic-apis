#!/bin/sh

if [ -f "/opt/healthcheck/healthy.txt" ]; then
  echo "Service healthy"
  exit 0
else
  echo "Service unhealthy"
  exit 1
fi;
