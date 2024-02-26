#全局设置环境变量
export GOPROXY=https://goproxy.cn,direct
# 定义变量
ifndef GOPATH
	GOPATH := $(shell go env GOPATH)
endif

GOBIN=$(GOPATH)/bin
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) mod tidy
HZ=$(GOBIN)/goctl

# 安装goctl代码生成工具
$(shell if [ ! -d $(HZ) ]; then \
	$(GOCMD) install github.com/zeromicro/go-zero/tools/goctl@latest; \
fi; \
)

MYSQL_INFO="root:oMbPi5munxCsBSsiLoPV@tcp(110.41.179.89:3306)/hertzdb"

all: deps build ## 默认的构建目标

clean: ## 清理目标
	$(GOCLEAN)
	rm -rf target

deps: ## 安装依赖目标
	$(GOGET) -v

build: ## 构建目标
	$(GOBUILD) -o target/zero-admin api/admin.go

start: build stop ## 运行目标
	nohup ./target/zero-admin -f ./api/etc/admin.yaml > /dev/null 2>&1 &

stop: ## 停止目标
	-pkill -f zero-admin

restart: stop start ## 重启项目

.DEFAULT_GOAL := all ## 默认构建目标是

format: ## 格式化代码
	$(HZ) api format --dir api/doc/system
proto:	## 生成代码
	$(HZ) api go -api api/doc/admin.api -dir ./api


model: ## 生成dao
	$(HZ) model mysql datasource -url=$(MYSQL_INFO) -table="*"  -dir="./api/internal/model" -c

image: ## 构建docker镜像
	docker build -t zero-admin:0.0.1 -f Dockerfile .

docker: image ## 启动docker容器
	docker run -itd --net=host --name=zero-admin zero-admin:0.0.1; \


deployment: image ## 部署k8s容器
	kubectl apply -f script/deployment.yaml; \

help: ## show help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[$$()% 0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
