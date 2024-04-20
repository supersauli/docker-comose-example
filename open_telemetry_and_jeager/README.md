## jaeger
### 说明
1. 网页地址 http://127.0.0.1::16686/
2. docker compose 说明

```
version: '3'
services:
  jaeger:
    image: jaegertracing/all-in-one:latest
    container_name: jaeger
    environment:
      - COLLECTOR_OTLP_ENABLED=true
    ports:
      # 网页访问地址
      - "16686:16686"
      # tcp 连接地址
      - "4317:4317"
      # http连接地址
      - "4318:4318"
    restart: unless-stopped
```


