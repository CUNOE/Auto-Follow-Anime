# Auto Follow Anime

一个自动追番的工具

## 原理

通过Kisssub网站的Rss链接来获取更新的番剧数据，将读取到的磁力链接通过JSONRPC发送到aria2下载，再通过cqhttp发送消息到指定的人或者群

## 配置

在`config.json`文件中你需要填写相关信息，具体请查看`config_example.json`文件

## 使用

通过docker构建并且部署即可
