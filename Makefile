.PHONY: build clean docker-build docker-dev

# App settings
APP_NAME=rick-and-morty-htmx
IMAGE_NAME=rick-and-morty-htmx

# Directories
DIST_DIR=dist
CMD_DIR=cmd/website

# Tailwind CSS build command
TAILWIND_BUILD_CMD=npm run build:css

build:
	# Clean up previous builds
	@echo "Cleaning up..."
	@mkdir -p $(DIST_DIR)
	@rm -rf $(DIST_DIR)/*
# Generate Tailwind CSS
	@echo "Generating Tailwind CSS..."
	@$(TAILWIND_BUILD_CMD)
# Build the Go app
	@echo "Building Go app..."
	@go build -o $(DIST_DIR)/$(APP_NAME) $(CMD_DIR)/main.go

	# Copy frontend assets
	@echo "Copying frontend assets..."
	@cp -R assets $(DIST_DIR)/

docker-build: build
	@echo "Building Docker image..."
	@docker build -t $(IMAGE_NAME) .

docker-dev:
	@docker-compose up

clean:
	@rm -rf $(DIST_DIR)
	@docker rmi $(IMAGE_NAME)
