PROJECT_PATH := $(shell pwd)
PROJECT_NAME ?= $(shell basename $(PROJECT_PATH))
PROJECT_VERSION ?= $(shell cat VERSION)
PROJECT_ARCH ?= "amd64"
RELEASE_FROM ?= "docker"
BASE_PATH ?= @pwd
CONFIG_PATH ?= "$(PROJECT_PATH)/config.yml"

## 打印信息
info:
	@echo "PROJECT PATH: $(PROJECT_PATH)"
	@echo "PROJECT NAME: $(PROJECT_NAME)"
	@echo "PROJECT VERSION: $(PROJECT_VERSION)"
	@echo "PROJECT ARCH: $(PROJECT_ARCH)"
	@echo "RELEASE FROM: $(RELEASE_FROM)"

## 格式检查
lint:
	@echo "格式检查开始"
	@cd server
	@golangci-lint run --modules-download-mode=vendor
	@cd ${BASE_PATH}
	@echo "格式检查完成"

### 测试
test:
	@echo "配置文件: $(CONFIG_PATH)"
	@echo "测试开始"
	@CONFIG_PATH=$(CONFIG_PATH) go test -p 1 -v -failfast -short -mod=vendor -race -coverprofile="bin/cover.out" ./...|tee bin/ut.tmp
	@tar -zcf bin/unitTestReport.tar.gz bin/ut.tmp
	@grep -E '===|---' bin/ut.tmp > bin/utStatistics.tmp
	@echo "测试完成"
	@echo "计算覆盖率"
	@go tool cover -func="bin/cover.out"|tee bin/ut_coverage.tmp
	@echo "计算覆盖率完成"

## 清理
clear:
	@echo "清理开始"
	@rm -rf bin/* || /bin/true
	@echo "清理完成"

## 编译

build:
	@echo "编译开始"
	@make info
	@make clear
	@version=`cat VERSION`
	@GOOS=linux GOARCH=amd64 go build -flags="-w -s -X system.Version=${PROJECT_VERSION}"
	@CGO_ENABLED=0 GOOS=linux GOARCH=$(PROJECT_ARCH) go build -mod=vendor -v -o bin/$(PROJECT_NAME)
	@CGO_ENABLED=0 GOOS=linux GOARCH=$(PROJECT_ARCH) go build -mod=vendor -v -o bin/config_tools ymer/tools/config_tools
	@CGO_ENABLED=0 GOOS=linux GOARCH=$(PROJECT_ARCH) go build -mod=vendor -v -o bin/pg2es ymer/tools/pg2es
	@CGO_ENABLED=0 GOOS=linux GOARCH=$(PROJECT_ARCH) go build -mod=vendor -v -o bin/user_log_update ymer/tools/user_log_update
	@echo "编译完成"

## 编译
build_arm:
	@make info
	@make clear
	@echo "编译开始"
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -mod=vendor -v -o bin/$(PROJECT_NAME)
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -mod=vendor -v -o bin/config_tools ymer/tools/config_tools
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -mod=vendor -v -o bin/pg2es ymer/tools/pg2es
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -mod=vendor -v -o bin/user_log_update ymer/tools/user_log_update
	@echo "编译完成"

## 发布
release:
ifeq ($(RELEASE_FROM),"docker")
	@make release_docker
else
	@make release_binary
endif

## 容器发布
release_docker:
	@make build
	@echo "容器发布开始"
	@cp config.yml bin/config.yml
	@docker build --build-arg path=bin --build-arg name=$(PROJECT_NAME) --tag $(PROJECT_REPOSITORY):$(PROJECT_VERSION) .
	@docker push $(PROJECT_REPOSITORY):$(PROJECT_VERSION)
	@echo "容器发布完成"

## 二进制发布
release_binary:
	@make build
	@echo "二进制发布开始"
	@mkdir -p bin/$(PROJECT_NAME)-$(PROJECT_VERSION)
	@mv bin/$(PROJECT_NAME) bin/$(PROJECT_NAME)-$(PROJECT_VERSION)
	@mv bin/config_tools bin/$(PROJECT_NAME)-$(PROJECT_VERSION)
	@mv bin/pg2es bin/$(PROJECT_NAME)-$(PROJECT_VERSION)
	@mv bin/user_log_update bin/$(PROJECT_NAME)-$(PROJECT_VERSION)
	@cp config.yml bin/$(PROJECT_NAME)-$(PROJECT_VERSION)/config.yml
	@cp sv_start.sh bin/$(PROJECT_NAME)-$(PROJECT_VERSION)/sv_start.sh
	@cp docker-compose.yml bin/$(PROJECT_NAME)-$(PROJECT_VERSION)/docker-compose.yml
	@cp -R dict bin/$(PROJECT_NAME)-$(PROJECT_VERSION)/dict
	@mkdir -p bin/$(PROJECT_NAME)
	@mv bin/$(PROJECT_NAME)-$(PROJECT_VERSION) bin/$(PROJECT_NAME)
	@ln -s $(PROJECT_NAME)-$(PROJECT_VERSION) bin/$(PROJECT_NAME)/latest
	@cd bin && tar -zcvf $(PROJECT_NAME)-$(PROJECT_VERSION).tar.gz $(PROJECT_NAME)/
	@echo "二进制发布完成"

## arm二进制发布
release_armbinary:
	@make build_arm
	@echo "arm二进制发布开始"
	@mkdir -p bin/$(PROJECT_NAME)-$(PROJECT_VERSION)
	@mv bin/$(PROJECT_NAME) bin/$(PROJECT_NAME)-$(PROJECT_VERSION)
	@mv bin/config_tools bin/$(PROJECT_NAME)-$(PROJECT_VERSION)
	@mv bin/pg2es bin/$(PROJECT_NAME)-$(PROJECT_VERSION)
	@mv bin/user_log_update bin/$(PROJECT_NAME)-$(PROJECT_VERSION)
	@cp config.yml bin/$(PROJECT_NAME)-$(PROJECT_VERSION)/config.yml
	@cp sv_start.sh bin/$(PROJECT_NAME)-$(PROJECT_VERSION)/sv_start.sh
	@cp docker-compose.yml bin/$(PROJECT_NAME)-$(PROJECT_VERSION)/docker-compose.yml
	@cp -R dict bin/$(PROJECT_NAME)-$(PROJECT_VERSION)/dict
	@mkdir -p bin/$(PROJECT_NAME)
	@mv bin/$(PROJECT_NAME)-$(PROJECT_VERSION) bin/$(PROJECT_NAME)
	@ln -s $(PROJECT_NAME)-$(PROJECT_VERSION) bin/$(PROJECT_NAME)/latest
	@cd bin && tar -zcvf $(PROJECT_NAME)-$(PROJECT_VERSION).tar.gz $(PROJECT_NAME)/
	@echo "arm二进制发布完成"
