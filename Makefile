packages = \
	./gql \

.PHONY: install
install: 
	@echo Installing dependencies...
	@go get ./...

.PHONY: build
build:
	@echo Building binary...
	@go build -v .

.PHONY: run
run: build
	@echo Running Droll GraphQL API...
	@if [ $(go env GOOS) == "windows" ]
		@droll-api.exe
	@else
		@./droll-api
	@fi

.PHONY: test
test:
	@$(foreach package,$(packages), \
		set -e; \
		go test -coverprofile $(package)/cover.out -covermode=count $(package);)

clear:
	@echo Cleaning all builds and test output...
	@$(foreach package,$(packages), \
		rm -rf *.out; \
		rm -rf *.xml; \
		rm -rf *.exe; \
		rm -rf ./**/*.out; \
		rm -rf ./**/*.xml;)

.PHONY: cover
cover:
	@set -e
	@echo "" > coverage.txt
	@$(foreach package,$(packages), \
	    go test -race -coverprofile=profile.out -covermode=atomic $(package); cat profile.out >> coverage.txt)
	@rm profile.out;

.PHONY: show
show:
	@echo Launching web browser to show overall coverage...
	@go tool cover -html=cover-all.out

