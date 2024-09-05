FROM alpine:latest
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata
COPY output/user_rpc/user.yaml user.yaml
COPY output/user_rpc/user_rpc user_rpc

ENTRYPOINT ["/user_rpc", "-f", "user.yaml"]