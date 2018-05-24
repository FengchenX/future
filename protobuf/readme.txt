下载protobuf.exe编译版本
放到go/bin目录下
将 C:\Users\fengc\go\bin\protoc-3.5.1-win32\bin 这个路径添加到环境变量中
go get -u github.com/golang/protobuf/protoc-gen-go
将 protoc-gen-go.exe添加到环境变量中

运行命令 protoc --go_out=. test.proto 生成文件
