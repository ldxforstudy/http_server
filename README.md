# 1.目的
用于测试Gateway项目而模拟的后端HTTP服务.

# 2.运行
```
cd $YOUR_GO_WORKSPACE/src/
git clone THIS_REPOSITORY

go run main.go
```

服务访问地址为 `http://localhost:8080`

# 3.API响应描述
成功:
```
{
    "retcode": 200,
    "res": {
        "id":"10086",
        "name":"user1"
    }
}
```
失败:
```
{
    "retcode": NOT_200_CODE,
    "msg": "ERR-MSG"
}
```

# 4.API列表
| URL        | Method           | Desc  |
| ------------- |:-------------:| -----:|
| /user/     |GET  |  用户信息|
| /order/      | GET      |   订单信息 |  