FROM golang:1.15-alpine AS base
RUN apk add ca-certificates \
    && rm -rf /var/cache/apk/*
    
FROM base as deps
WORKDIR "/hello"
ADD *.mod *.sum ./
RUN go mod download

FROM deps AS build-env
COPY cmd ./cmd
copy hello ./hello
ENV PORT 8080
EXPOSE 8080
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -X main.docker=true" -o server cmd/server/main.go
CMD ["./server"]

FROM scratch AS prod

WORKDIR /
ENV PORT 8080
EXPOSE 8080

COPY --from=build-env /hello/server /
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
CMD ["/server"]