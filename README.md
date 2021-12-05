# TODO_APP by golang React Typescript

## API

### USER

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

### TODO

todos/create

```
 curl -v -H "Accept: application/json"  -H "Content-Type: application/json" -d '{"user_id":3,"content":"test_todo_1"}' http://localhost:8080/todos/new
```

todos/show

```
curl -v -H "Accept: application/json"  -H "Content-Type: application/json" http://localhost:8080/todos/show/1
```

todos/update

```
curl -v -H "Accept: application/json"  -H "Content-Type: application/json" -d '{"user_id":3,"content":"update_content"}' http://localhost:8080/todos/update/1
```
