# Makefile
# https://www.alexedwards.net/blog/a-time-saving-makefile-for-your-go-projects

# Variables
BUILD_DIR = ./bin
CMD_DIR = ./cmd
CMD_NAMES = ccwc ccwc_bubbletea ccwc_cobra

# Set default goal (when you run make without specifying a goal)
.DEFAULT_GOAL := help

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	# Throws error if not confirmed
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	# Check for uncommitted changes
	git diff --exit-code


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...


# ==================================================================================== #
# DEPENDENCIES
# ==================================================================================== #

## deps: download dependencies
.PHONY: deps
deps:
	go mod download


# ==================================================================================== #
# TESTING
# ==================================================================================== #

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## lint: run linter
.PHONY: lint
lint:
	golangci-lint run ./...


# ==================================================================================== #
# COMPILING
# ==================================================================================== #

## all: build and run all executables
.PHONY: all
all: build run

## build: build all executables
.PHONY: build
build:
	@echo "Building executables..."
	mkdir -p $(BUILD_DIR)
	@$(foreach cmd, $(CMD_NAMES), \
		go build -o $(BUILD_DIR)/$(cmd) $(CMD_DIR)/$(cmd);)

## build/ccwc: build ccwc executable
.PHONY: build/ccwc
build/ccwc:
	@echo "Building ccwc executable..."
	mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/ccwc $(CMD_DIR)/ccwc

## build/ccwc_cobra: build ccwc_cobra executable
.PHONY: build/ccwc_cobra
build/ccwc_cobra:
	@echo "Building ccwc_cobra executable..."
	mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/ccwc_cobra $(CMD_DIR)/ccwc_cobra

## build/ccwc_bubbletea: build ccwc_bubbletea executable
.PHONY: build/ccwc_bubbletea
build/ccwc_bubbletea:
	@echo "Building ccwc_bubbletea executable..."
	mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/ccwc_bubbletea $(CMD_DIR)/ccwc_bubbletea


# ==================================================================================== #
# RUNNING
# ==================================================================================== #

## run: run all executables
.PHONY: run
run: build run/ccwc run/ccwc_cobra run/ccwc_bubbletea

## run/ccwc: run ccwc implementation
.PHONY: run/ccwc
run/ccwc: build/ccwc
	@echo "Executing ccwc..."
	@$(BUILD_DIR)/ccwc

## run/ccwc_cobra: run ccwc cobra implementation
.PHONY: run/ccwc_cobra
run/ccwc_cobra: build/ccwc_cobra
	@echo "Executing ccwc cobra implementation..."
	@$(BUILD_DIR)/ccwc_cobra

## run/ccwc_bubbletea: run ccwc bubbletea implementation
.PHONY: run/ccwc_bubbletea
run/ccwc_bubbletea: build/ccwc_bubbletea
	@echo "Executing ccwc bubbletea implementation..."
	@$(BUILD_DIR)/ccwc_bubbletea


# ==================================================================================== #
# LIVE RELOADING
# ==================================================================================== #

## run/live/ccwc_cobra: run ccwc_cobra application with live reloading
.PHONY: run/live/ccwc_cobra
run/live/ccwc_cobra:
	go run github.com/cosmtrek/air@latest \
		--build.cmd "make build/ccwc_cobra" \
		--build.bin "$(BUILD_DIR)/ccwc_cobra" \
		--build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"

## run/live/ccwc: run ccwc application with live reloading
.PHONY: run/live/ccwc
run/live/ccwc:
	go run github.com/cosmtrek/air@latest \
		--build.cmd "make build/ccwc" \
		--build.bin "$(BUILD_DIR)/ccwc" \
		--build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"

## run/live/ccwc_bubbletea: run ccwc_bubbletea application with live reloading
.PHONY: run/live/ccwc_bubbletea
run/live/ccwc_bubbletea:
	go run github.com/cosmtrek/air@latest \
		--build.cmd "make build/ccwc_bubbletea" \
		--build.bin "$(BUILD_DIR)/ccwc_bubbletea" \
		--build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"


# ==================================================================================== #
# CLEANING
# ==================================================================================== #

## clean: clean up all build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up..."
	go clean
	@rm -rf $(BUILD_DIR)/*


# ==================================================================================== #
# GIT
# ==================================================================================== #

## push: push changes to the remote Git repository, after running quality control checks
.PHONY: push
push: tidy audit no-dirty
	git push