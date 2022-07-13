FROM kong/go-plugin-tool:latest-alpine-latest as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN	cd /app && go version && \
	go mod download

# COPY config.yaml .

COPY main.go .
RUN go build -o key-checker

FROM kong:latest
WORKDIR /app
COPY --from=builder /app/key-checker /usr/local/bin/key-checker
# COPY --chown=kong:kong --from=builder /app/config.yaml /tmp/config.yml
RUN stat /usr/local/bin/key-checker
RUN /usr/local/bin/key-checker -dump