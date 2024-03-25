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

- 配置MySQL(和SQLite3二选一)
```text

```


- 配置MongoDB

```shell

```

use course2024
db.createUser({ user: "course2024", pwd: "course2024", roles: [{ role: "dbOwner", db: "course2024" }] })

- 安装protoc工具
```text

```
  

- 生成proto
```shell
python ./gen.py --proto_dir="./protobuf" --lang=go --gen_dir=./proto
python ./gen.py --proto_dir="./protobuf" --lang=csharp --gen_dir=../MonsterQuest/Assets/Assemblies/Proto
```

- Grpc测试
```text
使用BloomGrpc/代码自行测试
```

