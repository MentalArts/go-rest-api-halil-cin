## SIMPLE REST API with GO and Gin-Gonic
This task's main goal is creating a go app that can interact with a database and works in Docker environment with additional improvements.
## Preamble
My personal goal was creating a FUNCTIONING application rather than fancy one. I could implement ADDITIONAL FEATURES as our task giver stated as enhancements. Few more additions 
to docker-compose for grafana, prometheus would run smoothly. But I already tried to exceed my GO 101 knowledge. That would be overkill and it would be nothing but
go-rest-api-chatgpt. The project is already made with assistance of LLMs.. I learned a lot along the way. And wont stop trying to improve my capabilities on this programming 
language. But it will take practice and consistency. Only one project wont bring me practical level... 

## Requirements
You need to have 
- Docker
- Docker-compose
- Go 1.23 (O)
to be able to run this API smoothly. 

## Installation
 After cloning this repo to your local machine you should navigate to the directory.
```
docker compose up 
```
Will build the containers and run it.

## To TEST it 
You can either use Insomnia or visit; 
```
http://localhost:8080/"WRITE-A-ROUTE-HERE"
```
you can try and use all the endpoints listed below:

You dont have to use User system, Keep in mind all the endpoints are accessible for public except DELETE. It is a protected route only for Admins.
```
- /api/v1/auth/register
- /api/v1/auth/login
- /api/v1/auth/refresh-token

- /api/v1/books
- /api/v1/book/"Book ID"
- /api/v1/authors
- /api/v1/author/"Author ID"
- /api/v1/reviews
- /api/v1/review/"Review ID"
```
!! GET, POST, PUT requests can be done as visitor but to use DELETE you need to have registered as Admin.


## Swaagger DOCS
Can be accessed from http://localhost:8080/swagger/index.html

## Usage
Just don't..

## Support
There is no support since it works on my machine it should be working on yours too. Unless you do something wrong. Installation requires 10IQ so it should be doable for an average humanbeing.
Besided we are using docker for a reason.

## Author
Me and generous LLMs who helps me up to 20 questions with full power. (Free plan.)
