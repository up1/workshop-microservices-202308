# Working with PostgreSQL
* Create db.go => connect to database
* Update server.go
  * Create database connection
  * Inject to handler/router
* Working with Docker and Docker compose
* Testing


## How to inject all dependencies into Handler/Router ?

### Solution :: Send to SayHi(database)
* Add model.go
* Update server.go
* Update handler.go
* Update handler_test.go

server.go
```
// Database
connection := CreateDbConnection()

// Handler
hello.e.GET("/hello", SayHi(connection))
```

handler.go
```
func SayHi(connection *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		data, _ := GetData(connection)
		res := HelloResponse{Message: data}
		return c.JSON(http.StatusOK, res)
	}
}

// Get data from database
func GetData(connection *pgxpool.Pool) (string, error) {
	rows, err := connection.Query(context.Background(), "select * from message")
	if err != nil {
		return "", err
	}

	// iterate through the rows
	if rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return "", err
		}

		// convert DB types to Go types
		data := values[1].(string)
		return data, nil
	}
	return "", nil
}
```

Testing ...

## Run with docker
Start database
```
$docker compose build
$docker compose up -d db
$docker compose ps

NAME                       IMAGE               COMMAND                  SERVICE             CREATED              STATUS              PORTS
04-hello-postgresql-db-1   postgres            "docker-entrypoint.s…"   db                  About a minute ago   Up About a minute   0.0.0.0:5432->5432/tcp
```

Start Server
```
$go run cmd/main.go  

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.11.1
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323
```
or Start with Docker compose
```
$docker compose build
$docker compose up -d api
```


Testing with postman ...
```
{
    "message": "Hello world from database"
}
```

## Testing with go !!
Got error
```
$go test ./...
```

Try to fix !!

