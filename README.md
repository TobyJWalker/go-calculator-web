# go-calculator-web
Go calculator converted to a simple API. You can make POST requests to the API with a maths expression in the body and it will return the result. This project is to keep practicing Go and to learn how to make a simple API in Go. It will also save the maths expression and result to a database.

I will also experiment with building a Kubernetes cluster and separating the app into microservices. The microservices will include:

- Postgres database
- API
- Basic Frontend

Frontend will be a simple HTML page with a form to submit the maths expression. The API will be responsible for handling the request and sending it to the database. The database will be responsible for storing the maths expression and result. The API will then return the result to the frontend. It can be found here: 

https://github.com/TobyJWalker/go-calculator-frontend


## How to run

There are a few options to run this app. The first is to run it directly with Go. This can be done by running the following command in the root directory of the project:

```bash
go run .
```

The second option is to build the project and run the executable. This can be done by running the following commands in the root directory of the project:

```bash
go build -o bin/web-calculator
```


## How to use

Once the app is running, it will be active on port 8082. You can make a POST request to the API with a simple maths equation and it will return a response with the result. The body of the request should be a JSON object with the following structure:

```json
{
    "equation": "1 + 1"
}
```

The response will be a JSON object with the following structure:

```json
{
    "equation": "1 + 1",
    "result": 2
}
```
