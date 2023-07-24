# build stage
FROM golang:1.20.6 as build-env

ENV REPO="github.com/dca/chatbot-ai"
ENV CMD="message"
ENV GO111MODULE=on

WORKDIR /go/src/${REPO}

COPY go.mod /go/src/${REPO}/go.mod
COPY go.sum /go/src/${REPO}/go.sum
RUN go mod download

# COPY cmd/${CMD} /go/src/${REPO}/cmd/${CMD}
# COPY internal/app/${CMD} /go/src/${REPO}/internal/app/${CMD}
# COPY pkg /go/src/${REPO}/pkg
COPY cmd /go/src/${REPO}/cmd
COPY internal /go/src/${REPO}/internal

RUN env GOOS=linux GOARCH=amd64 go build -o /handler -v -ldflags "-s" ${REPO}/cmd/${CMD}
 
# ####################################################################

FROM alpine:3.18.2 AS runtime

FROM alpine:3.18.2
RUN apk --no-cache add ca-certificates libc6-compat \
    && addgroup -S app && adduser -S -g app app \
    && mkdir -p /home/app \
    && chown app /home/app

WORKDIR /home/app

COPY --from=build-env /handler .

RUN chown -R app /home/app
USER 1000

CMD ["./handler"]

EXPOSE 8081