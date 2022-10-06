# go-api-macaroons
shell command: `go run main.go to run`
hit endpoints with insomnia or postman

example:
POST: http://localhost:10000/macaroon
JSON body:
{
	"Location":"www.test.com/two",
	"Id":"3ad9e155-999c-4b22-a3b1-a49ba26e0707",
	"Secret":"verifyMe"
}

should return JOSN formatted macaroon.
```
cURL for shell
curl --request POST \
  --url http://localhost:10000/macaroon \
  --header 'Content-Type: application/json' \
  --data '{
	"Location":"www.test.com/two",
	"Id":"3ad9e155-999c-4b22-a3b1-a49ba26e0707",
	"Secret":"verifyMe"
}'
```