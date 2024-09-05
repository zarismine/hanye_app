FROM alpine:latest
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata
COPY output/order_api/order.yaml order.yaml
COPY output/order_api/order_api order_api

ENTRYPOINT ["/order_api", "-f", "order.yaml"]