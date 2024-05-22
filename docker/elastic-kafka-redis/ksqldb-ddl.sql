SET 'auto.offset.reset' = 'earliest';
CREATE STREAM HOTELS WITH(KAFKA_TOPIC='mysql-debezium-h5traveloto.h5traveloto.hotels', PARTITIONS=1, VALUE_FORMAT='AVRO');
CREATE STREAM HOTEL_FACILITIES WITH (KAFKA_TOPIC='mysql-debezium-h5traveloto.h5traveloto.hotel_facility_details', PARTITIONS=1, VALUE_FORMAT='AVRO');
CREATE STREAM ROOM_FACILITIES WITH (KAFKA_TOPIC='mysql-debezium-h5traveloto.h5traveloto.room_facility_details', PARTITIONS=1 ,VALUE_FORMAT='AVRO');
CREATE STREAM ROOM_TYPES WITH (KAFKA_TOPIC='mysql-debezium-h5traveloto.h5traveloto.room_types', PARTITIONS=1 ,VALUE_FORMAT='AVRO');
CREATE STREAM HOTEL_DETAILS WITH (KAFKA_TOPIC='mysql-debezium-h5traveloto.h5traveloto.hotel_details', PARTITIONS=1 ,VALUE_FORMAT='AVRO');
CREATE STREAM PROVINCES_BASE WITH (KAFKA_TOPIC='mysql-debezium-h5traveloto.h5traveloto.provinces', PARTITIONS=1 ,VALUE_FORMAT='AVRO');
CREATE STREAM DISTRICTS_BASE WITH (KAFKA_TOPIC='mysql-debezium-h5traveloto.h5traveloto.districts', PARTITIONS=1 ,VALUE_FORMAT='AVRO');
CREATE STREAM WARDS_BASE WITH (KAFKA_TOPIC='mysql-debezium-h5traveloto.h5traveloto.wards', PARTITIONS=1 ,VALUE_FORMAT='AVRO');
CREATE STREAM LANDMARKS WITH (KAFKA_TOPIC='mysql-debezium-h5traveloto.h5traveloto.landmarks', PARTITIONS=1 ,VALUE_FORMAT='AVRO');

CREATE TABLE PROVINCES_TABLE (ID VARCHAR PRIMARY KEY)
    WITH (KAFKA_TOPIC='mysql-debezium-h5traveloto.h5traveloto.provinces',
        VALUE_FORMAT='AVRO');
CREATE TABLE DISTRICTS_TABLE (ID VARCHAR PRIMARY KEY)
    WITH (KAFKA_TOPIC='mysql-debezium-h5traveloto.h5traveloto.districts',
        VALUE_FORMAT='AVRO');
CREATE TABLE WARDS_TABLE (ID VARCHAR PRIMARY KEY)
    WITH (KAFKA_TOPIC='mysql-debezium-h5traveloto.h5traveloto.wards',
        VALUE_FORMAT='AVRO');
CREATE TABLE HOTEL_FACILITY_LIST  AS
	SELECT HF.HOTEL_ID, COLLECT_LIST(FACILITY_ID) AS FACILITY_LIST
    FROM  HOTEL_FACILITIES HF
    GROUP BY HF.HOTEL_ID;
CREATE TABLE  ROOM_FACILITY_LIST  AS
    SELECT RF.ROOM_ID, COLLECT_LIST(FACILITY_ID) AS FACILITY_LIST
    FROM  ROOM_FACILITIES RF
    GROUP BY RF.ROOM_ID;
CREATE STREAM PROVINCES
       AS
SELECT
    P.CODE as "code",
    P.NAME as "name",
    P.FULL_NAME as "full_name"
FROM PROVINCES_BASE P
    PARTITION BY P.CODE;
CREATE STREAM DISTRICTS
       AS
SELECT
    D.CODE as "code",
    D.NAME as "name",
    D.FULL_NAME as "full_name",
    STRUCT("code" := P.CODE, "name" := P.NAME,"full_name" := P.FULL_NAME) AS "province"
    FROM DISTRICTS_BASE D
    JOIN PROVINCES_TABLE P ON 'Struct{code='+D.PROVINCE_CODE+'}' = P.ID
    PARTITION BY D.CODE;
CREATE STREAM WARDS
       AS
SELECT
    W.CODE as "code",
    W.NAME as "name",
    W.FULL_NAME as "full_name",
    STRUCT("code" := D.CODE, "name" := D.NAME,"full_name" := D.FULL_NAME ) AS "district",
    STRUCT("code" := P.CODE, "name" := P.NAME,"full_name" :=P.FULL_NAME) AS "province"
FROM WARDS_BASE W
    JOIN DISTRICTS_TABLE D ON 'Struct{code='+W.DISTRICT_CODE+'}' = D.ID
    JOIN PROVINCES_TABLE P ON 'Struct{code='+D.PROVINCE_CODE+'}' = P.ID

PARTITION BY W.CODE;




