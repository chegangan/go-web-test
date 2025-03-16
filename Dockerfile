FROM scratch

WORKDIR /go-web-test
COPY . /go-web-test

EXPOSE 8000
CMD ["./go-web-test"]