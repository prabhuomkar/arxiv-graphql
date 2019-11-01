packages = \
	./config \
	./resolver \

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
cover: test
	@echo "mode: count" > cover-all.out
	@$(foreach package,$(packages), \
		tail -n +2 $(package)/cover.out >> cover-all.out;)
	@gocover-cobertura < cover-all.out > cover-cobertura.xml

.PHONY: show
show:
	@echo Launching web browser to show overall coverage...
	@go tool cover -html=cover-all.out