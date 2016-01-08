FROM debian:8.2
ENV DEBIAN_FRONTEND "noninteractive"
ENV APT_FLAGS "-o Dpkg::Options::=--force-confold --assume-yes --quiet"
RUN apt-get $APT_FLAGS update && apt-get $APT_FLAGS install \
    python \
    python-pip \
    python-flake8
RUN pip install \
    requests \
    retrying \
    pytest
WORKDIR /mnt/factbeat
