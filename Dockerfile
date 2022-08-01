FROM golang:1.18-alpine as build
RUN go env -w CGO_ENABLED="0" GOPROXY="https://goproxy.cn,direct"
ADD . /go/src/github.com/cnt-camp/
RUN cd /go/src/github.com/cnt-camp/module2/ && go build -o /go/bin/httpserver
ENTRYPOINT ["/go/bin/httpserver"]

FROM scratch
COPY --from=build /go/bin/httpserver /bin/httpserver
ENTRYPOINT ["/bin/httpserver"]