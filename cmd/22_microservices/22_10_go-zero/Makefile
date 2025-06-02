# Makefile for Galactic Explorer API and Frontend
# Single command to generate code, transpile TypeScript, run backend, and serve frontend

# Variables
GOCTL := goctl
ESBUILD := esbuild
PROTO_FILE := explorer.api
BACKEND := explorer.go
FRONTEND_DIR := frontend
FRONTEND_SERVER := $(FRONTEND_DIR)/server.go
TS_CLIENT := $(FRONTEND_DIR)/src/client.ts
JS_CLIENT := $(FRONTEND_DIR)/src/client.js

# Single command to generate code, transpile, and run servers
run:
	@echo "Checking dependencies..."
	@if ! command -v $(GOCTL) >/dev/null 2>&1; then \
		echo "Error: goctl not found. Run 'go install github.com/zeromicro/go-zero/tools/goctl@latest'"; \
		exit 1; \
	fi
	@if ! command -v $(ESBUILD) >/dev/null 2>&1; then \
		echo "Error: esbuild not found. Run 'npm install -g esbuild'"; \
		exit 1; \
	fi
	@echo "Generating go-zero backend code..."
	$(GOCTL) api go -api $(PROTO_FILE) -dir .
	@echo "Generating TypeScript client code..."
	$(GOCTL) api ts -api $(PROTO_FILE) -dir $(FRONTEND_DIR)/src
	@echo "Transpiling TypeScript to JavaScript..."
	$(ESBUILD) $(TS_CLIENT) --outfile=$(JS_CLIENT) --format=iife --global-name=ClientModule
	@echo "Starting backend server..."
	go run $(BACKEND) &
	@echo "Starting frontend server..."
	go run $(FRONTEND_SERVER) &
	@echo "Servers running. Backend: http://localhost:8800, Frontend: http://localhost:3000"
	@wait

# Clean generated files (optional)
clean:
	@echo "Cleaning generated files..."
	rm -rf internal etc $(BACKEND) $(TS_CLIENT) $(JS_CLIENT)

.PHONY: run clean