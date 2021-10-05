NAME = "dnsctl"
VERSION = v0.0.1
ALIYUN_VERSION = v0.0.1
TENCENT_VERSION = v0.0.1
DATE = $(shell date '+%Y%m%d')
SCRIPT_PATH = $(shell cd "$(dirname "$0")"; pwd)
BIN_DIR = ${SCRIPT_PATH}/bin
PKG = github.com/chinatree/dnsctl

PREFIX=.
DESTDIR=
GOFLAGS=
BINDIR=${PREFIX}/bin
COVDIR=${PREFIX}/cov

OS =
ARCH = amd64
CLIS = dnsctl-aliyun
LIBS =

all: lint $(CLIS)

.PHONY: lint
lint:
	@echo "Code specification inspection ..."
	for pkg in $$(go list ./... | grep -v /vendor/) ; do \
		golint $$pkg ; \
	done
	@echo "Code specification inspection [OK]"
	goimports -l -w `find . -type f -name '*.go'`
	errcheck -verbose -ignoretests -blank `go list ./...`
	go vet -v `go list ./...`
	go fmt `go list ./...`



# 定义目标变量开始
%amd64: GOARCH := amd64
%arm64: GOARCH := arm64
# 定义目标变量结束



.PHONY: build-in-docker
build-in-docker: build-in-docker-amd64 build-in-docker-arm64

.PHONY: build-in-docker-amd64
build-in-docker-amd64: build-in-docker-dnsctl-amd64 build-in-docker-dnsctl-aliyun-amd64 build-in-docker-dnsctl-tencent-amd64

.PHONY: build-in-docker-arm64
build-in-docker-arm64: build-in-docker-dnsctl-arm64 build-in-docker-dnsctl-aliyun-arm64 build-in-docker-dnsctl-tencent-arm64

.PHONY: build-in-docker-dnsctl
build-in-docker-dnsctl: build-in-docker-dnsctl-amd64 build-in-docker-dnsctl-arm64

.PHONY: build-in-docker-dnsctl-aliyun
build-in-docker-dnsctl-aliyun: build-in-docker-dnsctl-aliyun-amd64 build-in-docker-dnsctl-aliyun-arm64

.PHONY: build-in-docker-dnsctl-tencent
build-in-docker-dnsctl-tencent: build-in-docker-dnsctl-tencent-amd64 build-in-docker-dnsctl-tencent-arm64

.PHONY: build-in-docker-dnsctl-amd64
build-in-docker-dnsctl-amd64:
	@docker run --rm \
		--name dnsctl-${GOARCH} \
		--platform=linux/${GOARCH} \
		-v ${GITHUB_GOPATH}:/src/github.com \
		-w /src/github.com/chinatree/dnsctl \
		chinatree/golangcompile:1.16.6-alpine \
		sh -c "make build-dnsctl-${GOARCH}"

.PHONY: build-in-docker-dnsctl-arm64
build-in-docker-dnsctl-arm64:
	@docker run --rm \
		--name dnsctl-${GOARCH} \
		--platform=linux/${GOARCH} \
		-v ${GITHUB_GOPATH}:/src/github.com \
		-w /src/github.com/chinatree/dnsctl \
		chinatree/golangcompile:1.16.6-alpine \
		sh -c "make build-dnsctl-${GOARCH}"

.PHONY: build-in-docker-dnsctl-aliyun-amd64
build-in-docker-dnsctl-aliyun-amd64:
	@docker run --rm \
		--name dnsctl-aliyun-${GOARCH} \
		--platform=linux/${GOARCH} \
		-v ${GITHUB_GOPATH}:/src/github.com \
		-w /src/github.com/chinatree/dnsctl \
		chinatree/golangcompile:1.16.6-alpine \
		sh -c "make build-dnsctl-aliyun-${GOARCH}"

.PHONY: build-in-docker-dnsctl-aliyun-arm64
build-in-docker-dnsctl-aliyun-arm64:
	@docker run --rm \
		--name dnsctl-aliyun-${GOARCH} \
		--platform=linux/${GOARCH} \
		-v ${GITHUB_GOPATH}:/src/github.com \
		-w /src/github.com/chinatree/dnsctl \
		chinatree/golangcompile:1.16.6-alpine \
		sh -c "make build-dnsctl-aliyun-${GOARCH}"

