# The makefile can quick execute command like test, build or other command.

.PHONY: help
help:
	@echo "make [Command] {Args}"
	@echo ""
	@echo "Commands:"
	@echo "- test: to test project by unit-test for go code."

# test: to test all code of project by unit-test. It will run 'go test ./...', and please run this command at repo root.
.PHONY: test
test:
	go test ./...
