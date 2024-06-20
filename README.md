# restful-api

#docker 运行 mysql
```
docker run --name host-mysql -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=host -e MYSQL_USER=hostroot -e MYSQL_PASSWORD=123456 -p 3306:3306 -d mysql:latest
```

# Go 语言中的类型断言
1. 在 Go 中，断言（Type Assertion）用于从接口类型断言为具体类型。它允许你检查一个接口变量是否保存了特定的具体类型值，并从接口中提取该具体类型值。

2. 基本语法是：x.(T)，其中 x 是接口类型的变量，T 是你要断言的具体类型。如果断言成功，结果是具体类型的值；如果断言失败，会触发运行时错误（panic）。
