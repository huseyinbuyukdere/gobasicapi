# Basic Go Api Structure
![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-98.8%25-brightgreen.svg?longCache=true&style=flat)

It provides you basic api structure for development.

## Development

You can add your api to pkg/api folder and register api to service in main.go file like below.

If you want authorization by JWT you can set "IsPublic" flag as false but you need to modify login api for verify credentials. 

```bash
var routes = []models.Route{
	models.Route{HandlerFunction: api.Login, Path: "/login", MethodName: "POST", IsPublic: true},
	models.Route{HandlerFunction: api.VerifyTest, Path: "/myNewApi", MethodName: "POST", IsPublic: false},
}
```

It includes logging middleware. Logging middleware is logging all requests to log file which seperated by date.

## Run

You can run the service with the following command.

You need to set your service variables in .env file before run

In cmd folder

```bash
go run main.go
```







