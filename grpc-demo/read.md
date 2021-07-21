
https://golang2.eddycjy.com/posts/ch3/02-simple-protobuf/

``` 
protoc --go_out=plugins=grpc:./proto ./proto/*.proto

–go_out：设置所生成 Go 代码输出的目录，该指令会加载 protoc-gen-go 插件达到生成 Go 代码的目的，
生成的文件以 .pb.go 为文件后缀，在这里 “:”（冒号）号充当分隔符的作用，后跟命令所需要的参数集，
在这里代表着要将所生成的 Go 代码输出到所指向 protoc 编译的当前目录。

plugins=plugin1+plugin2：指定要加载的子插件列表，我们定义的 proto 文件是涉及了 RPC 服务的，
而默认是不会生成 RPC 代码的，因此需要在 go_out 中给出 plugins 参数传递给 protoc-gen-go，告诉编译器，
请支持 RPC（这里指定了内置的 grpc 插件）。
```


file_proto_helloworld_proto_rawDesc
表示的是一个经过编译后的 proto 文 件，是对 proto 文件的整体描述，其包含了 proto 文件名、引用（import）内容、
包 （package）名、选项设置、所有定义的消息体（message）、所有定义的枚举（enum）、所有定义 的服务（ service）、
所有定义的方法（rpc method）等等内容，可以认为就是整个 proto 文件 的信息你都能够取到。


### hers 

go get -a github.com/golang/protobuf/protoc-gen-go

生成golang的服务代码
```
protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
```
这个指令支持*.proto模糊匹配。如果有许多文件可以使用helloworld/*.proto 来作为PROTO_FILES

```
protoc --go_out=plugins=grpc:. *.proto
```
