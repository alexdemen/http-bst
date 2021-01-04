FROM golang:1.15.6-alpine
WORKDIR /app
COPY . .
RUN cd ./cmd/http-bst && ls /app &&CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/bst-app

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /app/bst-app .
COPY --from=0 /app/env/init.json .
ENV TREE_INIT_FILE=/root/init.json
CMD ["./bst-app"]