FROM golang:1.18 as builder

COPY . /cute
WORKDIR /cute/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzzCute.a && \
    clang -fsanitize=fuzzer fuzzCute.a -o fuzzCute.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /cute/mayhem/fuzzCute.libfuzzer /