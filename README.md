# Auto Follow Anime

一个自动追番的工具

## 原理

通过Kisssub网站的Rss链接来获取更新的番剧数据，将读取到的磁力链接通过JSONRPC发送到aria2下载，再通过cqhttp发送消息到指定的人或者群

## 配置
请自行建立一个afa文件夹并且放在同一个目录或者通过volumes挂载到docker里面 可以参考下面的docker compose文件
目录如下
```
- path/to/dir
	- afa
		- config.json
		- sqlite.db
	- app
```


在`config.json`文件中你需要填写相关信息，具体请查看`config_example.json`文件

## 使用

通过docker构建并且部署即可

参考的docker compose

```yaml
  auto-follow-anime:
    image: cunoe/auto-follow-anime:latest
    container_name: cunoe-auto-follow-anime
    restart: always
    environment:
      - TZ=Asia/Shanghai
    volumes:
      - /home/cunoe/afa/:/root/afa
    networks:
      - cunoenet
  cqhttp:
    image: silicer/go-cqhttp:latest
    container_name: cunoe-qqbot
    restart: always
    environment:
      - TZ=Asia/Shanghai
    volumes:
      - /home/cunoe/cqhttp/:/data
    networks:
      - cunoenet
      
```

参考的cqhttp ws设置

```
servers:
  - ws-reverse:
      universal: ws://auto-follow-anime:9000/ws
      reconnect-interval: 3000
      middlewares:
        <<: *default
```

