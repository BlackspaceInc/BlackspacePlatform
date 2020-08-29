FROM docker.io/ubuntu:20.04 AS ENVIRONMENT
ENV DEBIAN_FRONTEND=noninteractive \
    TZ='America/New York' \
    PATH=/root/.cargo/bin:/root/.local/bin/:$PATH \
    LANG=en_US.UTF-8 \
    LANGUAGE=en_US.UTF-8 \
    LC_ALL=en_US.UTF-8

# Container junk
RUN echo $TZ > /etc/timezone

# Setup the env
RUN mkdir -p /git/timberio/Blackspace/{scripts,scripts/environment}
ADD scripts/environment/bootstrap-ubuntu-20.04.sh scripts/environment/prepare.sh \
    /git/timberio/Blackspace/scripts/environment/

# Setup the toolchain
WORKDIR /git/timberio/Blackspace
RUN ./scripts/environment/bootstrap-ubuntu-20.04.sh

ADD scripts/Gemfile scripts/Gemfile.lock \
    /git/timberio/Blackspace/scripts/
ADD rust-toolchain \
    /git/timberio/Blackspace/
RUN ./scripts/environment/prepare.sh

# Declare volumes
VOLUME /Blackspace
VOLUME /Blackspace/target
VOLUME /root/.cargo

# Prepare for use
ADD ./scripts/environment/entrypoint.sh /
ENTRYPOINT [ "/entrypoint.sh" ]
CMD [ "bash" ]
