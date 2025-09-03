# Makefile for Go Program Study Project

# 项目配置
BINARY_NAME := go_study
MODULE_NAME := github.com/GooglerLi/goProgramStudy
MAIN_FILE := main.go

# Go 相关配置
GO := go
GO_BUILD := $(GO) build
GO_TEST := $(GO) test
GO_CLEAN := $(GO) clean
GO_MOD_TIDY := $(GO) mod tidy
GOFMT := gofmt

# 构建标志
LDFLAGS := -w -s
BUILD_FLAGS := -ldflags "$(LDFLAGS)"

# 平台架构配置
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
BUILD_DIR := build
DIST_DIR := dist

# 默认目标
.PHONY: all
all: build

# 构建项目
.PHONY: build
build: tidy
	@echo "Building $(BINARY_NAME) for $(GOOS)/$(GOARCH)..."
	$(GO_BUILD) $(BUILD_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "Build completed: $(BUILD_DIR)/$(BINARY_NAME)"

# 安装依赖并整理模块
.PHONY: tidy
tidy:
	$(GO_MOD_TIDY)

# 运行程序
.PHONY: run
run: tidy
	$(GO) run $(MAIN_FILE)

# 清理构建文件
.PHONY: clean
clean:
	@echo "Cleaning build files..."
	$(GO_CLEAN)
	rm -rf $(BUILD_DIR) $(DIST_DIR)
	@echo "Clean completed"

# 代码格式化
.PHONY: fmt
fmt:
	$(GOFMT) -w .

# 运行测试
.PHONY: test
test:
	$(GO_TEST) -v ./...

# 显示帮助信息
.PHONY: help
help:
	@echo "Makefile for Go Program Study Project"
	@echo ""
	@echo "Usage:"
	@echo "  make build     - 构建项目，生成可执行文件"
	@echo "  make run       - 运行项目"
	@echo "  make test      - 运行测试"
	@echo "  make tidy      - 整理依赖"
	@echo "  make clean     - 清理构建文件"
	@echo "  make fmt       - 格式化代码"
	@echo "  make help      - 显示帮助信息"
	@echo ""
	@echo "高级构建:"
	@echo "  make build-linux   - 构建 Linux 版本"
	@echo "  make build-darwin  - 构建 macOS 版本"
	@echo "  make build-windows - 构建 Windows 版本"
	@echo "  make build-all     - 构建所有平台版本"

# 多平台构建
.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 $(MAKE) build

.PHONY: build-darwin
build-darwin:
	GOOS=darwin GOARCH=amd64 $(MAKE) build

.PHONY: build-darwin-arm64
build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 $(MAKE) build

.PHONY: build-windows
build-windows:
	GOOS=windows GOARCH=amd64 $(MAKE) build

# 构建所有平台
.PHONY: build-all
build-all: build-linux build-darwin build-darwin-arm64 build-windows

# 安装到系统路径 (需要 sudo 权限)
.PHONY: install
install: build
	@echo "Installing $(BINARY_NAME) to /usr/local/bin..."
	sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "Installation completed"

# 卸载
.PHONY: uninstall
uninstall:
	@echo "Uninstalling $(BINARY_NAME)..."
	sudo rm -f /usr/local/bin/$(BINARY_NAME)
	@echo "Uninstall completed"

# 显示版本信息
.PHONY: version
version:
	@echo "Module: $(MODULE_NAME)"
	@echo "Binary: $(BINARY_NAME)"
	@echo "Go version: $(shell go version)"
	@echo "Platform: $(GOOS)/$(GOARCH)"