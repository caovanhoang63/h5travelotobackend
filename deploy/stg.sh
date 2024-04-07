
#!/usr/bin/env bash

APP_NAME=h5traveloto

docker load -i ${APP_NAME}.tar
docker rm -f ${APP_NAME}

docker run -d --name ${APP_NAME} \
  --network h5traveloto \
  -e VIRTUAL_HOST="api.h5traveloto.site" \
  -e LETSENCRYPT_HOST="api.h5traveloto.site" \
  -e LETSENCRYPT_EMAIL="h5traveloto@gmail.com" \
  -p 8080:8080 \
  ${APP_NAME}