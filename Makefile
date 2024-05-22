.PHONY:	build

generate:
	@go generate ./parser/document.go


build:
	@go generate ./parser/document.go
	go build -o rexplorer ./cmd
