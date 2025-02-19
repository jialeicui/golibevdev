run-example:
	@cd example && go build && ./example

gen:
	@rm *_string.go
	@go generate
