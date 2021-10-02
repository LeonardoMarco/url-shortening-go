### Shortener GO
URL shortening project using GO language and Gin Web Framework.

#### Config

- Use Postgres Database
- Alter info database in setup.go
- Execute command line
```
psql -U {{user_database}} -d {{name_database} -a -f db/migrations/db.sql
```

<hr />

#### Run project

```
go run main.go
```

or execute binnary

```
./shortener
```

<hr />

#### Build project

```
go build
```

<hr />

#### Use project

###### Create new short url

Request:
```
curl --location --request POST 'http://localhost:8080/' \
--data-raw '[{
    "url": "https://leonardomarco.dev"
}]'
```

Response:
```
{
  "data": "ZwkL9OR"
}
```

<hr >


###### Access Url:

Access Url in browser

obs: Hash is data return in create short url route.

```
http://localhost:8080/ZwkL9OR
```


