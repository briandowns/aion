FROM ubuntu

WORKDIR /opt/aion
ADD ./aion /opt/aion/aion
ADD ./public /opt/aion/public

ENTRYPOINT ["/opt/aion/aion"]
