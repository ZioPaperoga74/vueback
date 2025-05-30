SHELL=cmd.exe
DSN=host=localhost port=5432 user=postgres password=password dbname=vueapi sslmode=disable timezone=UTC connect_timeout=5
BINARY_NAME=vueapi.exe

## build: builds all binaries
build:
	@go build -o ${BINARY_NAME} ./cmd/api
	@echo back end built!

run: build
	@echo Starting back end...
	set DSN=${DSN}&& start /B .\${BINARY_NAME} &
	@echo back end started!

clean:
	@echo Cleaning...
	@DEL ${BINARY_NAME}
	@go clean
	@echo Cleaned!

start: run

stop:
	@echo "Stopping the back end..."
	@taskkill /IM ${BINARY_NAME} /F
	@echo Stopped back end

restart: stop start