.PHONY: build-in-docker-dnsctl-tencent-amd64
build-in-docker-dnsctl-tencent-amd64:
	@docker run --rm \
		--name dnsctl-tencent-${GOARCH} \
		--platform=linux/${GOARCH} \
		-v ${GITHUB_GOPATH}:/src/github.com \
		-w /src/github.com/chinatree/dnsctl \
		chinatree/golangcompile:1.16.6-alpine \
		sh -c "make build-dnsctl-tencent-${GOARCH}"

.PHONY: build-in-docker-dnsctl-tencent-arm64
build-in-docker-dnsctl-tencent-arm64:
	@docker run --rm \
		--name dnsctl-tencent-${GOARCH} \
		--platform=linux/${GOARCH} \
		-v ${GITHUB_GOPATH}:/src/github.com \
		-w /src/github.com/chinatree/dnsctl \
		chinatree/golangcompile:1.16.6-alpine \
		sh -c "make build-dnsctl-tencent-${GOARCH}"



.PHONY: build
build: build-amd64 build-arm64

.PHONY: build-amd64
build-amd64: build-dnsctl-amd64 build-dnsctl-aliyun-amd64 build-dnsctl-tencent-amd64

.PHONY: build-arm64
build-arm64: build-dnsctl-arm64 build-dnsctl-aliyun-arm64 build-dnsctl-tencent-arm64

.PHONY: build-dnsctl
build-dnsctl: build-dnsctl-amd64 build-dnsctl-arm64

.PHONY: build-dnsctl-aliyun
build-dnsctl-aliyun: build-dnsctl-aliyun-amd64 build-dnsctl-aliyun-arm64

.PHONY: build-dnsctl-tencent
build-dnsctl-tencent: build-dnsctl-tencent-amd64 build-dnsctl-tencent-arm64

build-dnsctl-amd64: PACKAGE_NAME := dnsctl
build-dnsctl-amd64: GOOS := linux
build-dnsctl-amd64: CGO_ENABLED := 0
build-dnsctl-amd64:
	@cd ${PREFIX} \
	&& CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -o bin/${PACKAGE_NAME}-${GOOS}-${GOARCH} ${PKG}/cmd/dnsctl

build-dnsctl-arm64: PACKAGE_NAME := dnsctl
build-dnsctl-arm64: GOOS := linux
build-dnsctl-arm64: CGO_ENABLED := 0
build-dnsctl-arm64:
	@cd ${PREFIX} \
	&& CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -o bin/${PACKAGE_NAME}-${GOOS}-${GOARCH} ${PKG}/cmd/dnsctl

build-dnsctl-aliyun-amd64: PACKAGE_NAME := dnsctl-aliyun
build-dnsctl-aliyun-amd64: GOOS := linux
build-dnsctl-aliyun-amd64: CGO_ENABLED := 0
build-dnsctl-aliyun-amd64:
	@cd ${PREFIX} \
	&& CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -o bin/${PACKAGE_NAME}-${GOOS}-${GOARCH} ${PKG}/cmd/aliyun

build-dnsctl-aliyun-arm64: PACKAGE_NAME := dnsctl-aliyun
build-dnsctl-aliyun-arm64: GOOS := linux
build-dnsctl-aliyun-arm64: CGO_ENABLED := 0
build-dnsctl-aliyun-arm64:
	@cd ${PREFIX} \
	&& CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -o bin/${PACKAGE_NAME}-${GOOS}-${GOARCH} ${PKG}/cmd/aliyun

build-dnsctl-tencent-amd64: PACKAGE_NAME := dnsctl-tencent
build-dnsctl-tencent-amd64: GOOS := linux
build-dnsctl-tencent-amd64: CGO_ENABLED := 0
build-dnsctl-tencent-amd64:
	@cd ${PREFIX} \
	&& CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -o bin/${PACKAGE_NAME}-${GOOS}-${GOARCH} ${PKG}/cmd/tencent

