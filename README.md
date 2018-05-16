# loach golang 实现fabric-ca sdk

### 环境

```
针对版本：fabric-ca 1.1.0
```

### 项目依赖


```
go get -u golang.org/x/crypto/sha3
go get -u gopkg.in/yaml.v2
```

### 安装


```
go get -u learnergo/loach
```

### 测试


```
static\file.yaml 主要配置ca地址、名称
static下admin.key和admin.crt为管理员私钥和证书（根据自己替换）

主要用法参考 client_test.go
```

### TODO

```
目前只实现了Register和Enroll，后续会增加其他方法
```

### 项目应用

```
github.com/learnergo/cuttle 项目应用该项目
```
