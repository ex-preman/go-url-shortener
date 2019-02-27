# url-shortener-go
URL Shortener in Go

- Change db connection
- Execute db.sql
- go run main.go

Endpoint :
- Shorten
URL: localhost:8080/shorten

JSON Payload:
`{
	"url": "https://tutorialedge.net/golang/golang-mysql-tutorial/#performing-basic-sql-commands"
}`

- Redirect
URL: http://localhost:8080/[code from response shorten]
