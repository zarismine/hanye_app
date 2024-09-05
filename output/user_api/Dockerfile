FROM alpine:latest
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata
COPY output/user_api/user.yaml user.yaml
COPY output/user_api/user_api user_api

ENTRYPOINT ["/user_api", "-f", "user.yaml"]