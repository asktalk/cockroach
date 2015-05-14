#!/bin/bash

set -ex

if [ "$1" = "docker" ]; then
    time go install github.com/cockroachdb/build-cache
    time build-cache -race restore .
    time go test -race -v -i ./...
    time build-cache -race save .

    # Any package which uses stringer needs to be installed. Not quite
    # sure why, but stringer fails otherwise.
    pkgs=$(git grep -l 'go:generate stringer' | grep '\.go$' | xargs -n1 dirname | awk '{printf "./%s\n", $1}')
    for pkg in ${pkgs}; do
	time build-cache restore "${pkg}"
    done
    
    time go install ${pkgs}
    
    for pkg in ${pkgs}; do
	time build-cache save "${pkg}"
    done

    cmds=$(grep '^cmd' GLOCKFILE | grep -v glock | awk '{print $2}')
    for cmd in ${cmds}; do
	time build-cache restore "${cmd}"
    done

    time go install -v ${cmds}

    for cmd in ${cmds}; do
	time build-cache save "${cmd}"
    done
    exit 0
fi

git remote -v

gopath0="${GOPATH%%:*}"
cachedir="${gopath0}/pkg/cache"
tag="20150519.3857"

mkdir -p "${cachedir}"
du -sh "${cachedir}"
ls -lh "${cachedir}"

if ! docker images | grep -q "${tag}"; then
    # If there's a base image cached, load it. A click on CircleCI's "Clear
    # Cache" will make sure we start with a clean slate.
    rm -f "${cachedir}/builder.tar"
    builder="${cachedir}/builder.${tag}.tar"
    if [[ ! -e "${builder}" ]]; then
	time docker pull "cockroachdb/builder:${tag}"
	time docker save "cockroachdb/builder:${tag}" > "${builder}"
    else
	time docker load -i "${builder}"
    fi
fi

HOME= go get -d -u github.com/cockroachdb/build-cache
HOME= go get -u github.com/robfig/glock
grep -v '^cmd' GLOCKFILE | glock sync -n

# Pretend we're already bootstrapped, so that `make` doesn't try to get us
# started which is impossible without a working Go env.
touch .bootstrap && make '.git/hooks/*'

# Recursively invoke this script inside the builder docker container,
# passing "docker" as the first argument.
$(dirname $0)/circle-build.sh $0 docker

du -sh "${cachedir}"
ls -lh "${cachedir}"
