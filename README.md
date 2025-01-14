## Logical test
located in `cmd/logical_test`

## Post Service

### Prerequisites

To launch the entire service run the command from **dev** folder

```
docker-compose up
```

After launched you can make all requests

### Run Unit Tests

```
 go test ./... -v
```

### Example Requests / Responses

#### Create post

```
    curl -X POST http://localhost:9090/posts \
     -H "Content-Type: application/json" \
     -d '{
    "author" : "Denis",
    "title" : "Greetings",
    "content" : "Hi"
         }'
```

Get a response:

```
{
    "author": "Denis",
    "content": "Hi",
    "id": "2b50c8a2-542a-4882-840f-a2de1c1411cf",
    "title": "Greetings"
}
```

#### Get Posts

```
curl -X GET http://localhost:9090/posts \
     -H "Content-Type: application/json" \
```

Get a response example:

```
{
    "posts": [
        {
            "id": "33",
            "title": "Title 33",
            "content": "Ut dolore magnam ipsum dolorem tempora. Quaerat velit quisquam etincidunt porro labore voluptatem. Tempora dolor est dolor. Dolor eius dolor dolore eius. Numquam tempora consectetur labore porro dolorem sit quaerat. Est neque ipsum sit est labore ipsum. Dolorem dolore dolor sed sed neque ut. Eius dolorem dolore voluptatem numquam est. Magnam non sit tempora neque.",
            "author": "Author 33"
        },
         {
            "id": "82",
            "title": "Title 82",
            "content": "Adipisci sit sed quisquam. Eius modi porro sed porro. Dolorem adipisci sed neque porro neque. Eius est non quisquam. Adipisci consectetur sit porro. Sed dolor voluptatem ipsum quisquam eius magnam. Numquam voluptatem dolorem consectetur ut velit dolore non. Dolor quiquia eius est.",
            "author": "Author 82"
        }
     ]
}  

```

#### Get Post

```
curl -X GET http://localhost:9090/posts/2b50c8a2-542a-4882-840f-a2de1c1411cf \
     -H "Content-Type: application/json" \
```

Get a response example:

```
{
    "post": {
        "id": "2b50c8a2-542a-4882-840f-a2de1c1411cf",
        "title": "Greetings Again",
        "content": "Second chance",
        "author": "Denis"
    }
} 
```

#### Update post

```
    curl -X PUT http://localhost:9090/posts \
     -H "Content-Type: application/json" \
     -d '{
    "title" : "Greetings Again",
    "content" : "Second chance"
}'
```

Get a response:

```
{
    "post": {
        "id": "2b50c8a2-542a-4882-840f-a2de1c1411cf",
        "title": "Greetings Again",
        "content": "Second chance",
        "author": "Denis"
    }
}
```

#### Delete post

```
    curl -X DELETE http://localhost:9090/posts/2b50c8a2-542a-4882-840f-a2de1c1411cf \
     -H "Content-Type: application/json" \
```

Get a response:

```
{
    "post": {
        "id": "2b50c8a2-542a-4882-840f-a2de1c1411cf",
        "title": "Greetings Again",
        "content": "Second chance",
        "author": "Denis"
    }
}
```
