FROM debian:8.2
ENV DEBIAN_FRONTEND "noninteractive"
ENV DEB "puppetlabs-release-pc1-trusty.deb"
ENV APT_FLAGS "-o Dpkg::Options::=--force-confold --assume-yes --quiet"
ENV APT_GET "apt-get ${APT_FLAGS}"
ENV PATH /opt/puppetlabs/bin:$PATH
RUN $APT_GET update && $APT_GET install wget
RUN wget --quiet -O /tmp/${DEB} https://apt.puppetlabs.com/$DEB
RUN dpkg -i /tmp/$DEB
RUN $APT_GET update && $APT_GET install puppet-agent
CMD /mnt/factbeat/factbeat -e -d beat -c /mnt/factbeat/factbeat.yml
