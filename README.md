# TODO_APP by golang React Typescript

## USER

users/create

```
curl -v -H "Accept: application/json"  -H "Content-Type: application/json" -d '{"name":"user1","email":"sample@example.com","password":"pass"}' http://localhost:8080/users/new
```

users/show

```
curl -v -H "Accept: application/json"  -H "Content-Type: application/json" http://localhost:8080/users/show/1
```

users/update

```
curl -v -H "Accept: application/json"  -H "Content-Type: application/json" -d '{"name":"user2","email":"sample2@example.com","password":"pass"}' http://localhost:8080/users/update/1
```

users/delete

```
curl -v -H "Accept: application/json"  -H "Content-Type: application/json"  http://localhost:8080/users/delete/1
```

## TODO

todos/create

```

```
