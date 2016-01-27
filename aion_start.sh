#!/usr/bin/env bash

RUN_AION='aion -db-user aion -db-pass aion -db-host db -db-port 3306 -db-name aion -nsq-host nsqd'

command -v docker-compose &>/dev/null \
    && RUN_AION="docker-compose run --service-ports aion ${RUN_AION}"

eval "${RUN_AION} $@"
