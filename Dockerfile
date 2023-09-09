#Dockerfile
FROM  golang:latest

WORKDIR /app

COPY .  .

EXPOSE 8081

# 构建Go项目
RUN go build -o main .

ENTRYPOINT [ "./main" ]