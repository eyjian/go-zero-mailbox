# mooon-mailbox

一款简单的基于 go-zero 实现的信箱，支持：投递信件、批量查看信件、批量标记为已读和删除信件。

# 如何使用？

第一次编译前，请先执行一次"make tidy"，在运行之前依据 mooon_mailbox.sql 创建好 MySQL 表。下为一个启动示例：

```shell
./mooon_mailbox_service --dsn='dbuser:dbpassword@tcp(dbhost:dbport)/dbname?charset=utf8mb3&parseTime=True&loc=Local'
```

配置文件参考 etc 目录下的 mooon_mailbox.yaml 文件。其中 DB 数据源即可参数指定，也可配置文件指定，但参数优先。

# 如何测试？

测试方法，请参考 TEST.md 文件。