build-dnsctl-tencent-arm64: PACKAGE_NAME := dnsctl-tencent
build-dnsctl-tencent-arm64: GOOS := linux
build-dnsctl-tencent-arm64: CGO_ENABLED := 0
build-dnsctl-tencent-arm64:
	@cd ${PREFIX} \
	&& CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -o bin/${PACKAGE_NAME}-${GOOS}-${GOARCH} ${PKG}/cmd/tencent



.PHONY: package
package: package-amd64 package-arm64

.PHONY: package-amd64
package-amd64: package-dnsctl-amd64 package-dnsctl-aliyun-amd64 package-dnsctl-tencent-amd64

.PHONY: package-arm64
package-arm64: package-dnsctl-arm64 package-dnsctl-aliyun-arm64 package-dnsctl-tencent-arm64

.PHONY: package-dnsctl
package-dnsctl: package-dnsctl-amd64 package-dnsctl-arm64

.PHONY: package-dnsctl-aliyun
package-dnsctl-aliyun: package-dnsctl-aliyun-amd64 package-dnsctl-aliyun-arm64

.PHONY: package-dnsctl-tencent
package-dnsctl-tencent: package-dnsctl-tencent-amd64 package-dnsctl-tencent-arm64

package-dnsctl-amd64: PACKAGE_NAME := dnsctl
package-dnsctl-amd64: GOOS := linux
package-dnsctl-amd64:
	@ cd ${BINDIR} \
	&& mkdir -p ${PACKAGE_NAME}-${VERSION} \
	&& cp ${PACKAGE_NAME}-${GOOS}-${GOARCH} ${PACKAGE_NAME}-${VERSION} \
	&& cd ${PACKAGE_NAME}-${VERSION} \
	&& ln -s ${PACKAGE_NAME}-${GOOS}-${GOARCH} ${PACKAGE_NAME} \
	&& cd .. \
	&& tar czf ${PACKAGE_NAME}-${VERSION}-${GOARCH}.tar.gz ${PACKAGE_NAME}-${VERSION} \
	&& tar czf ${PACKAGE_NAME}-${VERSION}-${GOARCH}-${DATE}.tar.gz ${PACKAGE_NAME}-${VERSION} \
	&& rm -rf ${PACKAGE_NAME}-${VERSION}

package-dnsctl-arm64: PACKAGE_NAME := dnsctl
package-dnsctl-arm64: GOOS := linux
package-dnsctl-arm64:
	@ cd ${BINDIR} \
	&& mkdir -p ${PACKAGE_NAME}-${VERSION} \
	&& cp ${PACKAGE_NAME}-${GOOS}-${GOARCH} ${PACKAGE_NAME}-${VERSION} \
	&& cd ${PACKAGE_NAME}-${VERSION} \
	&& ln -s ${PACKAGE_NAME}-${GOOS}-${GOARCH} ${PACKAGE_NAME} \
	&& cd .. \
	&& tar czf ${PACKAGE_NAME}-${VERSION}-${GOARCH}.tar.gz ${PACKAGE_NAME}-${VERSION} \
	&& tar czf ${PACKAGE_NAME}-${VERSION}-${GOARCH}-${DATE}.tar.gz ${PACKAGE_NAME}-${VERSION} \
	&& rm -rf ${PACKAGE_NAME}-${VERSION}

package-dnsctl-aliyun-amd64: PACKAGE_NAME := dnsctl-aliyun
package-dnsctl-aliyun-amd64: GOOS := linux
package-dnsctl-aliyun-amd64:
	@ cd ${BINDIR} \
	&& mkdir -p ${PACKAGE_NAME}-${ALIYUN_VERSION} \
	&& cp ${PACKAGE_NAME}-${GOOS}-${GOARCH} ${PACKAGE_NAME}-${ALIYUN_VERSION} \
	&& cd ${PACKAGE_NAME}-${ALIYUN_VERSION} \
	&& ln -s ${PACKAGE_NAME}-${GOOS}-${GOARCH} ${PACKAGE_NAME} \
	&& cd .. \
	&& tar czf ${PACKAGE_NAME}-${ALIYUN_VERSION}-${GOARCH}.tar.gz ${PACKAGE_NAME}-${ALIYUN_VERSION} \
	&& tar czf ${PACKAGE_NAME}-${ALIYUN_VERSION}-${GOARCH}-${DATE}.tar.gz ${PACKAGE_NAME}-${ALIYUN_VERSION} \
	&& rm -rf ${PACKAGE_NAME}-${ALIYUN_VERSION}

