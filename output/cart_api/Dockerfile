FROM alpine:latest
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata
COPY output/cart_api/cart.yaml cart.yaml
COPY output/cart_api/cart_api cart_api

ENTRYPOINT ["/cart_api", "-f", "cart.yaml"]