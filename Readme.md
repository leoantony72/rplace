## Introduction

<P>For those who don't know reddit r/place was one of the biggest social experiments where anyone can place a pixel on a 1000x1000 pixel board. 

[reddit R/place Official](https://www.reddit.com/r/place/)

&nbsp;

## Technologies Used

&nbsp;

- Golang
- Gin
- gorilla websocket
- Redis
- Kafka
- zookeeper

&nbsp;

## Summary

&nbsp;

I am new to golang so the project structure might be weird for those who are checking it out. In the src there are controller, services, model, config. Folders such as services and model will be framework agnostic. Controllers leverage services and service folder(data access layer) contains all the business logic and will be able to interact with Model(database layer).

When request comes in Controllers pass it to service and does some work or is send to the database through the model.Results are passed back up to the service layer and then to the controllers.

<hr>

