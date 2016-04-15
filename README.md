
Document example
----------------

### outflow ###

```json
{
    'guid':         '237bb707-3b52-11e5-ba37-00d4ee77da23',
    'name':         'milk',
    'metric_unit':  'liter',
    'target':       '',
    'comment':      'drink',
    'keywords':     'cow',
    'satisfaction': '',
    'price':        2221,
    'amount':       1,
    'currency':     '',
    'destination':  ''
}
```

### inflow ###

```json
{
    'guid':         '237bb707-3b52-11e5-ba37-00d4ee77da23',
    'name':         ''
    'source':       '',
    'amount':       '',
    'currency':     '',
    'description':  ''
}
```

RESTful API
-----------

GET
`/document/{guid:[a-z0-9-]+}`

GET
`/inflow`
source []


type   [optional] inflow or outflow
```

guid [required]

/documents


```
