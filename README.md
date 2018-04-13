# loach

针对版本：fabric-ca 1.1.0

static目录：
  file.yaml 主要配置ca地址
  admin.key、 admin.crt为管理员私钥和证书（根据自己替换）

主要用法参考 client_test.go

目前只实现了Register和Enroll，后续会增加其他方法
