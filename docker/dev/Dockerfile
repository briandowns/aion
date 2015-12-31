FROM centos:6

RUN curl -L -o go1.5.2.linux-amd64.tar.gz https://storage.googleapis.com/golang/go1.5.2.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.5.2.linux-amd64.tar.gz

# node setup
RUN yum install -y nodejs-0.10.36 \
                   gcc-c++-4.4.7 \
                   gcc-4.4.7 \
                   http-parser-2.0 \
                   libuv-0.10.34 \
                   ruby-1.8.7.374 \
                   rubygems-1.3.7 \
                   ruby-devel-1.8.7.374