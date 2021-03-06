#!/bin/bash

result=FAIL
trap 'echo $result' EXIT
set -xe

cd $(dirname $0)
f="$PWD/getcli.sh"

docker run -v "$f:/cli" --rm centos sh -c "sh /cli && backyards --version"

docker run -v "$f:/cli" --rm centos sh -c "sh </cli && backyards --version"
docker run -v "$f:/cli" --rm opensuse/leap sh -c "zypper install -y curl; sh </cli && backyards --version"
docker run -v "$f:/cli" --rm alpine sh -c "sh </cli && backyards --version"
docker run -v "$f:/cli" --rm ubuntu sh -c "apt-get update; env TERM=dumb apt-get -y install --no-install-recommends ca-certificates wget; sh </cli && backyards --version"
docker run -v "$f:/cli" --rm ubuntu:16.04 sh -c "apt-get update; env TERM=dumb apt-get -y install --no-install-recommends ca-certificates wget; sh </cli && backyards --version"
docker run -v "$f:/cli" --rm debian sh -c "apt-get update; env TERM=dumb apt-get -y install --no-install-recommends ca-certificates curl; sh </cli && backyards --version"
docker run -v "$f:/cli" --rm fedora sh -c "sh </cli && backyards --version"

docker run -v "$f:/cli" --rm centos sh -c "bash </cli && backyards --version"
docker run -v "$f:/cli" --rm ubuntu sh -c "apt-get update; env TERM=dumb apt-get -y install --no-install-recommends ca-certificates curl; bash </cli && backyards --version"

! docker run -v "$f:/cli" --rm ubuntu sh -c "bash </cli && backyards --version"

docker run -v "$f:/cli" --rm ubuntu sh -c "apt-get update; env TERM=dumb apt-get -y install --no-install-recommends ca-certificates curl; rm /etc/os-release; bash </cli && backyards --version"

result=OK
