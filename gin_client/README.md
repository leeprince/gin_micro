# gin 项目基础架构

---


# gin 框架
```
go get -u github.com/gin-gonic/gin
```

# gorm
 ```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

# 启动
1. go mod tidy
2. go run main.go

# 测试
## 健康检查
```
curl http://localhost:8080/ping
```

## 根据ID获取指定用户信息
```
curl http://localhost:8080/v1/getUser\?userId\=1
```