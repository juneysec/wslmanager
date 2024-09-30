FROM ubuntu:24.04

RUN apt-get -qq update \
 && apt-get -qqy --no-install-recommends install apt-utils \
 && apt-get -qqy --no-install-recommends install build-essential \
    openjdk-17-jdk \
    openjdk-17-jre-headless \
    zip \
    unzip \
    sudo \
    curl \
    lldb \
    vim \
    language-pack-ja-base \
    language-pack-ja ibus-mozc \
    jq \
    maven \
    wget \
    git > /dev/null \
 && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* \
 && useradd -ms /bin/bash developer \
 && adduser developer root \
 && echo "developer ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

ENV LANG=ja_JP.UTF-8
ENV LANGUAGE="ja_JP:ja"
ENV TZ="Asia/Tokyo"

USER developer
WORKDIR /home/developer
RUN wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/7.8.0/openapi-generator-cli-7.8.0.jar -O openapi-generator-cli.jar \
 && curl https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/bin/utils/openapi-generator-cli.sh > /home/developer/openapi-generator-cli \
 && chmod u+x /home/developer/openapi-generator-cli \
 && /home/developer/openapi-generator-cli; exit 0

