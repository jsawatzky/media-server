FROM openjdk:9-jre

ARG uid
ARG gid

USER root

RUN apt-get update && \
        apt-get upgrade --yes --force-yes && \
        apt-get clean && \
        rm -rf /var/lib/apt/lists/*

RUN mkdir -p /minecraft \
    && chown ${uid}:${gid} /minecraft \
    && groupadd -g ${gid} minecraft \
    && useradd -d "/minecraft" -u ${uid} -g ${gid} -s /bin/bash minecraft

USER minecraft

EXPOSE 25565 25575

VOLUME ["/minecraft"]

WORKDIR "/minecraft"

CMD ["/bin/bash", "startserver.sh"]