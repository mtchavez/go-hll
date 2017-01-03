ROOT := $(CURDIR)

test:
	cd ${ROOT}/hll && go test -cover
