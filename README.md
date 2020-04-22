##本机配置

### 运行服务程序

```
go run server/main.go
```

### 运行客户端程序

```
vue ui
```

http://localhost:8083/

## 服务器配置

### 配置nginx

```
server {
    listen       80;
    server_name  vue.golango.cn;
    location / {
        root   /usr/local/src/vue;
        index  index.html index.htm;
    }

    location /api/ {
        proxy_pass http://localhost:8092/;
	    proxy_set_header HOST $host;
	    proxy_set_header X-Real-IP $remote_addr;
	    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }   
}

```