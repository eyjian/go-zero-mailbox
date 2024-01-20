**构建信箱服务步骤：**

* 执行"git clone"取下代码，或者 wget 方式下载源代码包
* 执行"make tidy"
* 执行 make 编译，成功后生成可执行的信箱服务文件 mooon_mailbox_service

**信箱服务启动示例：**

```shell
./mooon_mailbox_service --dsn='dbuser:dbpassword@tcp(dbhost:dbport)/dbname?charset=utf8mb3&parseTime=True&loc=Local'
```

注意，需修改参数"--dsn"的值，这启动之前还需要修改好配置文件 mooon_mailbox.yaml 。
