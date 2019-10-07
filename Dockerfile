FROM golang:1.13.1-alpine AS build

WORKDIR /go/src/sources.dev.pztrn.name/pztrn/giredore
COPY . .

WORKDIR /go/src/gitlab.com/pztrn/fastpastebin/cmd/fastpastebin

RUN cd /go/src/sources.dev.pztrn.name/pztrn/giredore/cmd/giredored && go build && cd ../giredorectl && go build

FROM alpine:latest
LABEL maintainer "Stanislav N. <pztrn@pztrn.name>"

COPY --from=build /go/src/sources.dev.pztrn.name/pztrn/giredore/cmd/giredored/giredored /usr/local/bin/giredored
COPY --from=build /go/src/sources.dev.pztrn.name/pztrn/giredore/cmd/giredorectl/giredorectl /usr/local/bin/giredorectl

EXPOSE 62222
ENTRYPOINT [ "/usr/local/bin/giredored" ]