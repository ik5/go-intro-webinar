info:
	@echo 'Go Info:'
	@go env
	@echo ''
	@echo "Go mod Info:"
	@GO111MODULE=on go mod graph
	@echo ''
	@echo "Git files Info:"
	@git ls-files --stage --abbrev=8 -v --others -c -d --eol --full-name -k -m -u

test: update-dep
	go test -v -a -coverpkg=all ./...

test-short: update-dep
	go test -v -a -coverpkg=all -short ./...

build: update-dep
	go build cmd/main
	go build cmd/users

update-dep:
	GO111MODULE=on go mod tidy -v
