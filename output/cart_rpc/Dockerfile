FROM alpine:latest
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata
COPY output/cart_rpc/cart.yaml cart.yaml
COPY output/cart_rpc/cart_rpc cart_rpc

ENTRYPOINT ["/cart_rpc", "-f", "cart.yaml"]