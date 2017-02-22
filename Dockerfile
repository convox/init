FROM heroku/cedar

ENV STACK=cedar-14 CURL_CONNECT_TIMEOUT=0 CURL_TIMEOUT=500 APP=/app

WORKDIR /tmp
RUN mkdir /tmp/cache

RUN  git clone https://github.com/heroku/heroku-buildpack-ruby    \
  && git clone https://github.com/heroku/heroku-buildpack-nodejs  \
  && git clone https://github.com/heroku/heroku-buildpack-clojure \
  && git clone https://github.com/heroku/heroku-buildpack-python  \
  && git clone https://github.com/heroku/heroku-buildpack-java    \
  && git clone https://github.com/heroku/heroku-buildpack-gradle  \
  && git clone https://github.com/heroku/heroku-buildpack-scala   \
  && git clone https://github.com/heroku/heroku-buildpack-php     \
  && git clone https://github.com/heroku/heroku-buildpack-go

COPY packs .
COPY bin/ /usr/local/bin
