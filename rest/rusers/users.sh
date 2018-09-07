curl -i -H 'Content-Type: application/json' \
    -d '{"Name":"Antoine"}' http://127.0.0.1:8080/users
curl -i http://127.0.0.1:8080/users/0
curl -i -X PUT -H 'Content-Type: application/json' \
    -d '{"Name":"Antoine Imbert"}' http://127.0.0.1:8080/users/0
curl -i -X DELETE http://127.0.0.1:8080/users/0
curl -i http://127.0.0.1:8080/users
