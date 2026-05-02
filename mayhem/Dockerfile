FROM golang:1.25 AS builder

COPY . /cute
WORKDIR /cute

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN GOFLAGS=-mod=mod go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o mayhem/fuzzCute.a ./mayhem && \
    clang -fsanitize=fuzzer mayhem/fuzzCute.a -o mayhem/fuzzCute.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /cute/mayhem/fuzzCute.libfuzzer /
