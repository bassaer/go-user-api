# go-user-api

[![Build Status](https://travis-ci.com/bassaer/go-user-api.svg?branch=master)](https://travis-ci.com/bassaer/go-user-api)

### build

```
❯ docker-compose up -d --build
```

### log
```
❯ docker-compose logs -f
```

### usage

```
❯ curl -sS '192.168.33.10:8080' -X POST -d "{\"name\": \"foobar\"}" | python -m json.tool
{
    "id": "173067cf-2c5a-4871-8309-17b53255576e",
    "name": "foobar",
    "created_at": "2019-08-25T02:55:43.762111941+09:00"
}
```
```
❯ curl -sS '192.168.33.10:8080?id=173067cf-2c5a-4871-8309-17b53255576e' | python -m json.tool
{
    "id": "173067cf-2c5a-4871-8309-17b53255576e",
    "name": "foobar",
    "created_at": "2019-08-25T02:55:44+09:00"
}
```
