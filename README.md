
Document example
----------------

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


GET `/transactions/all`


GET `/transactions/year`


GET `/transactions/halfyaer`


GET `/transactions/month`


GET `/transactions/week`

