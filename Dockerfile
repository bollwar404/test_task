FROM golang:latest
RUN mkdir /app/
RUN mkdir /app/src
RUN mkdir /app/bin
RUN mkdir /app/src/test_task
ADD . /app/src/test_task
ENV GOPATH=/app
ENV GIN_MODE=release
ENV PATH=$GOPATH/bin:$PATH:/app/src/test_task
WORKDIR /app/src/test_task
RUN curl https://glide.sh/get | sh
RUN glide install
RUN go build -o main_builded .
RUN ls /app/src/test_task
CMD ["/app/src/test_task/main_builded"]