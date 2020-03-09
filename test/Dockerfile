FROM alpine:3.6

# 设置locale
ENV LANG en_US.UTF-8
ENV LANGUAGE en_US:en
ENV LC_ALL en_US.UTF-8
ENV TZ=Asia/Shanghai

RUN mkdir /voice_test

WORKDIR /voice_test

COPY test /voice_test

RUN chmod +x /voice_test/test

ENV CLIENT=/voice_test/test
ENV PATH $CLIENT:$PATH
