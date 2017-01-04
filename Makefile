ROOT := $(CURDIR)

default: test

lint:
	@echo "[Lint] running golint"
	@cd ${ROOT}/hll && golint -set_exit_status

vet:
	@echo "[Vet] running go vet"
	@cd ${ROOT}/hll && go vet

ci: vet lint test

test:
	@echo "[Test] running tests"
	@cd ${ROOT}/hll && go test -cover

.PHONY: default golint test
