#
# Dockerfile for Dynamic Pricing API
#

FROM golang:onbuild

MAINTAINER Joan Oliva

RUN mkdir /dynamic-pricing
RUN mkdir /dynamic-pricing/bin
RUN mkdir /dynamic-pricing/cfg
RUN mkdir /dynamic-pricing/logs

ADD *.yml /dynamic-pricing/cfg/
ADD *.gcfg /dynamic-pricing/cfg/

ENV CONF_PATH /dynamic-pricing/cfg
ENV ENV pro

WORKDIR /dynamic-pricing
ENTRYPOINT /dynamic-pricing/bin/dynamic-pricing.linux

EXPOSE 8000