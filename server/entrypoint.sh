#!/bin/sh

set -e

whoami
ls -las /opt/app
exec /opt/app/generic_api
