FROM alpine:latest
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata
COPY output/order_rpc/order.yaml order.yaml
COPY output/order_rpc/order_rpc order_rpc

ENTRYPOINT ["/order_rpc", "-f", "order.yaml"]