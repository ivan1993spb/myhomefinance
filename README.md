
Document examples
-----------------

### transaction ###

```json
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

```json
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

```json
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

RESTful API
-----------


GET `/inflow`


GET `/inflow/{guid}`


GET `/inflow/grep`
'name'
'currency'
'description'
'source'

POST `/inflow/add`
'name'
'amount'
'currency'
'description'
'source'

Result: inflow guid


POST `/inflow/{guid}/delete`


POST `/inflow/{guid}/edit`

'name'
'amount'
'currency'
'description'
'source'

Result: inflow guid


GET `/outflow`


GET `/outflow/{guid}`


GET `/outflow/grep`
'name'
'currency'
'description'
'destination'
'target'
'satisfaction'


POST `/outflow/add`

'name'
'amount'
'currency'
'description'
'destination'
'metric_unit'
'target'
'satisfaction'

Result: outflow guid


POST `/outflow/{guid}/delete`


POST `/outflow/{guid}/edit`

'name'
'amount'
'currency'
'description'
'destination'
'metric_unit'
'target'
'satisfaction'

Result: outflow guid


GET `/transactions`


GET `/transactions/{guid}`


GET `/transactions/grep`
'name'
'currency'
'description'


GET `/transactions/year`


GET `/transactions/halfyaer`


GET `/transactions/month`


GET `/transactions/week`




Tools
-----

- [mux handler](github.com/gorilla/mux)
- [db sqlite3](github.com/mattn/go-sqlite3)
- [guid func](http://play.golang.org/p/4FkNSiUDMg)
- [guid lib](github.com/satori/go.uuid)
- [swagger editor](http://swagger.io/swagger-editor/)
