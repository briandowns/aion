FROM ubuntu

WORKDIR /opt/aion
ADD ./slam /opt/aion/aion
ADD ./public /opt/aion/public

ENTRYPOINT ["/opt/aion/aion"]
