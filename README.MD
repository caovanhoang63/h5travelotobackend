# H5traveloto
## Introduction
H5traveloto is an open-source project aimed at learning and research. It also serves as the final project for my Introduction to Mobile Applications and Introduction to Software Engineering courses.
Project Overview
H5traveloto is an integrated hotel management and booking system. It is a SaaS (Software as a Service) system that allows hotels to manage their properties and provides customers with the ability to search and book rooms. This repository contains the backend system. The web management interfaces and mobile application for booking can be found at the following URLs:
### [Web Management](https://github.com/caovanhoang63/h5traveloto-management)
### [Mobile Application](https://github.com/caovanhoang63/h5traveloto-booking)

The Api documentation can be found [here](https://documenter.getpostman.com/view/29234273/2sA35G2MX4)

### Key Features Provided
The project provides the following key features:
+ Hotel management: Hotels can manage their properties, including room types, room rates, and room availability.
+ Booking management: Customers can search for available rooms and book them.
+ User management: Hotels can manage their users, including customers and staff.
+ The real-time synchronization of data between databases and Elastic Search by Kafka-connect.
+ Real-time chat between customers and hotels by Socket.io.
+ Online payment with VNPAY.
+ Email notification for booking confirmation.
+ Image storage with AWS S3.
+ ...


### The backend system is implemented as a REST API, following the principles of Clean Architecture.
Technologies Used
The project utilizes the following technologies:
Golang
Docker
MySQL
MongoDB
Redis
Elastic Search
Kafka Connect
RabbitMQ
Nginx
...

## Installation
Prerequisites
Before setting up the project, ensure that you have the following prerequisites installed:

Golang 1.22 or higher
Docker
...

## Usage

1. Clone the repository:
```git clone https://github.com/caovanhoang204/h5travelotobackend.git ```
2. Navigate to the project directory:
``` cd h5travelotobackend ```
3. Install the dependencies:
``` go mod download ```
4. You need AWS S3 bucket to store images. Create a bucket and provide the necessary configuration values in the .env file. The tutorial can be found [here](https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-bucket.html).
5. You need and VNPAY account to use the online payment feature. Provide the necessary configuration values in the .env file. The tutorial can be found [here](https://sandbox.vnpayment.vn/apis/).
6. Configure the environment variables:
   + Create a .env file in the project root directory.
   + Provide the necessary configuration values in the .env file.
   + The sample [.env](/env-sample.env) file can be found in the project root directory.
7. You need to create the kafka-connect configuration file in the [elastic-kafka-redis](/docker/elastic-kafka-redis) folder. And create two .env file for two docker compose files in the [docker](/docker) folder. The sample [.env](/docker/env-sample.env) file can be found in the project root directory.
8. + You need to run the run.sh file in the [mysql-mongo-rabbit](/docker/mysql-mongo-rabbit) folder first to start all main database container.
   + You need to run the run.sh file in the [elastic-kafka-redis](/docker/elastic-kafka-redis) folder to start the docker-compose file to set up container to synchronize data from mysql to elastic search.
9. + If you run in your local machine, you just need to build the Dockerfile [here](/Dockerfile) and run the container.
   + If you run in the server, you need to set up the nginx configuration file in  the server (you need to uncomment all the setup in this [sh file](/docker/mysql-mongo-rabbit/run-database.sh) ) to redirect the request to the backend server. Then you need to change all your host Ips in all config file and run [deploy.sh file](/deploy.sh) if you are  using an SSH tool.
#### Note: If you encounter any issues with the Elasticsearch certificate, first ensure that you have provided all the necessary permissions for the certs folder. You can follow this [link](https://www.elastic.co/guide/en/elasticsearch/reference/8.13/encrypting-communications-certificates.html) for more information. 
#### If any unexpected errors occur or if you need further assistance, feel free to contact me via email at caovanhoang204@gmail.com or create an issue on the GitHub repository.

## License

This project is licensed under the MIT License - see the [LICENSE](/LICENSE) file for details.
