## 常用命令
```
gf gen dao  -path  ./app/api/ -t prac_user,prac_todo -r prac_
```

## 一些理解

### 项目结构思考

1. 项目结构
https://goframe.org/pages/viewpage.action?pageId=3672600
2. 一个目录下的文件中包名与目录名一致，可通过不同文件下开放的遵守一定命名规则的结构体来进行区分。如`app/api/controller/user.go`中，包名为controller，其中暴露的`userApi`结构体即可用于注册路由。那service、dao、model目录同样适用这样的规则。
3. 目录下还可嵌套新的目录以放置独立的包，如`app/api/controller/hello/`目录和dao、model目录下的internal目录都是如此。


## 今日学习参考

https://github.com/gogf/gf-demos/blob/master/app/api/user.go