## work 3 - ToDoList

### 项目使用框架: hertz+gorm

项目treer生成: notejs-treer



Bonus1 自动生成接口文档: 

使用了swaggo



Bonus2 使用三层架构设计: 并没有完成，但可以说思路上符合(?

表示层: 对用户发来的请求进行处理(如request body、param)与最后呈递用户所需数据(如JSON)(handler)

业务逻辑层:  通过验证来确认数据(如Authorization)是否传给DAL(还是handler)

数据访问层: 使用model和DB来操作数据库进行增删改查(service)



Bonus3 数据库交互安全性: 

避免字符串拼接以防止sql注入



Bonus4 思考一个比要求中的结构更优秀的返回结构: 

我觉得我的返回结构更混沌了! 甚至自己都没规范好什么时候该用model.Response和model.ErrorResponse，

比如在DAL层出现错误会返回Response，而在BLL、UI层才会是ErrorResponse



Bonus5 对项目使用Redis: 锐意策划中