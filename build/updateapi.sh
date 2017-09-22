#!/bin/bash
##############################################################
# YTAPI UPDATE API SCRIPT
##############################################################
# Shell script to update the YouTube API codes to the latest
# versions, including downloading the Partner API and
# generating the go library.
#
# Requires `curl` in order to download the Partner API from
# remote repository. There are no arguments to this script,
# so run using:
# 
# updateapi.sh
#
# You will then need to run the `build.sh` script in order
# to create the binary
##############################################################

CURRENT_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
CURL=`which curl`
GO=`which go`

##############################################################

cd "${CURRENT_PATH}/.."

##############################################################
# Sanity checks

if [ ! -d ${CURRENT_PATH} ] ; then
  echo "Not found: ${CURRENT_PATH}" >&2
  exit -1
fi
if [ ! -x ${GO} ] ; then
  echo "go not installed or executable" >&2
  exit -1
fi
if [ ! -x ${CURL} ] ; then
  echo "curl not installed or executable" >&2
  exit -1
fi

##############################################################
# Fetch API

# Get the API generator from remote source and build it
go get -u google.golang.org/api/google-api-go-generator || exit 1

# Create the youtubepartner/v1 folder
install -d youtubepartner/v1 || exit 1

#  Fetch the partner API code
curl https://www.googleapis.com/discovery/v1/apis/youtubePartner/v1/rest > youtubepartner/v1/youtubepartner.json || exit 1

##############################################################
# Generate Go Libraries

API_GEN=`which google-api-go-generator`
if [ ! -x ${API_GEN} ] ; then
  echo "google-api-go-generator not installed or executable" >&2
  exit -1
fi

# Generate the go library
google-api-go-generator -cache=false -gendir=. -api_json_file=youtubepartner/v1/youtubepartner.json -api_pkg_base=github.com/djthorpe/ytapi

