FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.23.4 AS build

ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

WORKDIR /collector-core
COPY . .

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build

FROM --platform=${TARGETPLATFORM:-linux/amd64} ubuntu:20.04 AS ship

RUN groupadd app
RUN useradd -g app app

RUN apt-get update
RUN apt install -y python3 python3-pip git

ARG USER
ARG PASS
ARG BUCKETNAME
ARG ACCOUNTID
ARG ACCOUNTSECRET

WORKDIR /home/app

EXPOSE 9090

ENV CF_USER=$USER
ENV CF_PASSWORD=$PASS
ENV R2_BUCKETNAME=$BUCKETNAME
ENV R2_ACCOUNTID=$ACCOUNTID
ENV R2_ACCOUNTSECRET=$ACCOUNTSECRET

COPY --from=build /collector-core .
COPY --from=build /collector-core/collector.yaml .

RUN chown -R app:app ./

RUN git clone https://github.com/siyak2/scamper.git && \ 
    cd scamper/scamper-cvs-20250505 && \
    ./configure && \
    make && \
    make install && \
    cd .. && \
    cd .. && \
    rm -rf scamper

RUN mkdir -p /home/app/data

WORKDIR /home/app

# USER app  
CMD ["./changelog"]
