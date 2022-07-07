# whoami
打印操作系统信息和 HTTP 请求等信息的服务
# docker
```
docker run -d -p 8080:8080 --name foo freemesh/whoami
```

# api
```
# 打印信息
curl http://localhost:8080
# 健康检查
curl -v http://localhost:8080/healthz
```

# docker
