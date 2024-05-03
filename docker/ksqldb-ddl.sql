SET 'auto.offset.reset' = 'earliest';

CREATE STREAM HOTELS WITH(KAFKA_TOPIC='mysql-debezium-h5traveloto.h5traveloto.hotels', PARTITIONS=1, VALUE_FORMAT='AVRO');
CREATE STREAM HOTEL_FACILITIES WITH(KAFKA_TOPIC='mysql-debezium-h5traveloto.h5traveloto.hotel_facility_details', PARTITIONS=1, VALUE_FORMAT='AVRO');
CREATE STREAM ROOM_FACILITIES WITH(KAFKA_TOPIC='mysql-debezium-h5traveloto.h5traveloto.room_facility_details', PARTITIONS=1 ,VALUE_FORMAT='AVRO');
CREATE STREAM ROOM_TYPES WITH (KAFKA_TOPIC='mysql-debezium-h5traveloto.h5traveloto.room_types', PARTITIONS=1 ,VALUE_FORMAT='AVRO');

CREATE TABLE PROVINCES (ID VARCHAR PRIMARY KEY)
    WITH (KAFKA_TOPIC='mysql-debezium-h5traveloto.h5traveloto.provinces',
        VALUE_FORMAT='AVRO');

CREATE TABLE DISTRICTS (ID VARCHAR PRIMARY KEY)
    WITH (KAFKA_TOPIC='mysql-debezium-h5traveloto.h5traveloto.districts',
        VALUE_FORMAT='AVRO');

CREATE TABLE WARDS (ID VARCHAR PRIMARY KEY)
    WITH (KAFKA_TOPIC='mysql-debezium-h5traveloto.h5traveloto.wards',
        VALUE_FORMAT='AVRO');

CREATE TABLE  HOTEL_FACILITY_LIST  AS
	SELECT HF.HOTEL_ID, COLLECT_LIST(FACILITY_ID) AS FACILITY_LIST
    FROM  HOTEL_FACILITIES HF
    GROUP BY HF.HOTEL_ID;

CREATE TABLE  ROOM_FACILITY_LIST  AS
    SELECT RF.ROOM_ID, COLLECT_LIST(FACILITY_ID) AS FACILITY_LIST
    FROM  ROOM_FACILITIES RF
    GROUP BY RF.ROOM_ID;

CREATE STREAM ROOM_TYPES_ENRICHED AS
    SELECT
        RT.ID ID,
        RT.HOTEL_ID,
        RT.NAME,
        RT.MAX_CUSTOMER,
        RT.AREA,
        RT.BED,
        RT.PRICE,
        RT.STATUS,
        RT.TOTAL_ROOM,
        RT.PAY_IN_HOTEL,
        RT.BREAK_FAST,
        RT.FREE_CANCEL,
        RT.CREATED_AT,
        RT.UPDATED_AT,
        RF.FACILITY_LIST
    FROM ROOM_TYPES RT
        LEFT JOIN ROOM_FACILITY_LIST RF ON RT.ID = RF.ROOM_ID
PARTITION BY ID;

CREATE STREAM HOTELS_ENRICHED
AS
    SELECT
        H.ID AS ID,
        H.OWNER_ID,
        H.NAME AS NAME,
        H.HOTEL_TYPE,
        H.LOGO,
        H.IMAGES,
        H.ADDRESS,
        P.NAME AS PROVINCE,
        D.NAME AS DISTRICT,
        W.NAME AS WARD,
        H.LAT,
        H.LNG,
        H.STAR,
        H.STATUS,
        H.CREATED_AT,
        H.UPDATED_AT,
        HF.FACILITY_LIST
    FROM HOTELS H
        LEFT JOIN PROVINCES P ON 'Struct{code='+H.PROVINCE_CODE+'}' = P.ID
        LEFT JOIN DISTRICTS D ON 'Struct{code='+H.DISTRICT_CODE+'}' = D.ID
        LEFT JOIN WARDS W ON 'Struct{code='+H.WARD_CODE+'}' = W.ID
        LEFT JOIN HOTEL_FACILITY_LIST HF ON H.ID = HF.HOTEL_ID
PARTITION BY H.ID;


