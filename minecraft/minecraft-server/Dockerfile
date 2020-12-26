FROM openjdk:9-jre

RUN apt-get update && \
        apt-get upgrade --yes --force-yes && \
        apt-get clean && \
        rm -rf /var/lib/apt/lists/*

EXPOSE 25565

VOLUME ["/minecraft"]

WORKDIR "/minecraft"

CMD ["/bin/bash", "startserver.sh"]