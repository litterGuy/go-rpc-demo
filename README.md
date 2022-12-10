A branch for protobuf & gRPC demo.

### 安装protoc和protoc-gen-go

#### protoc

Github的Release下载：

[https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases)

下载对应平台的压缩包(win/Linux);

解压之后将bin目录下的protoc复制到系统环境变量的路径下:

-   win: `C:\Windows\System32`
-   Linux: `/bin`

等等;

然后测试:

```bash
protoc --version
libprotoc 3.19.4
```

****

#### protoc-gen-go

Github的Release下载:

[https://github.com/golang/protobuf/releases](https://github.com/golang/protobuf/releases)

下载压缩包并解压;

解压后在`protoc-gen-go`目录下编译:

```bash
go build
```

然后将生成的二进制文件`protoc-gen-go`复制到系统环境变量的路径下

全部安装完成!

<br/>

### 目录结构

```bash
tree /F                                       
.                                             
│  go.mod                                       
│  go.sum                                       
│                                               
├─hello_client                                  
│      main.go                                  
│                                               
├─hello_server                                  
│      main.go                                  
│                                               
└─proto                                         
        hello.pb.go                             
        hello.proto                             
```

<br/>

### 编写proto

本例子使用proto3语法;

hello.proto

```protobuf
// 指定proto版本
syntax = "proto3";

package proto;
option go_package="./../proto";

// 定义请求结构
message HelloRequest{
    string name = 1;
}

// 定义响应结构
message HelloResponse{
    string message = 1;
}

// 定义Hello服务
service Hello{
    // 定义服务中的方法
    rpc SayHello (HelloRequest) returns (HelloResponse);
}
```

编写完成之后使用protoc生成`.pb.go`文件:

```bash
protoc -I . --go_out=. --go-grpc_out=. .\hello.proto
```

即可生成hello.pb.go文件

<br/>

### 初始化gomod项目

使用`go mod init protobuf_grpc_demo`初始化项目;

生成go.mod文件;

然后使用`go mod tidy`初始化;


