TIMESTAMP  	:= $(shell /bin/date "+%F %T")

fmt:
	@go fmt ./...

github: fmt
	@git add . &> /dev/null
	@git commit -m "$(TIMESTAMP)"
	@git push origin master

.PHONY: fmt github
