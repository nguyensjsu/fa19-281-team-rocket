FROM golang:latest 
EXPOSE 3000
RUN mkdir /payments 
ADD . /payments/ 
WORKDIR /payments 
ENV GOPATH /payments
ENV AWS_ACCESS_KEY_ID=AKIASNOQELJK2QV4SCZK
ENV AWS_SECRET_ACCESS_KEY=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
RUN cd /payments ; go install payments
CMD ["/payments/bin/payments"]