BINARY_NAME=aggrss

build:
	@go mod vendor
	@echo "Building Aggrss..."
	@go build -o ${BINARY_NAME} .
	@echo "Aggrss built!"

run: build
	@echo "Starting Aggrss..."
	@./${BINARY_NAME} &
	@echo "Aggrss started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned!"

start: run

stop:
	@echo "Stopping Aggrss..."
	@-pkill -SIGTERM -f "${BINARY_NAME}"
	@echo "Stopped Aggrss!"

restart: stop start

# export GOPATH=$HOME/devops-gallery/dev-env/go-lab
# export PATH=$PATH:/usr/share/go:$GOPATH/bin