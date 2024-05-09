#!/usr/bin/env bash
docker compose down


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

sleep 5
echo 'starting to copy jars'
docker cp kafka-connect:/usr/share/confluent-hub-components ./data/connect-jars
echo 'done copying jars'

echo 'installed plugins'
curl -s localhost:8083/connector-plugins|jq '.[].class'|egrep 'MySqlConnector|ElasticsearchSinkConnector'

echo 'waiting for control-center start...'

# Add permissions to the certs folder
chmod 777 certs
chmod 777 certs/es01
chmod 777 certs/es01/keystore.jks
chmod 777 certs/es01/truststore.jks

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

# map the location field to geo_point

curl --silent --show-error -k -XPUT -u elastic:oc2nq0mhv8bju1e -H 'Content-Type: application/json' \
    https://localhost:9200/_index_template/rmoff_template01/ \
    -d'{
        "index_patterns": [ "hotels*" ],
        "template": {
            "mappings": {
                "properties": {
                    "location_geo_point": {
                        "type": "geo_point"
                }
            }
        }
    }}'



#keytool -import -alias elasticsearch -file es01.crt -keystore truststore.jks
#keytool -import -alias elasticsearch -file es01.crt -keystore keystore.jks

# Start Elasticsearch connector
curl -i -X POST -H "Accept:application/json" -H  "Content-Type:application/json" http://localhost:8083/connectors/ -d @es-sink-enriched.conf.json
curl -i -X POST -H "Accept:application/json" -H  "Content-Type:application/json" http://localhost:8083/connectors/ -d @es-sink-suggest.conf.json


# Check the status of the connector
curl localhost:8083/connectors/elastic-sink-enriched/status/




