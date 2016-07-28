#!/bin/bash
##############################################################
# YTAPI BUILD SCRIPT
##############################################################
# Shell script to determine what the current build of the
# repository is, and encode it into a go file - then compile
# the ytapi binary. To use the command, you can run the
# shell without any flags, but you can also wrap a client
# secrets and service account file into the binary so that
# it's easy to install these by running the ytapi binary
# without having to do lots of copying afterwards. So:
#
# build.sh
# build.sh -c <client_secret.json>
# build.sh -c <client_secret.json> -s <service_account.json>
#
# ...are all valid ways of compiling a binary of the software
# the resulting binary will be in ${GOBIN} and called 'ytapi'
# which you can then distribute!
#
# To build for Windows architecture, use with  the -w flag
# which results in a ytapi.exe application instead.
#
##############################################################

if [ -z "${TMPDIR}" ]; then
  echo "TMPDIR is unset or set to the empty string."
  TMPDIR=`dirname $(mktemp -u -t tmp.XXXXXXXXXX)`
  echo "Now set to: ${TMPDIR}"
fi
CURRENT_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
JSON_PATH="${TMPDIR}/version.json"
TEMPLATE_PATH="${CURRENT_PATH}/version.go.tmpl"
VERSION_PATH="${CURRENT_PATH}/../version.go"
GIT=`which git`
GO=`which go`

##############################################################

cd "${CURRENT_PATH}/.."

##############################################################
# Sanity checks

if [ ! -d ${TMPDIR} ] ; then
  echo "Temporary directory $TMPDIR not found" >&2
  exit -1
fi
if [ ! -d ${CURRENT_PATH} ] ; then
  echo "Not found: ${CURRENT_PATH}" >&2
  exit -1
fi
if [ ! -f ${TEMPLATE_PATH} ] ; then
  echo "Template not found: ${TEMPLATE_PATH}" >&2
  exit -1
fi
if [ ! -x ${GIT} ] ; then
  echo "git not installed or executable" >&2
  exit -1
fi
if [ ! -x ${GO} ] ; then
  echo "go not installed or executable" >&2
  exit -1
fi

##############################################################
# Check for client secret (-c) and/or service account flag (-s)
# which include the client secret info and service account
# info when generating the version.json file

while getopts ':c:s:w' FLAG ; do
  case ${FLAG} in
    c)
	  CLIENT_SECRET=${OPTARG}
      ;;
    s)
	  SERVICE_ACCOUNT=${OPTARG}
      ;;
    w)
	  GOOS="windows"
	  GOARCH="amd64"
      ;;
    \?)
      echo "Invalid option: -$OPTARG"
	  exit -1
      ;;
    :)
      echo "Option -$OPTARG requires an argument"
      exit -1
      ;;
  esac
done

if [ "${CLIENT_SECRET}" != "" ] && [ ! -f ${CLIENT_SECRET} ] ; then
  echo "Not found: ${CLIENT_SECRET}" >&2
  exit -1
fi

if [ "${SERVICE_ACCOUNT}" != "" ] && [ ! -f ${SERVICE_ACCOUNT} ] ; then
  echo "Not found: ${SERVICE_ACCOUNT}" >&2
  exit -1
fi

##############################################################
# Determine version

# get our version info from git
TAG=`${GIT} describe --tags 2>/dev/null`
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
echo "  \"goversion\":\"${GOVERSION}\"," >> ${JSON_PATH}
if [ "${GOOS}X" != "X" ] ; then
	echo "  \"target_os\":\"${GOOS}\"," >> ${JSON_PATH}
fi
if [ "${GOARCH}X" != "X" ] ; then
	echo "  \"target_arch\":\"${GOARCH}\"," >> ${JSON_PATH}
fi
echo "  \"client_secret\":\"${CLIENT_SECRET}\"," >> ${JSON_PATH}
echo "  \"service_account\":\"${SERVICE_ACCOUNT}\"" >> ${JSON_PATH}
echo "}" >> ${JSON_PATH}

# get dependencies
${GO} get

# build the command line tool
${GO} run build/build.go ${JSON_PATH} ${TEMPLATE_PATH} > ${VERSION_PATH}
if [ $? -ne 0 ] ; then
  exit $?
fi
GOOS=${GOOS} GOARCH=${GOARCH} ${GO} install -ldflags "-s -w" ytapi.go version.go
if [ $? -ne 0 ] ; then
  exit $?
fi

# output
cat ${JSON_PATH}

