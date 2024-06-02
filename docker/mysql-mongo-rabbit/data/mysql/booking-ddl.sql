

DROP TABLE IF EXISTS `hotel_types`;
CREATE TABLE `hotel_types` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL,
    `description` varchar(250),
    `status` int DEFAULT 1,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `status` (`status`) USING BTREE
)ENGINE=InnoDB;

DROP TABLE IF EXISTS `hotel_details`;
CREATE TABLE `hotel_details` (
  `id` int NOT NULL AUTO_INCREMENT,
  `hotel_id` int NOT NULL,
  `number_of_floor` int NOT NULL DEFAULT 1,
  `distance_to_center_city` float NOT NULL DEFAULT 0,
  `description` text,
  `location_detail` text,
  `check_in_time` time,
  `check_out_time` time,
  `require_document` boolean,
  `minimum_age` int,
  `cancellation_policy` float default 0,
  `smoking_policy` enum('free','limit','ban') default 'free',
  `pet_policy` enum('free','large','small','ban') default 'free',
   `additional_policies` text,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   PRIMARY KEY (`id`),
   KEY `hotel_id` (`hotel_id`) USING BTREE
) ENGINE=InnoDB;


DROP TABLE IF EXISTS `hotels`;
CREATE TABLE `hotels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `owner_id` int NOT NULL,
  `name` varchar(50) NOT NULL,
  `address` varchar(255) NOT NULL,
  `hotel_type` int NOT NULL,
  `hotline` varchar(20) NOT NULL,
  `logo` json,
  `images` json,
  `province_code`  varchar(20) NOT NULL,
  `district_code` varchar(20) NOT NULL,
  `ward_code` varchar(20) NOT NULL,
  `lat` double NOT NULL,
  `lng` double NOT NULL,
  `rating` float default 0,
  `avg_price` float default 0,
  `total_room_type` int default 0,
  `total_rating` int default 0,
  `star` int default 1,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `owner_id` (`owner_id`) USING BTREE,
  KEY `province_code` (`province_code`) USING BTREE,
  KEY `district_code` (`district_code`) USING BTREE,
  KEY `hotel_type` (`hotel_type`) USING BTREE,
  KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB;

ALTER TABLE `hotels`
ADD COLUMN `total_room` int NOT NULL DEFAULT 0 AFTER `total_room_type`;




DROP TABLE IF EXISTS `room_types`;
CREATE TABLE `room_types` (
    `id` int NOT NULL AUTO_INCREMENT,
    `hotel_id` int NOT NULL,
    `name` varchar(50) NOT NULL,
    `max_customer` integer DEFAULT 1,
    `area` double,
    `bed` json NOT NULL,
    `price` DECIMAL(19,4)  NOT NULL,
    `cur_available_room` int,
    `images` json,
    `description` text,
    `total_room` int NOT NULL DEFAULT 0,
    `pay_in_hotel` bool DEFAULT false,
    `break_fast` boolean DEFAULT false,
    `free_cancel` boolean DEFAULT false,
    `status` int NOT NULL DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `hotels_id` (`hotel_id`) USING BTREE,
    KEY `status` (`status`) USING BTREE
)ENGINE=InnoDB;



DROP TABLE IF EXISTS `workers`;
CREATE TABLE `workers` (
    `user_id` int NOT NULL,
    `hotel_id`int NOT NULL,
    `status` int NOT NULL DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`user_id`,`hotel_id`),
    KEY `hotel_id` (`hotel_id`) USING BTREE,
    KEY `status` (`status`) USING BTREE
)ENGINE=InnoDB;


DROP TABLE IF EXISTS `rooms`;
CREATE TABLE `rooms` (
    `id` int NOT NULL AUTO_INCREMENT,
    `hotel_id` int NOT NULL,
    `room_type_id` int NOT NULL,
    `name` varchar(50) NOT NULL,
    `floor` int NOT NULL DEFAULT  1,
    `condition` enum('available','booked','dirty','fixing','reserved') DEFAULT 'available',
    `status` int NOT NULL DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `hotels_id` (`hotel_id`) USING BTREE,
    KEY `room_type_id` (`room_type_id`) USING BTREE,
    KEY `status` (`status`) USING BTREE
)ENGINE=InnoDB;




DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(50) NOT NULL,
  `fb_id` varchar(50) DEFAULT NULL,
  `gg_id` varchar(50) DEFAULT NULL,
  `address` varchar(100) DEFAULT NULL,
  `password` varchar(255) NOT NULL,
  `salt` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) NOT NULL,
  `first_name` varchar(50) NOT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `role` enum('customer','admin','staff','manager','owner') NOT NULL DEFAULT 'customer',
  `avatar` json DEFAULT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `status` (`status`) USING BTREE,
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB;

ALTER TABLE `users`
ADD COLUMN `date_of_birth` date DEFAULT NULL AFTER `avatar`,
ADD COLUMN `gender` enum('male','female','other') default 'other' AFTER `date_of_birth`;


DROP TABLE IF EXISTS `bookings`;
CREATE TABLE `bookings` (
    `id` int NOT NULL AUTO_INCREMENT,
    `hotel_id` int NOT NULL,
    `user_id` int,
    `adults` int NOT NULL,
    `children` int NOT NULL,
    `room_type_id` int NOT NULL,
    `room_quantity` int ,
    `deal_id` int default null,
    `total_amount` decimal(19,4),
    `discount_amount` decimal(19,4),
    `final_amount` decimal(19,4),
    `currency` varchar(3),
    `start_date` date NOT NULL,
    `end_date` date NOT NULL ,
    `status` int NOT NULL DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `hotel_id` (`hotel_id`) USING BTREE ,
    KEY `user_id` (`user_id`) USING BTREE ,
    KEY `status` (`status`) USING BTREE
)ENGINE=InnoDB;


ALTER TABLE `bookings`
ADD COLUMN `state` enum('pending','paid','canceled','checked-in','checked-out','deleted','expired') default 'pending' after currency ,
ADD COLUMN `pay_in_hotel` boolean default false after state;
ALTER TABLE `bookings`
ADD  KEY `state` (`state`) USING BTREE,
ADD KEY `pay_in_hotel` (`pay_in_hotel`) USING BTREE ;



DROP TABLE IF EXISTS `booking_details`;
CREATE TABLE `booking_details` (
    `booking_id` int NOT NULL,
    `room_id` int NOT NULL,
    `status` int NOT NULL DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`booking_id`,`room_id`),
    KEY `room_id` (`room_id`) USING BTREE ,
    KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB;


DROP TABLE IF EXISTS `booking_trackings`;
CREATE TABLE `booking_trackings` (
    `id` int NOT NULL AUTO_INCREMENT,
    `booking_id` int NOT NULL,
    `state` enum('pending','paid','confirmed','canceled','checked-in','checked-out','deleted','expired') NOT NULL DEFAULT 'pending',
    `status` int NOT NULL DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `booking_id` (`booking_id`) USING BTREE ,
    KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `deals`;
CREATE TABLE `deals` (
    `id` int NOT NULL AUTO_INCREMENT,
    `hotel_id` int NOT NULL,
    `room_type_id` int,
    `name` varchar(100),
    `image` json,
    `description` text,
    `total_quantity` int default 0,
    `available_quantity` int default 0,
    `min_price` decimal default 0,
    `discount_type` enum('percent','cash') default 'percent',
    `discount_amount` float default 0,
    `discount_percent` float default 0,
    `is_unlimited` boolean default false,
    `start_date` timestamp,
    `expiry_date` timestamp,
    `status` int NOT NULL DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `hotel_id` (`hotel_id`) USING BTREE ,
    KEY `room_type_id` (`room_type_id`) USING BTREE ,
    KEY `status` (`status`) USING BTREE
) ENGINE =InnoDb;

# PAYMENT
DROP TABLE IF EXISTS `payment_events`;
CREATE TABLE `payment_events` (
    `txn_id` varchar(50) not null ,
    `hotel_id` int not null,
    `customer_id` int not null default 0,
    `is_payment_done` boolean default false,
    `payment_type` enum('pay_in','pay_out'),
    `status` int NOT NULL DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`txn_id`),
    KEY `hotel_id` (hotel_id) USING BTREE ,
    KEY `customer_id` (customer_id) USING BTREE,
    KEY `status` (`status`) USING BTREE
) ENGINE = InnoDb;



DROP TABLE IF EXISTS `payment_bookings`;
CREATE TABLE `payment_bookings`
(
    `booking_id` int NOT NULL,
    `txn_id` varchar(50),
    `customer_id` int NOT NULL ,
    `hotel_id` int NOT NULL,
    `amount` decimal(19,4) NOT NULL,
    `currency` varchar(3) NOT NULL,
    `method` enum('vnpay','paypal','hbanking','visa'),
    `payment_status` enum('not_started','executing','success','failed','expired') default 'not_started',
    `ledger_updated` boolean default false,
    `wallet_updated` boolean default false,
    `status` int NOT NULL DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`booking_id`,`txn_id`),
    KEY `hotel_id` (`hotel_id`) USING BTREE,
    KEY `status` (`status`) USING BTREE
) ENGINE = InnoDb;



DROP TABLE IF EXISTS `hotel_wallets`;
CREATE TABLE `hotel_wallets` (
    `id` int not null auto_increment,
    `hotel_id` int not null,
    `balance` varchar(255),
    `currency` varchar(3) NOT NULL,
    `status` int NOT NULL DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `hotel_id` (`hotel_id`) USING BTREE,
    KEY `status` (`status`) USING BTREE
) ENGINE = InnoDb;

DROP TABLE IF EXISTS `ledgers`;
CREATE TABLE `ledgers` (
    `id` int not null auto_increment,
    `transaction_id` int not null,
    `account_id` int not null,
    `debit` decimal(19,4) ,
    `credit` decimal(19,4),
    `currency` VARCHAR(3),
    `status` int NOT NULL DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `account_id` (`account_id`) USING BTREE ,
    KEY `transaction_id` (`transaction_id`) USING BTREE ,
    KEY `status` (`status`) USING BTREE
) ENGINE = InnoDb;

DROP TABLE IF EXISTS `hotel_payment_info`;
CREATE TABLE `hotel_payment_info` (
    `id` int  auto_increment,
    `type` enum('paypal','bank_transfer') not null ,
    `detail_id` int not null,
    `status` int NOT NULL DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `detail_id` (`detail_id`) USING BTREE,
    KEY `type` (`type` ) USING BTREE,
    UNIQUE (`detail_id`,`type`),
    KEY `status` (`status`) USING BTREE
) ENGINE = InnoDb;


DROP TABLE IF EXISTS `paypal_info`;
CREATE TABLE `paypal_info` (
    `id` INT AUTO_INCREMENT,
    `hotel_id` INT,
    `salt` varchar(50) DEFAULT NULL,
    `email` VARCHAR(255),
    `status` int NOT NULL DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `hotel_id` (hotel_id) USING BTREE,
    KEY `status` (`status`) USING BTREE,
    UNIQUE KEY `email` (`email`)
)ENGINE = InnoDb;

DROP TABLE IF EXISTS `bank_info`;
CREATE TABLE `bank_info` (
    `id` INT AUTO_INCREMENT,
    `hotel_id` INT  NOT NULL ,
    `bank_name` VARCHAR(50) NOT NULL ,
    `bank_branch` VARCHAR(50) NOT NULL ,
    `salt` varchar(50) DEFAULT NULL,
    `account_name` VARCHAR(50) NOT NULL ,
    `account_number` VARCHAR(255) NOT NULL ,
    `swift_code` VARCHAR(50) NOT NULL ,
    `status` int NOT NULL DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `hotel_id` (hotel_id) USING BTREE,
    KEY `status` (`status`) USING BTREE,
    UNIQUE(`swift_code`,`account_number`)
)ENGINE = InnoDb;



DROP TABLE IF EXISTS `hotels_saved`;
CREATE TABLE `hotels_saved` (
    `user_id` INT NOT NULL ,
    `hotel_id` INT NOT NULL,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`,`hotel_id`),
    KEY `hotel_id` (hotel_id) USING BTREE
)ENGINE = InnoDb;

SELECT * From `hotels_saved`;


DROP TABLE IF EXISTS `hotels_collections`;
CREATE TABLE `hotels_collections` (
    `id` INT NOT NULL  AUTO_INCREMENT ,
    `user_id` INT NOT NULL ,
    `name` VARCHAR(250) NOT NULL DEFAULT 'untitled',
    `cover` json,
    `status` int NOT NULL DEFAULT '1',
    `is_private` boolean DEFAULT true,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `user_id` (user_id) USING BTREE,
    KEY `status` (`status`) USING BTREE
)ENGINE = InnoDb;

DROP TABLE IF EXISTS `hotels_collection_details`;
CREATE TABLE `hotels_collection_details` (
    `collection_id` INT NOT NULL,
    `hotel_id` INT NOT NULL ,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`collection_id`,`hotel_id`)
)ENGINE = InnoDb;

