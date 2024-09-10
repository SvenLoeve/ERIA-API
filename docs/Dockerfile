FROM golang:latest as openApiGen

ADD . .

RUN go run ./docs/openAPI

FROM registry.jetbrains.team/p/writerside/builder/writerside-builder:241.15989 as build

ARG INSTANCE=Writerside/eria-docs

RUN mkdir /opt/sources

WORKDIR /opt/sources

COPY --from=openApiGen /go/docs/openApi.json .

ADD ./docs/Writerside ./Writerside

RUN export DISPLAY=:99 && \
Xvfb :99 & \
/opt/builder/bin/idea.sh helpbuilderinspect -source-dir /opt/sources --product $INSTANCE --runner other --output-dir /opt/wrs-output/

WORKDIR /opt/wrs-output

RUN unzip webHelpERIA-DOCS2-all.zip -d /opt/wrs-output/unzipped-artifact

FROM httpd:2.4 as http-server

COPY --from=build /opt/wrs-output/unzipped-artifact/ /usr/local/apache2/htdocs/