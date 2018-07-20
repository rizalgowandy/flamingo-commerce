CONTEXT?=dev
.PHONY: up localup update test

up:
	rm -rf vendor/
	dep ensure -v

update:
	rm -rf vendor/
	dep ensure -v -update flamingo.me/flamingo

localup: up
	rm -rf vendor/flamingo.me/flamingo
	ln -sf ../../../flamingo vendor/flamingo.me/flamingo
	rm -rf vendor/flamingo.me/flamingo/vendor
	
test:
	go test ./...