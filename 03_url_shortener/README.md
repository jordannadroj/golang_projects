# GO URL SHORTENER

This is project 3 of my 52 in 52 golang projects. This url shortener project was done based on the tutorial from [eddywm](https://github.com/eddywm/go-shortener-wm)

## Tech Stack
- GO
- Redis
- [Gin](https://github.com/gin-gonic/gin)

## Install

### Clone the repository
```shell
git clone git@github.com:jordannadroj/52_in_52.git
cd 03_url_shortener
```
 ### Install Dependencies

```shell
go get
```

### Install Redis

[Installing Redis](https://redis.io/docs/getting-started/installation/)


## Run Application

### Start Redis Server
Open a terminal and run:
```shell
redis-server
```

### Run Application
Open a new terminal and run: 
```shell
go run main.go
```

## How to Use Example

### Test Connection is working properly

```shell
$ curl --location --request GET 'http://localhost:9808'

{"message":"Welcome to the URL Shortener API"}
```

### Shorten a URL

```shell
$ curl --location --request POST 'http://localhost:9808/create-short-url' \
--data-raw '{
"long_url": "https://example-url.com"
}'

{"message":"short url created successfully","short_url":"http://localhost:9808/5uxkLeZu"}%
```

### Use Shortened URL
```shell
curl --location --request GET 'http://localhost:9808/5uxkLeZu'
```
 or simple paste the shortened url in your browser while the application is running
