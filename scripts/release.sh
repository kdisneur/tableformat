#! /usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

_applicationName="tableformat";

_rootFolder=$(cd $(dirname ${0})/..; pwd);
_configFile=${1:-~/.config/github};
_version=$(cat VERSION);
_currentSHA=$(git rev-parse HEAD);

if [ -z ${GITHUB_API_TOKEN:-} ]; then
  if [ ! -f ${_configFile} ]; then
    &>2 echo "You must have a GITHUB_API_TOKEN environment variable or ~/.config/github file available"
    exit 1
  else
    GITHUB_API_TOKEN=$(cat ${_configFile});
  fi
fi

_release=$(
  curl \
    --silent \
    --request POST \
    --header "Authorization: bearer ${GITHUB_API_TOKEN}" \
    --header "Content-Type: application/json" \
    --data '
      { "tag_name": "'${_version}'"
      , "target_commitish": "'${_currentSHA}'"
      , "body": "Release '${_version}'"
      , "name": "'${_version}'"
      , "draft": false
      , "prerelease": false
      }' \
    https://api.github.com/repos/kdisneur/tableformat/releases);

_uploadUrl=$(
  grep -oE "upload_url\":[ ]+\"([^\"]+)\"" <<< ${_release} \
    | grep -oE 'https:.*assets');

curl \
  --silent \
  --request POST \
  --header "Authorization: bearer ${GITHUB_API_TOKEN}" \
  --header "Content-Type: application/octet-stream" \
  --data-binary @${_rootFolder}/${_applicationName} \
  "${_uploadUrl}?name=${_applicationName}&label=Executable" > /dev/null;
