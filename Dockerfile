FROM golang:1.5
RUN go get github.com/constabulary/gb/...
COPY . /go
ENV PATH=$PATH:/go/bin
RUN gb build all

