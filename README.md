<p align="center">
  <img src="http://i.imgur.com/jf7Dz7y.png)">
</p>

# Blob

A simple storage API that using Redis allowing for extremely fast response times!

## Setup 

ensure you point to a valid JWTea microservice for auth! You can also customize the configuration variables via enviroment variables:

```
// redis enviroment variables
"REDIS_HOST" 
"REDIS_PORT"
"REDIS_PASSWORD"
"REDIS_DB"

// port enviroment variables
"PORT"

// jwtea enviroment variables
"JWTEA_URL" 
```

setup up is simple with the use of glide:

```
glide install
go build
./blob
```

## Routes

#### GET "/status"
Get the status of the application.

#### GET "/get/{key}"
Get value based on key

#### GET "/get/search/{key}"
Get all values that match the key 

#### POST "/login"
Login using JWTea for token based auth, there is no need for tokens for reading data only writing!. Example post body (JSON)
```
{
    "username": "test",
    "password": "shhhhh!"
}
```

#### POST "/set"
Set value based on key. Example post body (JSON)
```
{
    "token": "VALID TOKEN RETURNED FROM LOGIN",
    "key": "hello",
    "value": "hello world!"
}
```

