# go-user-api

[![Build Status](https://travis-ci.com/bassaer/go-user-api.svg?branch=master)](https://travis-ci.com/bassaer/go-user-api)

```
❯ docker-compose up -d --build
```

```
❯ curl '192.168.33.10:8080'
{"id":"testid","name":"testname","created_at":"2019-08-21T14:06:53.145415733Z"}
```

```
❯ mysql -u test -p -s -t --host=127.0.0.1
Enter password:
mysql> select * from userdb.users;
+--------+----------+---------------------+
| id     | name     | created_at          |
+--------+----------+---------------------+
| testid | testuser | 2019-08-22 10:31:08 |
+--------+----------+---------------------+
```
