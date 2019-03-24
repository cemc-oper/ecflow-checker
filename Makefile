.PHONY: ecflow_checker

ecflow_checker:
	go build -o bin/ecflow_checker ecflow_checker/main.go