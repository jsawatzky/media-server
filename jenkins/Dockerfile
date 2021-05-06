FROM jenkins/jenkins:lts

ARG uid
ARG gid
ARG docker_gid

ENV JENKINS_OPTS --httpPort=80

USER root

RUN apt-get update && apt-get install -y libcap2-bin apt-transport-https \
       ca-certificates curl gnupg2 \
       software-properties-common
RUN curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add -
RUN apt-key fingerprint 0EBFCD88
RUN add-apt-repository \
       "deb [arch=amd64] https://download.docker.com/linux/debian \
       $(lsb_release -cs) stable"
RUN apt-get update && apt-get install -y docker-ce-cli

RUN setcap 'cap_net_bind_service=+ep' $(which java)
RUN ln -s /opt/java/openjdk/lib/amd64/jli/libjli.so /usr/lib
RUN ldconfig

RUN groupadd -g ${docker_gid} docker \
    && groupmod -g ${gid} jenkins \
    && usermod -aG docker -u ${uid} -g ${gid} jenkins \
    && chown ${uid}:${gid} /var/jenkins_home

USER jenkins

EXPOSE 80