ARG BUILD_FROM=alpine:3.14.0

FROM ${BUILD_FROM}

LABEL maintainer="Docker Maintainers <chinatree2012@gmail.com>"

ARG DOWNLOAD_URL=https://dockerfile.yuntree.com
ARG DOWNLOAD_URI=
ARG PACKAGE_NAME=dnsctl
ARG PACKAGE_VERSION=v0.0.1
ARG RUN_USER=yuntree
ARG RUN_GROUP=devops
ARG WORKDIR=/opt

RUN set -eux; \
    apkArch="$(apk --print-arch)"; \
    case "${apkArch}" in \
        armhf) arch='armv6' ;; \
        aarch64) arch='arm64' ;; \
        x86_64) arch='amd64' ;; \
        *) echo >&2 "error: unsupported architecture: ${apkArch}"; exit 1 ;; \
    esac; \
    # apk repositorie replace it with aliyun
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories; \
    # setting Shanghai timezone
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime; \
    \
    apk add --no-cache \
        'su-exec>=0.2' \
        tzdata; \
    addgroup -g 1000 ${RUN_GROUP}; \
    adduser -D -u 1000 -G ${RUN_GROUP} ${RUN_USER}; \
    cd /opt; \
    file=${PACKAGE_NAME}-${PACKAGE_VERSION}-${arch}.tar.gz; \
    prefix=${PACKAGE_NAME}-${PACKAGE_VERSION}; \
    wget -O ${file} ${DOWNLOAD_URL}${DOWNLOAD_URI}/${file}; \
    tar xf ${file}; \
    rm -f ${file}; \
    ln -s ${prefix} ${PACKAGE_NAME}; \
    chown -Rf ${RUN_USER}:${RUN_GROUP} ${prefix}; \
    ln -s ${WORKDIR}/${prefix}/${PACKAGE_NAME} /usr/local/bin/;

WORKDIR ${WORKDIR}

USER ${RUN_USER}

ENV ALIYUN_REGION_ID="default" \
    ALIYUN_SECRET_ID="" \
    ALIYUN_SECRET_KEY="" \
    TENCENT_REGION="" \
    TENCENT_SECRET_ID="" \
    TENCENT_SECRET_KEY=""

ENTRYPOINT ["dnsctl"]
