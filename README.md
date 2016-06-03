domeos/sender
=============

## Notice

domeos/sender模块是以open-falcon原生sender模块为基础，为适应DomeOS监控报警需求而设计修改的，包名已修改为`github.com/domeos/sender`

原生open-falcon系统中，alarm处理报警event可能会产生报警短信或者报警邮件，alarm不负责发送，只是把报警邮件、短信写入redis队列，sender负责读取并发
送。sender的配置文件cfg.json中配置了api:sms和api:mail，即两个http接口，以适应不同公司发送短信和邮件的需求。

当要发送短信的时候，sender就会调用api:sms中配置的http接口，post方式，参数是：

- tos：用逗号分隔的多个手机号
- content：短信内容

当要发送邮件的时候，sender就会调用api:mail中配置的http接口，post方式，参数是：

- tos：用逗号分隔的多个邮箱地址
- content：邮件正文
- subject：邮件标题

在DomeOS中，api:sms和api:mail改为由DomeOS的数据库全局配置中读取得到。当DomeOS全局配置改变时DomeOS服务器将主动调用sender接口/config/api/reload
从DomeOS数据库中更新新的短信和邮件API接口；同时sender也会每隔60s自动向DomeOS数据库拉取短信和邮件API接口。

## Installation

```bash
# set $GOPATH and $GOROOT
mkdir -p $GOPATH/src/github.com/domeos
cd $GOPATH/src/github.com/domeos
git clone https://github.com/domeos/sender.git
cd sender
go get ./...
./control build
# vi cfg.json modify configuration
./control start
```

## Configuration

- database: DomeOS数据库地址，需提供用户名、密码、地址与对应端口
- maxIdle: MySQL连接池最大空闲连接数
- http: 监听的http端口
- redis: redis地址需要和alarm、judge使用同一个
- queue: 维持默认即可，需要和alarm的配置一致
- worker: 最多同时有多少个线程调用短信、邮件发送接口
- api: 发送短信和邮件的接口

## Run In Docker Container

首先构建domeos/sender镜像：

```bash
sudo docker build -t="domeos/sender:latest" ./docker/
```

启动docker容器：
```bash
sudo docker run -d --restart=always \
    -p <_sender_http_port>:6066 \
    -e DATABASE="\"<_domeos_db_user>:<_domeos_db_passwd>@tcp(<_domeos_db_addr>)/domeos?loc=Local&parseTime=true\"" \
    -e REDIS_ADDR="\"<_redis>\"" \
    --name sender \
    pub.domeos.org/domeos/sender:1.0
```

参数说明：

- _sender_http_port: sender服务http端口，主要用于状态检测、调试等。
- _domeos_db_user: DomeOS中MySQL数据库的用户名。
- _domeos_db_passwd: DomeOS中MySQL数据库的密码。
- _domeos_db_addr: DomeOS中MySQL数据库的地址，格式为IP:Port。
- _redis: 用于报警的redis服务地址，格式为IP:Port。

样例：

```bash
sudo docker run -d --restart=always \
    -p 6066:6066 \
    -e DATABASE="\"root:root@tcp(10.16.42.199:3306)/domeos?loc=Local&parseTime=true\"" \
    -e REDIS_ADDR="\"10.16.42.199:6379\"" \
    --name sender \
    pub.domeos.org/domeos/sender:1.0
```

验证：

通过curl -s localhost:<_sender_http_port>/health命令查看运行状态，若运行正常将返回ok。

DomeOS仓库中domeos/sender对应版本：pub.domeos.org/domeos/sender:1.0
