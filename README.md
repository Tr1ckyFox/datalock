# datalock

Follow the rules on http://bit.ly/2rdoNTn

## Information for developers

You can specify the format by adding `&_format=xml` or `&_format=json`; JSON is used as the default data format.

### Season details

`GET /api/info_season?url=/serial-15825-Nelyudi-0-season.html`

Output format:

```json
{
    "title": "...",
    "id": 0,
    "serial": 0,
    "keywords": "...",
    "description": "..."
}
```

### Available seasons

`GET /api/all_seasons?url=/serial-15825-Nelyudi-0-season.html`

Output format:

```json
[{
    "title": "...",
    "link": "..."
}]
```

### Available playlists

`GET /api/all_series?url=/serial-15825-Nelyudi-0-season.html`

Output format:

```json
[{
    "name": "...",
    "playlist": [{
        "id": "...",
        "title": "...",
        "subtitle": "...",
        "file": "...",
        "galabel": "..."
    }]
}]
```

### Feeds

#### Updated series

`GET /api/updated_series`

Output format:

```json
[{
    "name": "...",
    "link": "...",
    "comment": "..."
}]
```

#### Popular series

`GET /api/popular_series`

Output format:

```json
[{
    "name": "...",
    "link": "...",
    "comment": "..."
}]
```


#### New series

`GET /api/new_series`

Output format:

```json
[{
    "name": "...",
    "link": "...",
    "comment": "..."
}]
```
