.PHONY: release dist clean dist-N

OS := darwin linux
VERSION := $(shell go run main.go --version | awk '{print $$2}')

release:
	git commit -m $(VERSION)
	git tag -a v$(VERSION) -m $(VERSION)
	git push origin v$(VERSION)
	git push origin master

dist:
	@set -e; \
	for os in $(OS); do \
		script/build.sh $$os amd64 $(VERSION); \
	done
