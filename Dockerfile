FROM bufbuild/buf:1.72.0 AS buf

FROM golang:1.25-alpine

ENV GOPROXY=https://proxy.golang.org,https://goproxy.cn/,direct

# Install protobuf compiler and git (needed for some go modules)
RUN apk add --no-cache protobuf git

# Copy the pre-built buf binary from the official image
COPY --from=buf /usr/local/bin/buf /usr/local/bin/buf

# Install protoc-gen-go
ENV PATH="/root/go/bin:${PATH}"
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

WORKDIR /workspace
COPY . .

RUN buf generate
