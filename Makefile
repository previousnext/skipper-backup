#!/usr/bin/make -f

build:
	./hack/build.sh linux server skipper-backup github.com/previousnext/skipper-backup

test:
	./hack/test.sh github.com/previousnext/skipper-backup/...
