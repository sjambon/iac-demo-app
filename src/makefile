.PHONY: clean build

build: clean build-mailfunction build-relayfunction

build-relayfunction:
	cd relayfunction; go mod download; \
	GOOS=windows GOARCH=amd64 go build -o ../output/relayfunction/function.exe cmd/main.go
	cp relayfunction/host.json output/relayfunction/
	mkdir output/relayfunction/RelayHttpFunction && cp relayfunction/RelayHttpFunction/function.json output/relayfunction/RelayHttpFunction/function.json

build-mailfunction:
	cd mailfunction; go mod download; \
	GOOS=windows GOARCH=amd64 go build -o ../output/mailfunction/function.exe cmd/main.go
	cp mailfunction/host.json output/mailfunction/
	mkdir output/mailfunction/MailHttpFunction && cp mailfunction/MailHttpFunction/function.json output/mailfunction/MailHttpFunction/function.json
	mkdir output/mailfunction/MailEventFunction && cp mailfunction/MailEventFunction/function.json output/mailfunction/MailEventFunction/function.json

clean:
	rm -rf output
