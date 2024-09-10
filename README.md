# WHOIM-user-service

kitex -module github.com/lvkeliang/WHOIM-user-service -service UserService thrift/user.thrift

kitex -module github.com/your_project_name/RPC -service UserService your_thrift_file.thrift

go get github.com/kitex-contrib/registry-etcd