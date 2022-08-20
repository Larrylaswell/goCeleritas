BINARY_NAME=celaritas

## build: builds all binaries
build:
	@go echo "Building Celaritas.exe"
    @go build -o -x tmp/${BINARY_NAME}.exe 
    @echo "Bricks built!"

run: build
    @echo "Starting Bricks..."
    @.\tmp\${BINARY_NAME} &
    @echo "Bricks started!"

clean:
    @echo "Cleaning..."
    @go clean
    @del .\tmp\${BINARY_NAME}
    @echo "Cleaned!"

test:
    @echo "Testing..."
    @go test ./...
    @echo "Done!"

start: run