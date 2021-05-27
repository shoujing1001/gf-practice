## 请求参数映射

GF框架将输入和输出定义为struct结构体，便于结构体参数维护。其支持非常便捷的参数->结构体对象转换。
参考文档：
[请求输入-对象处理](https://goframe.org/pages/viewpage.action?pageId=1114185)

[类型转换-Struct转换](https://goframe.org/pages/viewpage.action?pageId=1114345)

### 默认映射规则

1. struct中需要匹配的属性必须为公开属性(首字母大写)。
2. 参数名称会自动按照 `不区分大小写` 且 `忽略-/_/空格符号` 的形式与struct属性进行匹配。
3. 如果匹配成功，那么将键值赋值给属性，如果无法匹配，`那么忽略该键值`。

匹配示例：
```
map键名    struct属性     是否匹配
name       Name           match
Email      Email          match
nickname   NickName       match
NICKNAME   NickName       match
Nick-Name  NickName       match
nick_name  NickName       match
nick name  NickName       match
NickName   Nick_Name      match
Nick-name  Nick_Name      match
nick_name  Nick_Name      match
nick name  Nick_Name      match
```

### 自定义映射规则
自定义参数映射通过struct属性绑定tag实现，tag名称可为`p/param/params`

如：
```golang
type User struct{
  Id    int
  Name  string
  Pass1 string `p:"password1"`
  Pass2 string `p:"password2"`
}
```
该结构体利用tag的`p`标签指定了该属性的password1映射为Pass1。

## 请求参数校验

请求参数校验通过结构体属性绑定`v`标签实现。
参考文档: [请求输入-请求校验](https://goframe.org/pages/viewpage.action?pageId=1114244)

[数据校验-结构体校验](https://goframe.org/pages/viewpage.action?pageId=1114179)

如：
```golang
// 注册请求数据结构
type RegisterReq struct {
	Name  string `p:"username"  v:"required|length:6,30#请输入账号|账号长度为:min到:max位"`
	Pass  string `p:"password1" v:"required|length:6,30#请输入密码|密码长度不够"`
	Pass2 string `p:"password2" v:"required|length:6,30|same:password1#请确认密码|密码长度不够|两次密码不一致"`
}
```

参数校验规则参考文档：[数据校验-校验规则](https://goframe.org/pages/viewpage.action?pageId=1114367)

### 校验错误处理

上述例子中，所有不满足的条件都会返回，如
```shell
$ curl "http://127.0.0.1:8199/register?name=john&password1=123456&password2=12345"
{"code":1,"error":"密码长度不够; 两次密码不一致","data":null}
```
这样对用户体验不太友好，可以将校验错误转换为`*gvalid.Error`对象，随后可以通过灵活的方法控制错误的返回。

如：
```golang
var req *RegisterReq
if err := r.Parse(&req); err != nil {
  // Validation error.
  // 通过err.(gvalid.Error)断言的方式判断错误是否为校验错误
  // 如果是的话则返回第一条校验错误，而不是所有都返回
  if v, ok := err.(gvalid.Error); ok {
    r.Response.WriteJsonExit(RegisterRes{
      Code:  1,
      Error: v.FirstString(),
    })
  }
  // Other error.
  r.Response.WriteJsonExit(RegisterRes{
    Code:  1,
    Error: err.Error(),
  })
}
```
这样设置后就只返回第一条校验错误信息：
```shell
$ curl "http://127.0.0.1:8199/register"
{"code":1,"error":"请输入账号","data":null}
```

更详细的错误控制方法，参考文档 [数据校验-校验结果](https://goframe.org/pages/viewpage.action?pageId=1114251) 。

## 请求参数默认值

从v1.15版本开始，Request请求对象支持通过struct tag的方式为输入对象的属性绑定默认值。默认值的struct tag名称为d(也可以使用default)。

例： 
```golang
type ContentServiceGetListReq struct {
	Type       string                                    // 内容模型
	CategoryId uint   `p:"cate"`                         // 栏目ID
	Page       int    `d:"1"  v:"min:0#分页号码错误"`      // 分页号码
	Size       int    `d:"10" v:"max:50#分页数量最大50条"` // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}
```