package-dnsctl-aliyun-arm64: PACKAGE_NAME := dnsctl-aliyun
package-dnsctl-aliyun-arm64: GOOS := linux
package-dnsctl-aliyun-arm64:
	@ cd ${BINDIR} \
	&& mkdir -p ${PACKAGE_NAME}-${ALIYUN_VERSION} \
	&& cp ${PACKAGE_NAME}-${GOOS}-${GOARCH} ${PACKAGE_NAME}-${ALIYUN_VERSION} \
	&& cd ${PACKAGE_NAME}-${ALIYUN_VERSION} \
	&& ln -s ${PACKAGE_NAME}-${GOOS}-${GOARCH} ${PACKAGE_NAME} \
	&& cd .. \
	&& tar czf ${PACKAGE_NAME}-${ALIYUN_VERSION}-${GOARCH}.tar.gz ${PACKAGE_NAME}-${ALIYUN_VERSION} \
	&& tar czf ${PACKAGE_NAME}-${ALIYUN_VERSION}-${GOARCH}-${DATE}.tar.gz ${PACKAGE_NAME}-${ALIYUN_VERSION} \
	&& rm -rf ${PACKAGE_NAME}-${ALIYUN_VERSION}

package-dnsctl-tencent-amd64: PACKAGE_NAME := dnsctl-tencent
package-dnsctl-tencent-amd64: GOOS := linux
package-dnsctl-tencent-amd64:
	@ cd ${BINDIR} \
	&& mkdir -p ${PACKAGE_NAME}-${TENCENT_VERSION} \
	&& cp ${PACKAGE_NAME}-${GOOS}-${GOARCH} ${PACKAGE_NAME}-${TENCENT_VERSION} \
	&& cd ${PACKAGE_NAME}-${TENCENT_VERSION} \
	&& ln -s ${PACKAGE_NAME}-${GOOS}-${GOARCH} ${PACKAGE_NAME} \
	&& cd .. \
	&& tar czf ${PACKAGE_NAME}-${TENCENT_VERSION}-${GOARCH}.tar.gz ${PACKAGE_NAME}-${TENCENT_VERSION} \
	&& tar czf ${PACKAGE_NAME}-${TENCENT_VERSION}-${GOARCH}-${DATE}.tar.gz ${PACKAGE_NAME}-${TENCENT_VERSION} \
	&& rm -rf ${PACKAGE_NAME}-${TENCENT_VERSION}

package-dnsctl-tencent-arm64: PACKAGE_NAME := dnsctl-tencent
package-dnsctl-tencent-arm64: GOOS := linux
package-dnsctl-tencent-arm64:
	@ cd ${BINDIR} \
	&& mkdir -p ${PACKAGE_NAME}-${TENCENT_VERSION} \
	&& cp ${PACKAGE_NAME}-${GOOS}-${GOARCH} ${PACKAGE_NAME}-${TENCENT_VERSION} \
	&& cd ${PACKAGE_NAME}-${TENCENT_VERSION} \
	&& ln -s ${PACKAGE_NAME}-${GOOS}-${GOARCH} ${PACKAGE_NAME} \
	&& cd .. \
	&& tar czf ${PACKAGE_NAME}-${TENCENT_VERSION}-${GOARCH}.tar.gz ${PACKAGE_NAME}-${TENCENT_VERSION} \
	&& tar czf ${PACKAGE_NAME}-${TENCENT_VERSION}-${GOARCH}-${DATE}.tar.gz ${PACKAGE_NAME}-${TENCENT_VERSION} \
	&& rm -rf ${PACKAGE_NAME}-${TENCENT_VERSION}

.PHONY: clean
clean:
	@cd ${BINDIR} \
	&& rm -f *.tar.gz
