# WHOIM-user-service

kitex -module github.com/lvkeliang/WHOIM-user-service -service UserService thrift/user.thrift

kitex -module whotest -service UserService thrift/user.thrift

go get github.com/kitex-contrib/registry-etcd

------------------------

docker exec -it cassandra cqlsh

CREATE KEYSPACE whoim WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

USE whoim;

CREATE TABLE users (
id UUID PRIMARY KEY,
username TEXT,
password_hash TEXT,
email TEXT,
created_at TIMESTAMP,
updated_at TIMESTAMP
);

CREATE INDEX ON users (username);

------------------------

HSET user:user123:devices device456 server1
HSET device:device456 userID user123
HSET device:device456 serverAddress server1
