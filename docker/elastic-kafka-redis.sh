#!/usr/bin/env bash

DEPLOY_CONNECT=root@159.89.200.57

scp -r -o StrictHostKeyChecking=no ./elastic-kafka-redis ${DEPLOY_CONNECT}:~


