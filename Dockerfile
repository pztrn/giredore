FROM code.pztrn.name/containers/mirror/golang:1.18.3-alpine AS build

WORKDIR /go/src/sources.dev.pztrn.name/pztrn/giredore
COPY . .

ENV CGO_ENABLED=0
RUN cd /go/src/sources.dev.pztrn.name/pztrn/giredore/cmd/giredored && go build -tags netgo -ldflags '-w -extldflags "-static"' && cd ../giredorectl && go build -tags netgo -ldflags '-w -extldflags "-static"'

FROM code.pztrn.name/containers/mirror/alpine:3.16.0
LABEL maintainer "Stanislav N. <pztrn@pztrn.name>"

COPY --from=build /go/src/sources.dev.pztrn.name/pztrn/giredore/cmd/giredored/giredored /usr/local/bin/giredored
COPY --from=build /go/src/sources.dev.pztrn.name/pztrn/giredore/cmd/giredorectl/giredorectl /usr/local/bin/giredorectl

EXPOSE 62222
ENTRYPOINT [ "/usr/local/bin/giredored" ]
