DSN="host=localhost port=5432 user=postgres password=password dbname=vueapi sslmode=disable timezone=UTC connect_timeout=5"
BINARY_NAME=vueapi
BINARY_FILE_PATH=./dist/${BINARY_NAME}

## build: Build binary
build:
	@echo "[+] Building web api server."
	@go build -o ${BINARY_FILE_PATH} ./cmd/api/
	@echo "[+] Created the output file in ${BINARY_FILE_PATH}"

## run: builds and runs the application
run: build
	@echo "[+] Starting web api server."
	@env DSN=${DSN} ${BINARY_FILE_PATH} &
	@sleep 1
	@echo "[+] Web api server was started in the background."

## clean: runs go clean and deletes binaries
clean:
	@echo "[+] Cleanup the directory."
	@go clean
	@rm -rf ${BINARY_FILE_PATH}
	@echo "[+] Cleanup was successful."

## start: an alias to run
start: run

## stop: stops the running application
stop:
	@echo "[+] Stoping web api server."
	@-pkill -SIGTERM -f "${BINARY_FILE_PATH}"
	@echo "[+] Web api server was stopped."

## restart: stops and starts the running application
restart: stop start