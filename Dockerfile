FROM golang:1.14 as builder

WORKDIR /root
COPY ./server  ./
RUN export GO111MODULE=on \
 && go env -w GOPROXY=https://goproxy.cn,direct \
 && go get -d -v ./... \
 && CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM node:12-alpine as node-builder
WORKDIR /root
COPY ./web ./
RUN npm run build-prod

# FROM alpine:latest
# WORKDIR /root
# COPY --from=builder /root/app ./
## COPY --from=node-builder /root/dist dist
# COPY ./start.sh ./
FROM nginx
WORKDIR /root
COPY --from=builder /root/app ./
COPY --from=node-builder /root/dist /usr/share/nginx/html

COPY ./start.sh ./
EXPOSE 8083/tcp

CMD ["/root/start.sh"]
