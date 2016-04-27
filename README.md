
Document examples
-----------------

### transaction ###

```
{
    'guid':         '237bb707-3b52-11e5-ba37-00d4ee77da23',
    'type':         '',
    'datetime':     '',
    'name':         'milk',
    'amount':       1,
    'currency':     '',
    'description':  '',
}
```

### outflow ###

```
{
    'guid':         '237bb707-3b52-11e5-ba37-00d4ee77da23',
    'datetime':     '',
    'name':         'milk',
    'amount':       1,
    'currency':     '',
    'description':  '',
    'destination':  ''
    'count':        ''
    'metric_unit':  'liter',
    'target':       '',
    'satisfaction': '',
}
```

### inflow ###

```
{
    'guid':         '237bb707-3b52-11e5-ba37-00d4ee77da23',
    'datetime':     '',
    'name':         '',
    'amount':       '',
    'currency':     '',
    'description':  '',
    'source':       '',
}
```

Server params
-------------

-addr=127.0.0.1:8080
-file=sql_file.db
-currency=RUB

RESTful API
-----------

```

$ go get -u github.com/go-swagger/go-swagger/cmd/swagger
$ go install github.com/go-swagger/go-swagger/cmd/swagger
$ swagger generate server -f swagger.yaml

```

Tools
-----

- [mux handler](github.com/gorilla/mux)
- [db sqlite3](github.com/mattn/go-sqlite3)
- [guid func](http://play.golang.org/p/4FkNSiUDMg)
- [guid lib](github.com/satori/go.uuid)
- [swagger editor](http://swagger.io/swagger-editor/)
- [JSON API for foreign exchange rates and currency conversion](http://fixer.io/)

Code generation:

- [swagger-codegen](https://github.com/swagger-api/swagger-codegen)
- [bin data](https://github.com/jteeuwen/go-bindata)
- [go-swagger](https://github.com/go-swagger/go-swagger)
- [swagger-js](https://github.com/swagger-api/swagger-js)

Links
-----

- [currency symbols](http://www.currencysymbols.in/)
