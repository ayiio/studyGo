### go.mod
记录当前项目依赖的第三方包信息和版本信息，第三方依赖包保存在`GOPATH/pkg/mod`目录下

### go.sum
详细的包名和版本信息

### 常用命令
```bash
go mod init [包名]  // 初始化项目
go mod tidy        // 检查代码里的依赖，更新go.mod文件中的依赖
go get             // 下载依赖
go mod download    // 下载依赖
```
