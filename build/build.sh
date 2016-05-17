#!/bin/bash
# Shell script to determine what the current build of the
# repository is

##############################################################

CURRENT_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
JSON_PATH="${CURRENT_PATH}/version.json"
TEMPLATE_PATH="${CURRENT_PATH}/version.go.tmpl"
VERSION_PATH="${CURRENT_PATH}/../version.go"
GIT=`which git`
GO=`which go`

##############################################################
# Sanity checks

cd "${CURRENT_PATH}/.."
if [ ! -d ${CURRENT_PATH} ] ; then
    echo "Not found: ${CURRENT_PATH}"
fi
if [ ! -f ${TEMPLATE_PATH} ] ; then
  echo "Not found: ${TEMPLATE_PATH}"
fi
if [ ! -x ${GIT} ] ; then
  echo "Not found: ${GIT}"
fi
if [ ! -x ${GO} ] ; then
  echo "Not found: ${GO}"
fi


##############################################################
# Determine version

# get our version info from git
TAG=`${GIT} describe --tags`
BRANCH=`${GIT} name-rev HEAD --name-only --always`
HASH=`${GIT} rev-parse HEAD`
DATE=`date -u`
GOVERSION=`go version`

# output JSON with the information in it
echo "{" > ${JSON_PATH}
echo "  \"branch\":\"${BRANCH}\"," >> ${JSON_PATH}
echo "  \"tag\":\"${TAG}\"," >> ${JSON_PATH}
echo "  \"hash\":\"${HASH}\"," >> ${JSON_PATH}
echo "  \"date\":\"${DATE}\"," >> ${JSON_PATH}
echo "  \"goversion\":\"${GOVERSION}\"" >> ${JSON_PATH}
echo "}" >> ${JSON_PATH}

# build the command line tool
${GO} run build/build.go ${JSON_PATH} ${TEMPLATE_PATH} > ${VERSION_PATH}
${GO} install ytapi.go version.go

ls -l ${GOBIN}/ytapi

