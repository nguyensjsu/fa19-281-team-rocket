FROM golang:latest 
EXPOSE 3001
RUN mkdir /login 
ADD . /login/
WORKDIR /login 
ENV GOPATH /login
ENV AWS_ACCESS_KEY_ID=AKIASNOQELJK2QV4SCZK
ENV AWS_SECRET_ACCESS_KEY=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
RUN cd src/login ; go get ; go install
CMD ["/login/bin/login"]

