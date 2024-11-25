# url-shortener-go
URL Shortener in Go

- Change db connection
- Execute db.sql
- go run main.go

Endpoint :
- Shorten
```
URL: localhost:8080/shorten
JSON Payload:
{
	"url": "https://go.dev/play"
}
```

- Curl example:
curl -d '{"url":"https://go.dev/play"}' -H "Content-Type: application/json" -X POST http://localhost:8080/shorten

- Redirect
URL: http://localhost:8080/{shorten_code}
