FROM golang:1.14.4-alpine3.12 AS build
ENV CGO_ENABLED=0 \
    GOOS=linux
ADD . /go/src
WORKDIR /go/src/github.com/kaissi/fct7-devops-kubernetes-desafio
RUN go build -ldflags '-w -s' -a -installsuffix cgo -o /go/bin/desafio-k8s \
    && apk add --no-cache upx \
    && upx --brute /go/bin/desafio-k8s

FROM scratch
COPY --from=build /go/bin/desafio-k8s /usr/local/bin/
EXPOSE 8000
ENTRYPOINT ["desafio-k8s"]

