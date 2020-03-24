# go-leaf
借鉴美团分布式id生成框架leaf思想，基于golang开发分布式id生成工具，以下是具体的步骤


## 依赖包获取

1、http web框架gin

go get -u github.com/gin-gonic/gin

2、数据库orm框架 gorm

go get -u github.com/jinzhu/gorm

数据库驱动

go get -u github.com/go-sql-driver/mysql

## 代码框架
--base      //通用包

  --db      //数据库连接初始化
  
  --error   //异常包
  
  --utils   //工具包
  
--cmd       //启动包

  --main.go
  
--conf      //配置文件包

--ids       //id生成模块

  --controller  //handler
  
  --dao         //数据库操作
  
  --entity      //映射对象
  
  --serilaizer  //序列化包
  
  --service     //业务代码包
  
  --router.go