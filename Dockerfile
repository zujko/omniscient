FROM centos

ENV omni /home/omni
WORKDIR ${omni}

ADD . ${omni}
