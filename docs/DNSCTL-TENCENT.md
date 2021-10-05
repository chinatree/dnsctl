# 腾讯云 DNS 客户端工具使用手册

[TOC]

## 帮助

```bash
$ dnsctl-tencent -h
dnsctl-tencent is a simple command line client for tencent dns.

Usage:
  dnsctl-tencent [command]

Available Commands:
  domain      域名管理
  group       域名分组管理
  help        Help about any command
  record      域名解析管理
  version     版本号

Flags:
  -h, --help                help for dnsctl-tencent
      --region string       区域名称
      --secret-id string    密钥ID
      --secret-key string   加密密钥

Use "dnsctl-tencent [command] --help" for more information about a command.
```

## 鉴权

该工具的鉴权需要依赖于 Region、 SecretID 和 SecretKey，支持两种方式进行配置:

- 参数方式
  -  --region
  - --secret-id
  - --secret-key
- 环境变量
  - TENCENT_REGION
  - TENCENT_SECRET_ID
  - TENCENT_SECRET_KEY

## 域名管理

### 帮助

```bash
$ dnsctl-tencent domain
域名管理

Usage:
  dnsctl-tencent domain [command]

Available Commands:
  add         添加域名
  delete      删除域名
  get         获取域名信息
  list        获取域名列表
  lock        锁定域名
  logs        获取域名的操作日志
  remark      修改域名的备注
  status      设置域名的状态
  unlock      域名锁定解锁

Flags:
  -h, --help   help for domain

Global Flags:
      --region string       区域名称
      --secret-id string    密钥ID
      --secret-key string   加密密钥

Use "dnsctl-tencent domain [command] --help" for more information about a command.
```

### 获取域名列表

```bash
$ dnsctl-tencent domain list
```

### 获取域名信息

```bash
$ dnsctl-tencent domain get --name yuntree.com
```

### 添加域名

```bash
$ dnsctl-tencent domain add --name yuntree.com
```

### 修改域名的备注

```bash
$ dnsctl-tencent domain remark --name yuntree.com --remark 国际域名
```

### 删除域名

```bash
$ dnsctl-tencent domain delete --name yuntree.com
```

### 锁定域名

```bash
$ dnsctl-tencent domain lock --name yuntree.com --day 7
```

### 域名锁定解锁

```bash
$ dnsctl-tencent domain unlock --name yuntree.com --code <code>
```

### 获取域名操作日志

```bash
$ dnsctl-tencent domain logs --name yuntree.com
```

## 域名分组管理

### 帮助

```bash
$ dnsctl-tencent group
域名分组管理

Usage:
  dnsctl-tencent group [command]

Available Commands:
  add         添加域名分组

Flags:
  -h, --help   help for group

Global Flags:
      --region string       区域名称
      --secret-id string    密钥ID
      --secret-key string   加密密钥

Use "dnsctl-tencent group [command] --help" for more information about a command.
```

### 添加域名分组

```bash
$ dnsctl-tencent group add --name 测试分组
```

## 域名解析管理

### 帮助

```bash
$ dnsctl-tencent record -h
域名解析管理

Usage:
  dnsctl-tencent record [command]

Available Commands:
  add         添加解析记录
  delete      删除解析记录
  get         获取解析记录信息
  list        获取域名解析列表
  remark      修改解析记录的备注
  status      设置解析记录的状态
  update      修改解析记录

Flags:
  -d, --domain string    域名名称
      --domain-id uint   域名ID。如果指定该参数，将忽略 --name 参数
  -h, --help             help for record

Global Flags:
      --region string       区域名称
      --secret-id string    密钥ID
      --secret-key string   加密密钥

Use "dnsctl-tencent record [command] --help" for more information about a command.
```

### 获取域名解析列表

```bash
$ dnsctl-tencent record list --domain yuntree.com
```

### 获取解析记录信息

```bash
$ dnsctl-tencent record get --domain yuntree.com --id 927117504
```

### 添加解析记录

```bash
$ dnsctl-tencent record add --domain yuntree.com --rr www --type A --value 1.1.1.1
```

### 修改解析记录

```bash
$ dnsctl-tencent record update --domain yuntree.com --id 927117504 --rr www --type A --value 1.1.1.1
```

### 修改解析记录的备注

```bash
$ dnsctl-tencent record remark --domain yuntree.com --id 927117504 --remark 官网
```

### 设置解析记录状态

```bash
$ dnsctl-tencent record status --domain yuntree.com --id 927117504 --status ENABLE
```

### 删除解析记录

```bash
$ dnsctl-tencent record delete --domain yuntree.com --id 927117504
```

## 参考

1. [云解析 DNS API 文档](https://cloud.tencent.com/document/product/1427/56193)
2. [API 概览](https://cloud.tencent.com/document/product/1427/56194)
