# 阿里云 DNS 客户端工具使用手册

[TOC]

## 帮助

```bash
$ dnsctl-aliyun -h
dnsctl-aliyun is a simple command line client for aliyun dns.

Usage:
  dnsctl-aliyun [command]

Available Commands:
  domain      域名管理
  group       域名分组管理
  help        Help about any command
  record      域名解析管理
  version     版本号

Flags:
  -h, --help                help for dnsctl-aliyun
      --region-id string    区域ID (default "default")
      --secret-id string    密钥ID
      --secret-key string   加密密钥

Use "dnsctl-aliyun [command] --help" for more information about a command.
```

## 鉴权

该工具的鉴权需要依赖于 RegionID、 AccessKey 和 AccessSecret，支持两种方式进行配置:

- 参数方式
  -  --region-id
  - --secret-id
  - --secret-key
- 环境变量
  - ALIYUN_REGION_ID
  - ALIYUN_SECRET_ID
  - ALIYUN_SECRET_KEY

## 域名管理

### 帮助

```bash
$ dnsctl-aliyun domain
域名管理

Usage:
  dnsctl-aliyun domain [command]

Available Commands:
  add         添加域名
  delete      删除域名
  get         获取域名信息
  getMain     获取主域名名称
  list        获取域名列表
  logs        获取域名操作日志
  remark      修改域名的备注

Flags:
  -h, --help   help for domain

Global Flags:
      --region-id string    区域ID (default "default")
      --secret-id string    密钥ID
      --secret-key string   加密密钥

Use "dnsctl-aliyun domain [command] --help" for more information about a command.
```

### 获取域名列表

```bash
$ dnsctl-aliyun domain list
```

### 获取域名信息

```bash
$ dnsctl-aliyun domain get --name yuntree.com
```

### 添加域名

```bash
$ dnsctl-aliyun domain add --name yuntree.com
```

### 修改域名的备注

```bash
$ dnsctl-aliyun domain remark --name yuntree.com --remark 国际域名
```

- 删除域名

```bash
$ dnsctl-aliyun domain delete --name yuntree.com
```

### 获取主域名名称

```bash
$ dnsctl-aliyun domain getMain --name cdn.yuntree.com
```

### 获取域名操作日志

```bash
$ dnsctl-aliyun domain logs --lang zh
```

## 域名分组管理

### 帮助

```bash
$ dnsctl-aliyun group
域名分组管理

Usage:
  dnsctl-aliyun group [command]

Available Commands:
  add         添加域名分组
  change      更改域名分组
  delete      删除域名分组
  list        获取域名分组列表
  update      修改域名分组

Flags:
  -h, --help   help for group

Global Flags:
      --region-id string    区域ID (default "default")
      --secret-id string    密钥ID
      --secret-key string   加密密钥

Use "dnsctl-aliyun group [command] --help" for more information about a command.
```

### 获取域名分组列表

```bash
$ dnsctl-aliyun group list
```

### 添加域名分组

```bash
$ dnsctl-aliyun group add --name 测试分组
```

### 修改域名分组

```bash
$ dnsctl-aliyun group update --id fd31288b67e94dbdbab444e70937fb19 --name 测试分组
```

### 删除域名分组

```bash
$ dnsctl-aliyun group delete --id fd31288b67e94dbdbab444e70937fb19
```

### 更改域名分组

```bash
$ dnsctl-aliyun group change --domain yuntree.com --id fd31288b67e94dbdbab444e70937fb19
```

## 域名解析管理

### 帮助

```bash
$ dnsctl-aliyun record -h
域名解析管理

Usage:
  dnsctl-aliyun record [command]

Available Commands:
  add         添加解析记录
  delete      删除解析记录
  get         获取解析记录信息
  list        获取域名解析列表
  listSub     获取子域名解析记录列表
  logs        获取解析记录操作日志
  remark      修改解析记录的备注
  status      设置解析记录状态
  update      修改解析记录

Flags:
  -d, --domain string   域名名称
  -h, --help            help for record

Global Flags:
      --region-id string    区域ID (default "default")
      --secret-id string    密钥ID
      --secret-key string   加密密钥

Use "dnsctl-aliyun record [command] --help" for more information about a command.
```

### 获取域名解析列表

```bash
$ dnsctl-aliyun record list --domain yuntree.com
```

### 获取子域名解析记录列表

```bash
$ dnsctl-aliyun record listSub --domain yuntree.com --sub-domain cdn.yuntree.com
```

### 添加解析记录

```bash
$ dnsctl-aliyun record add --domain yuntree.com --rr www --type A --value 1.1.1.1
```

### 修改解析记录

```bash
$ dnsctl-aliyun record update --id 721046463884277760 --rr www --type A --value 1.1.1.1
```

### 修改解析记录的备注

```bash
$ dnsctl-aliyun record remark --id 721046463884277760 --remark 官网
```

### 设置解析记录状态

```bash
$ dnsctl-aliyun record status --id 721046463884277760 --status ENABLE
```

### 删除解析记录

```bash
$ dnsctl-aliyun record delete --id 721046463884277760
```

### 获取解析记录操作日志

```bash
$ dnsctl-aliyun record logs --domain yuntree.com --lang zh
```

## 参考

1. [云解析 DNS API 文档](https://help.aliyun.com/document_detail/29738.html)
2. [域名管理接口](https://help.aliyun.com/document_detail/29748.html)
3. [域名分组接口](https://help.aliyun.com/document_detail/29761.html)
4. [解析管理接口](https://help.aliyun.com/document_detail/29771.html)
