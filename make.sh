#!/usr/bin/bash

RELEASE=${GOPATH}/release
VERSION=$(git tag | tail -n 1)

function check_err() {
    if [[ $? != 0 ]]
    then
        exit $?
    fi
}

pushd utilities/sysmon && make -j 8 $@ || check_err && popd

pushd utilities/whids && make -j 8 $@ || check_err && popd

pushd utilities/manager && make -j 8 $@ || check_err && popd

pushd ${RELEASE}
# Remove previous bundles
rm *.zip
7z a -tzip whids-${VERSION}-release-bundle.zip *

popd


