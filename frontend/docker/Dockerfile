FROM ubuntu:24.04

RUN apt-get -qq update \
 && apt-get -qqy --no-install-recommends install apt-utils \
 && apt-get -qqy --no-install-recommends install build-essential \
    sudo \
    language-pack-ja-base \
    language-pack-ja ibus-mozc \
    wget \
    vim \
    curl \
    nodejs \
    npm \
 && useradd -ms /bin/bash developer \
 && adduser developer root \
 && echo "developer ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

RUN npm install n -g \
 && n stable \
 && apt purge -y nodejs npm \
 && sudo apt autoremove -y

ENV LANG=ja_JP.UTF-8
ENV LANGUAGE="ja_JP:ja"
ENV TZ="Asia/Tokyo"
#ENV DISPLAY=host.docker.internal:0.0

USER developer
WORKDIR /home/developer

# src ディレクトリをバインドする場合
RUN echo sudo chmod 770 /home/developer/src >> .bashrc


