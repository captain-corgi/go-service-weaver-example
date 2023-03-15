run:
	go run cmd/local/main.go

tidy:
	go mod tidy
	go mod vendor

sample:
	curl "localhost:12345/hello?name=Anh%20Le"

wstatus:
	weaver single status
wdashboard:
	weaver single dashboard