
DROP TABLE IF EXISTS `hotel_facility_types`;
CREATE TABLE `hotel_facility_types` (
    `id` int auto_increment,
    `name` varchar(50),
    `name_en` varchar(50),
    `name_vn` varchar(50),
    `description_en` varchar(250),
    `description_vn` varchar(250),
    `status` int NOT NULL DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `status` (`status`) USING BTREE
)ENGINE=InnoDB;


DROP TABLE IF EXISTS `hotel_facilities`;
CREATE TABLE `hotel_facilities` (
    `id` int auto_increment,
    `name` varchar(50),
    `name_en` varchar(50),
    `name_vn` varchar(50),
    `description_en` varchar(250),
    `description_vn` varchar(250),
    `type` int NOT NULL,
    `status` int NOT NULL DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `type` (`type`) USING BTREE,
    KEY `status` (`status`) USING BTREE
)ENGINE=InnoDB;


DROP TABLE IF EXISTS `hotel_facility_details`;
CREATE TABLE `hotel_facility_details` (
    `hotel_id` int NOT NULL,
    `facility_id` int NOT NULL,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`hotel_id`,`facility_id`),
    KEY `facility_id` (`facility_id`) USING BTREE
)ENGINE=InnoDB;


DROP TABLE IF EXISTS `room_facility_types`;
CREATE TABLE `room_facility_types` (
    `id` int auto_increment,
    `name` varchar(50),
    `name_en` varchar(50),
    `name_vn` varchar(50),
    `description_en` varchar(250),
    `description_vn` varchar(250),
    `status` int NOT NULL DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `status` (`status`) USING BTREE
)ENGINE=InnoDB;


DROP TABLE IF EXISTS `room_facilities`;
CREATE TABLE `room_facilities` (
    `id` int auto_increment,
    `name` varchar(50),
    `name_en` varchar(50),
    `name_vn` varchar(50),
    `description_en` varchar(250),
    `description_vn` varchar(250),
    `type` int NOT NULL,
    `status` int NOT NULL DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `type` (`type`) USING BTREE,
    KEY `status` (`status`) USING BTREE
)ENGINE=InnoDB;

DROP TABLE IF EXISTS `room_facility_details`;
CREATE TABLE `room_facility_details` (
    `room_id` int NOT NULL,
    `facility_id` int NOT NULL,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`room_id`,`facility_id`),
    KEY `facility_id` (`facility_id`) USING BTREE
)ENGINE=InnoDB;