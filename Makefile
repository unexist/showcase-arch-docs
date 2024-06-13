.DEFAULT_GOAL := build
.ONESHELL:

EXE_GIN := todo-service.bin
EXE_STRUCT := structurizr.bin

define JSON_TODO
curl -X 'POST' \
  'http://localhost:8080/todo' \
  -H 'accept: */*' \
  -H 'Content-Type: application/json' \
  -d '{
  "description": "string",
  "done": true,
  "title": "string"
}'
endef
export JSON_TODO
# Tools
todo:
	@echo $$JSON_TODO | bash

list:
	@curl -X 'GET' 'http://localhost:8080/todo' -H 'accept: */*' | jq .

swagger:
	@$(SHELL) -c "cd todo-service-gin; swag init"

open-swagger:
	open http://localhost:8080/swagger/index.html

c4:
	@$(SHELL) -c "cd todo-service-gin/tools; ./$(EXE_STRUCT) && plantuml *.plantuml"

# Build
build-struct:
	@$(SHELL) -c "cd todo-service-gin/tools; GO111MODULE=on; go mod download; go build -o $(EXE_STRUCT)"

build-gin:
	@$(SHELL) -c "cd todo-service-gin; GO111MODULE=on; go mod download; go build -o $(EXE_GIN)"

# Analysis
vet-gin:
	@$(SHELL) -c "cd todo-service-gin; go vet"

# Run
run-gin:
	@$(SHELL) -c "cd todo-service-gin; ./$(EXE_GIN)"

# Tests
test-gin:
	@$(SHELL) -c "cd todo-service-gin; go test -v ./test"

# Helper
clear:
	rm -rf todo-service-gin/$(EXE_GIN)

install:
	go install braces.dev/errtrace/cmd/errtrace@latest
	go install golang.org/x/tools/cmd/deadcode@latest
	go install goa.design/model/cmd/mdl@latest
	go install goa.design/model/cmd/stz@latest
