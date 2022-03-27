# gorm-layer-architecture-sample

#### To run the application
```shell
go run main.go
```

## Architecture
### Simple Layer Architecture
![Layer Architecture](https://camo.githubusercontent.com/d9b21eb50ef70dcaebf5a874559608f475e22c799bc66fcf99fb01f08576540f/68747470733a2f2f63646e2d696d616765732d312e6d656469756d2e636f6d2f6d61782f3830302f312a4a4459546c4b3030796730496c556a5a392d737037512e706e67)

### Layer Architecture with full features
![Layer Architecture with standard features: config, health check, logging, middleware log tracing](https://camo.githubusercontent.com/aa7b739a4692eaf2b363cf9caf8b021c60082c77c98d3f8c96665b5cf4640628/68747470733a2f2f63646e2d696d616765732d312e6d656469756d2e636f6d2f6d61782f3830302f312a6d79556b504969343265593477455f494446526176412e706e67)
#### [core-go/search](https://github.com/core-go/search)
- Build the search model at http handler
- Build dynamic SQL for search
  - Build SQL for paging by page index (page) and page size (limit)
  - Build SQL to count total of records

### Search users: Support both GET and POST 
#### POST /users/search
##### *Request:* POST /users/search
In the below sample, search users with these criteria:
- get users of page "1", with page size "20"
- email="tony": get users with email starting with "tony"
- dateOfBirth between "min" and "max" (between 1953-11-16 and 1976-11-16)
- sort by phone ascending, id descending
```json
{
    "page": 1,
    "limit": 20,
    "sort": "phone,-id",
    "email": "tony",
    "dateOfBirth": {
        "min": "1953-11-16T00:00:00+07:00",
        "max": "1976-11-16T00:00:00+07:00"
    }
}
```
##### GET /users/search?page=1&limit=2&email=tony&dateOfBirth.min=1953-11-16T00:00:00+07:00&dateOfBirth.max=1976-11-16T00:00:00+07:00&sort=phone,-id
In this sample, search users with these criteria:
- get users of page "1", with page size "20"
- email="tony": get users with email starting with "tony"
- dateOfBirth between "min" and "max" (between 1953-11-16 and 1976-11-16)
- sort by phone ascending, id descending

#### *Response:*
- total: total of users, which is used to calculate numbers of pages at client 
- list: list of users
```json
{
    "list": [
        {
            "id": "ironman",
            "username": "tony.stark",
            "email": "tony.stark@gmail.com",
            "phone": "0987654321",
            "dateOfBirth": "1963-03-24T17:00:00Z"
        }
    ],
    "total": 1
}
```

## API Design
### Common HTTP methods
- GET: retrieve a representation of the resource
- POST: create a new resource
- PUT: update the resource
- PATCH: perform a partial update of a resource
- DELETE: delete a resource

## API design for health check
To check if the service is available.
#### *Request:* GET /health
#### *Response:*
```json
{
    "status": "UP",
    "details": {
        "sql": {
            "status": "UP"
        }
    }
}
```

## API design for users
#### *Resource:* users

### Get all users
#### *Request:* GET /users
#### *Response:*
```json
[
    {
        "id": "spiderman",
        "username": "peter.parker",
        "email": "peter.parker@gmail.com",
        "phone": "0987654321",
        "dateOfBirth": "1962-08-25T16:59:59.999Z"
    },
    {
        "id": "wolverine",
        "username": "james.howlett",
        "email": "james.howlett@gmail.com",
        "phone": "0987654321",
        "dateOfBirth": "1974-11-16T16:59:59.999Z"
    }
]
```

### Get one user by id
#### *Request:* GET /users/:id
```shell
GET /users/wolverine
```
#### *Response:*
```json
{
    "id": "wolverine",
    "username": "james.howlett",
    "email": "james.howlett@gmail.com",
    "phone": "0987654321",
    "dateOfBirth": "1974-11-16T16:59:59.999Z"
}
```

### Create a new user
#### *Request:* POST /users 
```json
{
    "id": "wolverine",
    "username": "james.howlett",
    "email": "james.howlett@gmail.com",
    "phone": "0987654321",
    "dateOfBirth": "1974-11-16T16:59:59.999Z"
}
```
#### *Response:* 1: success, 0: duplicate key, -1: error
```json
1
```

### Update one user by id
#### *Request:* PUT /users/:id
```shell
PUT /users/wolverine
```
```json
{
    "username": "james.howlett",
    "email": "james.howlett@gmail.com",
    "phone": "0987654321",
    "dateOfBirth": "1974-11-16T16:59:59.999Z"
}
```
#### *Response:* 1: success, 0: not found, -1: error
```json
1
```

### Delete a new user by id
#### *Request:* DELETE /users/:id
```shell
DELETE /users/wolverine
```
#### *Response:* 1: success, 0: not found, -1: error
```json
1
```

## Common libraries
- [core-go/health](https://github.com/core-go/health): include HealthHandler, HealthChecker, SqlHealthChecker
- [core-go/config](https://github.com/core-go/config): to load the config file, and merge with other environments (SIT, UAT, ENV)
- [core-go/log](https://github.com/core-go/log): log and log middleware

### core-go/health
To check if the service is available, refer to [core-go/health](https://github.com/core-go/health)
#### *Request:* GET /health
#### *Response:*
```json
{
    "status": "UP",
    "details": {
        "sql": {
            "status": "UP"
        }
    }
}
```
To create health checker, and health handler
```go
    db, err := sql.Open(conf.Driver, conf.DataSourceName)
    if err != nil {
        return nil, err
    }

    sqlChecker := s.NewSqlHealthChecker(db)
    healthHandler := health.NewHealthHandler(sqlChecker)
```

To handler routing
```go
    r := mux.NewRouter()
    r.HandleFunc("/health", healthHandler.Check).Methods("GET")
```

### core-go/config
To load the config from "config.yml", in "configs" folder
```go
package main

import "github.com/core-go/config"

type Root struct {
    DB DatabaseConfig `mapstructure:"db"`
}

type DatabaseConfig struct {
    Driver         string `mapstructure:"driver"`
    DataSourceName string `mapstructure:"data_source_name"`
}

func main() {
    var conf Root
    err := config.Load(&conf, "configs/config")
    if err != nil {
        panic(err)
    }
}
```

### core-go/log *&* core-go/middleware
```go
import (
    "github.com/core-go/config"
    "github.com/core-go/log"
    m "github.com/core-go/middleware"
    "github.com/gorilla/mux"
)

func main() {
    var conf app.Root
    config.Load(&conf, "configs/config")

    r := mux.NewRouter()

    log.Initialize(conf.Log)
    r.Use(m.BuildContext)
    logger := m.NewLogger()
    r.Use(m.Logger(conf.MiddleWare, log.InfoFields, logger))
    r.Use(m.Recover(log.ErrorMsg))
}
```
To configure to ignore the health check, use "skips":
```yaml
middleware:
  skips: /health
```
