FROM --platform=linux/amd64 ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
  apt-get install -y \
  ca-certificates \
  curl \
  gcc-mingw-w64 \
  gcc \
  libc6 \
  libc6-dev \
  libxml2 \
  libxml2-dev \
  pkg-config

RUN curl -Lf "https://go.dev/dl/go1.18.3.linux-amd64.tar.gz" -o "go.tar.gz" \
  && tar -C /usr/local -xzf go.tar.gz \
  && rm -f go.tar.gz \
  && mkdir -p /home/test

ENV PATH=/usr/local/go/bin:$PATH

WORKDIR /home/test
COPY . ./
RUN go mod tidy
