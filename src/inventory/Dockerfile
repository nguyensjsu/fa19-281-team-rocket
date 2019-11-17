FROM golang:latest
EXPOSE 3000
RUN mkdir /inventory
ADD . /inventory
WORKDIR /inventory
ENV GOPATH /inventory
RUN cd /inventory ; go install inventory
CMD ["/inventory/bin/inventory"]
