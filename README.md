# golang_examples

# Work with PostgreSQL

When you run the sqlExecute file you should setup a postgresql with given parameters in sqlExecute.go file.
<br/>
If you have docker, you can easily install.

    $docker pull postgres
    $docker run --name postgres_database -e POSTGRES_PASSWORD=123456 -d postgres


After setup postgresql, you should create database and create table with **users** name.
<br/>
Run below commands in postgresql database.

CREATE DATABASE test;

CREATE TABLE users (<br/>
    &nbsp;&nbsp;&nbsp;id serial PRIMARY KEY,<br/>
    &nbsp;&nbsp;&nbsp;email VARCHAR(100) NOT NULL,<br/>
    &nbsp;&nbsp;&nbsp;password VARCHAR(30) NOT NULL<br/>
);

Now you can run sqlExecute.go file with:
    
    $go run sqlExecute.go

<br/>

# Work with Redis

If you don't have redis you should setup manually or you can install with docker.

    $docker pull redis
    $docker run --name redis-test-instance -p 6379:6379 -d redis

After that step, you can start redisConnection.go file:

    $go run redisConnection.go

# Work with Endpoints

I create some basic endpoints with net/http package. Its working on 8081 port. You can check basicRest.go file.<br/>
I also use gorilla/mux. Package [gorilla/mux](https://github.com/gorilla/mux) implements a request router and dispatcher for matching incoming requests to their respective handler.<br/>
Run the basicRest.go file:

    $go run basicRest.go

Open new terminal and test endpoints:

    $ curl http://localhost:8081/  # it will return a message for every request
    $ curl http://localhost:8081/articles # it will return a list of articles
    $ curl -X POST http://localhost:8081/articles # it will return a message for post request.
