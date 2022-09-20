# The makefile can quick execute command like test, build or other command.

.PHONY: test
test:
	go test ./...
