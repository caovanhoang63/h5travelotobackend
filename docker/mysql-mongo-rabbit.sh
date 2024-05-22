#!/usr/bin/env bash

DEPLOY_CONNECT=root@157.245.201.95

scp -r -o StrictHostKeyChecking=no ./mysql-mongo-rabbit ${DEPLOY_CONNECT}:~