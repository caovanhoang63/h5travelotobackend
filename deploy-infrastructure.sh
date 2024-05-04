#!/usr/bin/env bash

DEPLOY_CONNECT1=root@157.245.201.95
DEPLOY_CONNECT2=root@206.189.83.53


scp -r -o StrictHostKeyChecking=no ./docker ${DEPLOY_CONNECT2}:~
