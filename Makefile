ROOT := $(CURDIR)
GOPKGS = \
		golang.org/x/tools/cmd/cover

default: test

deps:
	@echo "[Deps] installing dependencies"
	@go get -v $(GOPKGS)

vet:
	@echo "[Vet] running go vet"
	@cd ${ROOT}/hll && go vet

ci: deps vet test

test:
	@echo "[Test] running tests"
	@go test -v ./... -cover -coverprofile=c.out -bench=".*"

.PHONY: default deps vet ci test
