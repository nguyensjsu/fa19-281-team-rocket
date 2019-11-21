FROM golang:latest 
EXPOSE 3000
RUN mkdir /payments 
ADD . /payments/ 
WORKDIR /payments 
ENV GOPATH /payments
RUN cd /payments ; go install payments
CMD ["/payments/bin/payments"]