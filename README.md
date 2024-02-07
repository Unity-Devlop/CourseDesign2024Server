# Server

- 初始化项目
```shell
go mod init
go mod tidy
```

- 安装SQLite3
```text
链接:https://www.sqlite.org/download.html
配置SQLite3到环境变量:PATH
```
  

- 生成proto
```shell
python ./gen.py --proto_dir="./protobuf" --lang=go --gen_dir=./proto
python ./gen.py --proto_dir="./protobuf" --lang=csharp --gen_dir=../CourseDesign2024/Assets/Assemblies/Proto
```

- Grpc服务测试

使用bloomgrpc测试grpc服务
