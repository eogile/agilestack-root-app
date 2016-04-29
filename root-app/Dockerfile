FROM node:5.8
MAINTAINER EOGILE "agilestack@eogile.com"

ENV name=root-app \
    workdir=/plugin \
    HTTP_DIR=/shared/root-app
ENV baseUrl=/

EXPOSE 8081

LABEL SERVICE_TAGS="urlprefix-$baseUrl" \
      SERVICE_CHECK_HTTP="/stats" \
      SERVICE_CHECK_INTERVAL="10s"

WORKDIR $workdir

ADD $name $workdir/$name

ENTRYPOINT ["./root-app"]
