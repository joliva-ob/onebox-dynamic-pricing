#
# Dockerfile for Dynamic Pricing API
#

FROM golang:onbuild

MAINTAINER Joan Oliva

RUN mkdir /dynamic-pricing
RUN mkdir /dynamic-pricing/bin
RUN mkdir /dynamic-pricing/cfg
RUN mkdir /dynamic-pricing/logs
ADD *.yml /dynamic-pricing/cfg
ADD *.gcfg /dynamic-pricing/cfg

RUN GOOS=linux GOARCH=386 go build -o dynamic-pricing.linux dynamic-pricing.go
ADD dynamic-pricing.linux /dynamic-pricing/bin

WORKDIR /dynamic-pricing

CMD ["/dynamic-pricing.linux /cfg pre"]

EXPOSE 8000

# COPY run.sh /run.sh
# RUN chmod 777 /run.sh

# CMD "/run.sh"
