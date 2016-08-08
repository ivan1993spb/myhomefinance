
MyHomeFinance [![Build Status](https://travis-ci.org/ivan1993spb/myhomefinance.svg?branch=master)](https://travis-ci.org/ivan1993spb/myhomefinance)
=============

Install
-------

```bash
make deps
make test
make build
```

Document examples
-----------------

### Outflow ###

```
type Outflow struct {
	Id           int64
	DocumentGUID string
	Time         time.Time
	Name         string
	Amount       float64
	Description  string
	Destination  string
	Target       string
	Count        float64
	MetricUnit   string
	Satisfaction float32
}
```

### inflow ###

```
type Inflow struct {
	Id           int64
	DocumentGUID string
	Time         time.Time
	Name         string
	Amount       float64
	Description  string
	Source       string
}
```

Server params
-------------

```
-addr=127.0.0.1:8080
-file=sql_file.db
-default_satisfaction=0.5
```

Tools
-----

- [mux handler](github.com/gorilla/mux)
- [db sqlite3](github.com/mattn/go-sqlite3)
- [guid func](http://play.golang.org/p/4FkNSiUDMg)
- [guid lib](github.com/satori/go.uuid)
- [swagger editor](http://swagger.io/swagger-editor/)
- [JSON API for foreign exchange rates and currency conversion](http://fixer.io/)
- [React-bootstrap](https://github.com/react-bootstrap/react-bootstrap)

Code generation:

- [swagger-codegen](https://github.com/swagger-api/swagger-codegen)
- [bin data](https://github.com/jteeuwen/go-bindata)
- [go-swagger](https://github.com/go-swagger/go-swagger)
- [swagger-js](https://github.com/swagger-api/swagger-js)

Links
-----

- [currency symbols](http://www.currencysymbols.in/)
