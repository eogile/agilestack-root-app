FROM node:4.2.3
MAINTAINER EOGILE "agilestack@eogile.com"

ENV name=root-app-builder \
    workdir=/plugin \
    baseUrl=/api/root \
    HTTP_PORT=8080

ENV SOURCES_DIRECTORY=$workdir \
    HTTP_DIRECTORY=/shared/root-app/ \
    SHARED_MODULES_DIRECTORY=/shared/root-app-builder/web_modules \
    HTTP_APP_BASE_URL=/

EXPOSE $HTTP_PORT

LABEL SERVICE_TAGS="urlprefix-$baseUrl" \
      SERVICE_CHECK_HTTP="/status" \
      SERVICE_CHECK_INTERVAL="10s"

# Adding the ReactJS application sources
ADD node_modules $workdir/node_modules/
ADD css $workdir/css
ADD img $workdir/img
ADD js  $workdir/js
ADD index.html .babelrc .eslintrc .htaccess makewebpackconfig.js manifest.json package.json server.prod.js webpack.prod.config.js $workdir/

WORKDIR $workdir
ADD $name $workdir/$name

ENTRYPOINT ["/plugin/root-app-builder"]
