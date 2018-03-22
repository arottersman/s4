# s4
_Simple, Stupid, Storage Service_

A very basic key-value store service that can set a string, numeric or boolean
value to a key, and retrieve it again.

## Requirements
- [Docker v17+](https://docs.docker.com/install/)

## Get started
Start the server:
```bash
$ ./scripts/server.sh
```

Set a key:
```bash
$ curl -v -X PUT http://localhost:8080/items -d '{"key": "name", "value": "Pepper"}'

# > PUT /items HTTP/1.1                                                                       [0/1676]> Host: localhost:8080
# > User-Agent: curl/7.54.0
# > Accept: */*
# > Content-Length: 34
# > Content-Type: application/x-www-form-urlencoded
# >
# * upload completely sent off: 34 out of 34 bytes
# < HTTP/1.1 201 Created
# < Content-Type: application/json; charset=UTF-8
# < Date: Thu, 22 Mar 2018 02:09:19 GMT
# < Content-Length: 0
```

Read that key back out again
```bash
$ curl http://localhost:8080/items/name

# > GET /items/name HTTP/1.1
# > Host: localhost:8080
# > User-Agent: curl/7.54.0
# > Accept: */*
# >
# < HTTP/1.1 200 OK
# < Content-Type: application/json; charset=UTF-8
# < Date: Thu, 22 Mar 2018 02:34:32 GMT
# < Content-Length: 32
# <
# {"key":"name","value":"Pepper"}
```

## API

### PUT `/items`
Set a key-value pair. The key must be URL-encodable and the value
must be a string, number, or boolean.
#### `application/json` Request Body
```json
{
    "key": "keyYouWantToStoreYourDataAt",
    "value": "Some string, numeric or boolean value" 
}
```
#### JSON Response Body
N/A

### GET `/items/{key}`
Get the value stored at the key. Responds
with 404 if the key is not found.

#### `application/json` Request Body
N/A
#### JSON Response Body
```json
{
    "key": "theKey",
    "value": "Your value as a string"
}
```

## Scripts

| Script | Description                  | Ports |
|--------|------------------------------|-------|
| server | start the development server | 8080  |
| test   | run the unit tests           |       |