CREATE STREAM ROOM_TYPES_ENRICHED AS
    SELECT
        RT.ID ID,
        RT.HOTEL_ID AS "hotel_id",
        RT.NAME AS "name",
        RT.MAX_CUSTOMER AS "max_customer",
        RT.AREA  AS "area",
        RT.IMAGES AS "images_str",
        RT.BED AS "bed_str",
        RT.PRICE AS "price",
        RT.STATUS AS "status",
        RT.TOTAL_ROOM AS "total_room",
        RT.PAY_IN_HOTEL AS "pay_in_hotel",
        RT.BREAK_FAST AS "break_fast",
        RT.FREE_CANCEL AS "free_cancel",
        RT.CREATED_AT AS "created_at",
        RT.UPDATED_AT AS "updated_at",
        RF.FACILITY_LIST AS "facility_list"
    FROM ROOM_TYPES RT
        LEFT JOIN ROOM_FACILITY_LIST RF ON RT.ID = RF.ROOM_ID
PARTITION BY ID;
CREATE STREAM HOTELS_ENRICHED
AS
SELECT
    H.ID AS ID,
    H.OWNER_ID AS "owner_id",
    H.NAME AS "name" ,
    H.HOTEL_TYPE AS "hotel_type",
    H.LOGO AS "logo_str",
    H.IMAGES AS "images_str",
    H.ADDRESS AS "address",
    STRUCT("code" := P.CODE, "name" := P.NAME,"full_name" := P.FULL_NAME) AS "province",
    STRUCT("code" := D.CODE, "name" := D.NAME,"full_name" := D.FULL_NAME) AS "district",
    STRUCT("code" := W.CODE, "name" := W.NAME,"full_name" := W.FULL_NAME) AS "ward",
    STRUCT("lat" := H.LAT, "lon" := H.LNG) AS "location",
    CAST(H.LAT AS VARCHAR)  + ',' + CAST(H.LNG AS VARCHAR) AS "location_geo_point",
    H.STAR as "star",
    STRUCT(
    "number_of_floor" := HD.NUMBER_OF_FLOOR,
    "distance_to_center_city" := HD.DISTANCE_TO_CENTER_CITY,
    "description" := HD.DESCRIPTION,
    "location_detail" := HD.LOCATION_DETAIL,
    "check_in_time" := HD.CHECK_IN_TIME,
    "check_out_time" := HD.CHECK_OUT_TIME,
    "require_document" := HD.REQUIRE_DOCUMENT,
    "minimum_age" := HD.MINIMUM_AGE,
    "cancellation_policy" := HD.CANCELLATION_POLICY,
    "smoking_policy" := HD.SMOKING_POLICY,
    "additional_policies" := HD.ADDITIONAL_POLICIES
    ) AS "hotel_details",
    H.STATUS as "status",
    H.CREATED_AT as "created_at",
    H.UPDATED_AT as "updated_at",
    HF.FACILITY_LIST AS "facility_list"
FROM HOTELS H
         LEFT JOIN PROVINCES_TABLE P ON 'Struct{code='+H.PROVINCE_CODE+'}' = P.ID
         LEFT JOIN DISTRICTS_TABLE D ON 'Struct{code='+H.DISTRICT_CODE+'}' = D.ID
         LEFT JOIN WARDS_TABLE W ON 'Struct{code='+H.WARD_CODE+'}' = W.ID
         LEFT JOIN HOTEL_FACILITY_LIST  HF ON H.ID = HF.HOTEL_ID
         LEFT JOIN HOTEL_DETAILS HD WITHIN 30 SECONDS ON H.ID = HD.HOTEL_ID
    PARTITION BY H.ID;

CREATE STREAM LANDMARKS_ENRICHED
AS
SELECT
    H.ID AS ID,
    H.NAME AS "name" ,
    H.IMAGE AS "images_str",
    STRUCT("code" := P.CODE, "name" := P.NAME,"full_name" := P.FULL_NAME) AS "province",
    STRUCT("code" := D.CODE, "name" := D.NAME,"full_name" := D.FULL_NAME) AS "district",
    STRUCT("code" := W.CODE, "name" := W.NAME,"full_name" := W.FULL_NAME) AS "ward",
    STRUCT("lat" := H.LAT, "lon" := H.LNG) AS "location",
    CAST(H.LAT AS VARCHAR)  + ',' + CAST(H.LNG AS VARCHAR) AS "location_geo_point",
    H.STATUS as "status",
    H.CREATED_AT as "created_at",
    H.UPDATED_AT as "updated_at"
FROM LANDMARKS H
         LEFT JOIN PROVINCES_TABLE P ON 'Struct{code='+H.PROVINCE_CODE+'}' = P.ID
         LEFT JOIN DISTRICTS_TABLE D ON 'Struct{code='+H.DISTRICT_CODE+'}' = D.ID
         LEFT JOIN WARDS_TABLE W ON 'Struct{code='+H.WARD_CODE+'}' = W.ID
    PARTITION BY H.ID;
