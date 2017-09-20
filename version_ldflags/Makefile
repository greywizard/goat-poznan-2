DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)

M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: run
run: ; 
	$(info $(M) running script…) @ ## Run main script
	@go run \
		-ldflags '-X main.Version=$(VERSION) -X main.BuildTime=$(DATE)' \
		program.go


		