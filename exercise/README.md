## 项目运行

1. 执行 git clone 命令将仓库代码 copy 到本地；
2. 指定 sqlite-serve.bat 脚本，启动本地 Windows 下的 sqlite 环境；
    1. 如果安装路径不一致，修改 bat 批处理命令后再尝试或弃用该脚本；
3. 运行命令：

```shell
cd src/main
# 执行目标函数，23 指的是随机确认的 userId 值
go run sqlite.go sqlite_arg.go sqlite_cal.go 23
```

4. 如果运行时间比较长，可以自行修改其中的参数，减少生成的用户或者订单数量。修改处为：`src/main/sqlite.go:16`，分别表示生成的用户以及订单的数量。

```go
const userCount = 10000
const orderCount = 100000
```

5. 运行之后可以看到打印输出的结果：

![image.png](https://img-blog.csdnimg.cn/img_convert/d8be488c645dc7322d4947dba6b9f193.png)
