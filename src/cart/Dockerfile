FROM golang:latest
EXPOSE 3000
RUN mkdir /cart
ADD . /cart
WORKDIR /cart
ENV GOPATH /cart
RUN cd /cart ; go install cart
CMD ["/cart/bin/cart"]
