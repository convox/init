FROM heroku/cedar

ENV CURL_CONNECT_TIMEOUT=0 STACK=cedar-14

WORKDIR /tmp

RUN  git clone https://github.com/heroku/heroku-buildpack-ruby    \
  && git clone https://github.com/heroku/heroku-buildpack-nodejs  \
  && git clone https://github.com/heroku/heroku-buildpack-clojure \
  && git clone https://github.com/heroku/heroku-buildpack-python  \
  && git clone https://github.com/heroku/heroku-buildpack-java    \
  && git clone https://github.com/heroku/heroku-buildpack-gradle  \
  && git clone https://github.com/heroku/heroku-buildpack-scala   \
  && git clone https://github.com/heroku/heroku-buildpack-php     \
  && git clone https://github.com/heroku/heroku-buildpack-go

COPY bin/ /usr/local/bin