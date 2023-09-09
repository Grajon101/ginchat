项目使用：
1.Gorm  "gorm.io/gorm" 
Gorm一个强大的Go编程语言中的ORM（对象关系映射）库。ORM是一种技术，它将数据库表中的数据映射到面向对象的模型中，从而简化了数据库操作
文档：https://gorm.io/zh_CN/docs/

2.mysql "gorm.io/driver/mysql" 
开源关系型数据库驱动

3.viper "github.com/spf13/viper" 
主要是用于处理各种格式的配置文件，简化程序配置的读取问题，1.支持Yaml、Json、 TOML、HCL 等格式的配置 2.可以从文件、io、环境变量、command line中提取配置 3.支持自动转换的类型解析  4.可以远程从etcd中读取配置
文档：https://www.cnblogs.com/QiaoPengjun/p/17489207.html

4.gin "github.com/gin-gonic/gin"
Gin 是一个用 Go语言 (Golang) 编写的 Web 框架
文档：http://www.topgoer.com/gin%E6%A1%86%E6%9E%B6/%E7%AE%80%E4%BB%8B.html

5.ginSwagger "github.com/swaggo/gin-swagger" "github.com/swaggo/files"
swagger是一套基于OpenAPI规范构建的开源工具，使用RestApi。swagger-ui 呈现出来的是一份可交互式的API文档，可以直接在文档页面尝试API的调用 
更多用法：https://blog.csdn.net/qq_41630102/article/details/128411210

6.govalidator "github.com/asaskevich/govalidator"
数据验证
翻译中文：https://blog.csdn.net/qq_42887507/article/details/120934568



一、go mod 
go modules 官方定义为：
模块是相关Go包的集合。modules是源代码交换和版本控制的单元。
go命令直接支持使用modules，包括记录和解析对其他模块的依赖性。modules替换旧的基于GOPATH的方法来指定在给定构建中使用哪些源文件。

详细命令
1. init
代码：go mod init
生成 go.mod 文件，此命令会在当前目录中初始化和创建一个新的go.mod文件，手动创建go.mod文件再包含一些module声明也等同该命令，而go mod init命令便是帮我们简便操作，可以帮助我们自动创建。

2.download
代码：go mod download
下载 go.mod 文件中指明的所有依赖，使用此命令来下载指定的模块，模块的格式可以根据主模块依赖的形式或者path@version形式指定。
3.tidy
代码：go mod tidy
整理现有的依赖，使用此命令来下载指定的模块，并删除已经不用的模块
4.graph
代码：go mod graph
查看现有的依赖结构，生成项目所有依赖的报告，但可读性太差，图形化更方便。
5.edit
代码：go mod edit
编辑 go.mod 文件，之后通过 download 或 edit 进行下载
6.vendor
代码：go mod vendor
导出项目所有的依赖到vendor目录，从mod中拷贝到项目的vendor目录下，IDE可以识别这样的目录。
7.verify
代码：go mod verify
校验一个模块是否被篡改过，查询某个常见的模块出错是否已被篡改
8.why
代码：go mod why
查看为什么需要依赖某模块，查询某个不常见的模块是否是哪个模块的引用

二.git
git init
//所有文件
git add .
git commit -m "first commit"
git commit -m "更新"
git branch -M main
第一次
git remote add origin git@github.com:Grajon101/ginchat.git
git push -u origin main

三.docker

构建镜像
docker build -t ginchat .

运行容器
docker run -d -p 8080:8081 --name mychat ginchat

docker run -d --name minio \
    --publish 9000:9000 \
    --publish 9001:9001 \
    --env MINIO_ROOT_USER="root" \
    --env MINIO_ROOT_PASSWORD="19911106" \
    bitnami/minio:latest


三.错误解决
1.
gin + gin-swagger，swag init 命令执行错误 cannot find type definition
需要加参数  swag init 命令后面添加两个参数 --parseDependency --parseInternal 
如：swag init --parseDependency --parseInternal 
或者简写 swag init --pd 
2.

