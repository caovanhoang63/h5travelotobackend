#!/usr/bin/env bash

docker compose down

echo "Compose Down Done"


docker compose up -d

echo "Compose Up Done"

echo "Waiting for Kafka Connect to start"


bash -c ' \
echo -e "\n\n=============\nWaiting for Kafka Connect to start listening on localhost ⏳\n=============\n"
while [ $(curl -s -o /dev/null -w %{http_code} http://localhost:8083/connectors) -ne 200 ] ; do
  sleep 5
done
echo -e $(date) "\n\n--------------\n\o/ Kafka Connect is ready! Listener HTTP state: " $(curl -s -o /dev/null -w %{http_code} http://localhost:8083/connectors) "\n--------------\n"
'

echo 'starting to copy jars'
docker cp kafka-connect:/usr/share/confluent-hub-components ./data/connect-jars
echo 'done copying jars'

echo 'installed plugins'
curl -s localhost:8083/connector-plugins|jq '.[].class'|egrep 'MySqlConnector|ElasticsearchSinkConnector'

echo 'waiting for control-center start...'



# Start MySQL connector
curl -i -X POST -H "Accept:application/json" -H  "Content-Type:application/json" http://localhost:8083/connectors/ -d @source.conf.json

echo '-------'

sleep 2
# Check the status of the connector
curl localhost:8083/connectors/mysql-connector/status/

echo 'waiting to run ksqldb...'

sleep 5

cat ./ksqldb-ddl.sql | docker exec -i ksqldb ksql http://localhost:8088


echo 'Done running ksqldb-ddl.sql'

sleep 2

echo '-------'

# Start Elasticsearch connector
curl -i -X POST -H "Accept:application/json" -H  "Content-Type:application/json" http://localhost:8083/connectors/ -d @es-sink-enriched.conf.json


# Check the status of the connector
curl localhost:8083/connectors/elastic-sink-enriched/status/
