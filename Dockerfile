FROM heroku/cedar

ENV CURL_CONNECT_TIMEOUT=0 STACK=cedar-14

WORKDIR /tmp

RUN git clone https://github.com/heroku/heroku-buildpack-ruby
RUN git clone https://github.com/heroku/heroku-buildpack-python

COPY bin/ /usr/local/bin