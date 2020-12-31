
SIGNUP:
curl --header "Content-Type: application/json" --request POST --data '{"user":"apascuaslco@gmail.com","password":"12345"}' -v http://localhost:8080/api/v1/signup

LOGIN:
curl --header "Content-Type: application/json" --request POST --data '{"user":"apascuaslco@gmail.com","password":"12345"}' -v http://localhost:8080/api/v1/login

TOKEN HEADER:
curl -H "Content-Type; application/json" -H "Authorization: Bearer TOKE-LOGIN" --request GET -v http://localhost:8080/api/v1/hello
