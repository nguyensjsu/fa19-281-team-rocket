FROM golang:latest 
EXPOSE 3001
RUN mkdir /login 
ADD . /login/
WORKDIR /login 
ENV GOPATH /login
RUN cd src/login ; go get ; go install
CMD ["/login/bin/login"]

