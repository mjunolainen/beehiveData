
# Project beehiveData for hive health monitoring [0/12]

## Introduction
The goal of this project is to set up a monitoring environment for beehives. Sensors to monitor the variables will be added one by one, as not to make the 
too complicated. Current plan is as follows: temperature/humidity/barometric pressure inside the hive --> hive weight --> temperature/humidity/barometric pressure outside the hive --> camera --> bee counter --> sound analysis. 


## User Stories
- [ ] As an administrator I want to have a section in README.md a list of steps by subsystems how to set up and configure different used services  (Raspberry Pi and it's sensors, database, Grafana and website.
- [ ] As a developer and administrator I want to keep track of my work in git, regularly updating README.md other files with concise and descriptive commit messages.
- [ ] As a developer I want to have Raspberry Pi connected to its sensors the right way and all the software installed in order to be able to run receive data from the sensors and run the script on it. 
- [ ] As a developer I want to have a solid code running on Raspberry Pi that is receiving data from the sensors and storing it to the database. 
- [ ] As an administrator and a developer I want to have the database either configured (if for example using BigQuery or similar) or set up if I decide to use a self hosted version. 
- [ ] As a developer I want to connect Grafana to my chosen database to see the received events. 
- [ ] As an administrator I want to have the Apollo GraphQL server configured with queries that fetch events from the database with parameters 'from'/'until' that specify time range.
- [ ] As a developer I want to create a JS or React project for displaying the analytics dashboard.
- [ ] As an developer I want to integrate Apollo GraphQL with my project. 
- [ ] As a developer I want toquery data from Apollo GraphQL within my project using time range and possibly other variables. 
- [ ] As a data analyst I want to see a visual representation of received events in the app dashboard. 
- [ ] As an administrator I want to host the app somewhere and run it @ www.urbanbees.earth


## Resources
- React app with Apollo GraphQL https://www.freecodecamp.org/news/apollo-graphql-how-to-build-a-full-stack-app-with-react-and-node-js/
- Docker 1h intro https://www.youtube.com/watch?v=pTFZFxd4hOI
- Docker and node.js app https://nodejs.org/en/docs/guides/nodejs-docker-webapp/


# Setup and running

## Raspberry Pi and sensors
1. ... define steps here urls to visit, how to verify the service is configured and working...
2. https://projects.raspberrypi.org/en/projects/raspberry-pi-setting-up/1
3. https://medium.com/initial-state/how-to-build-a-raspberry-pi-temperature-monitor-8c2f70acaea9#
4. Connecting DHT22 temperature / humidity sensor to Raspberry Pi - https://github.com/d2r2/go-dht 
5. https://periph.io/
6. https://blog.inkdrop.app/a-simple-room-air-quality-visualizer-using-raspberry-pi-and-golang-46e6464b2e8d
7. https://willj.net/posts/humidity-temperature-and-pressure-sensing-on-a-raspberry-pi-with-go/
8. https://www.jeremymorgan.com/tutorials/go/get-temperature-raspberry-pi-go/
9. BME280 I2C or SPI Temperature Humidity Pressure Sensor - https://www.adafruit.com/product/2652
10. https://github.com/InitialState/fona-pi-zero/wiki
11. 
 

## Database
1. ... define steps here urls to visit, how to verify the service is configured and working...
2. https://firebase.google.com/products/firestore

## Apollo GraphQL
1. ... define steps here urls to visit, how to verify the service is configured and working...
2. https://medium.com/@lukepighetti/yes-you-can-query-firebase-with-graphql-e79a45990f22

## Grafana
1. ... define steps here urls to visit, how to verify the service is configured and working...


## 'dashboard' app
1. ... define steps here urls to visit, how to verify the service is configured and working...
2. https://medium.com/initial-state/build-a-serverless-iot-application-using-aws-amplify-and-vue-efee0126ae33
