#!/bin/bash

CHANGESET_OPTION="--no-execute-changeset"

if [ $# = 1 ] && [ $1 = "deploy" ]; then
  echo "deploy mode"
  CHANGESET_OPTION=""
fi

readonly CFN_TEMPLATE="$(dirname $0)/10-evidently.yml"
readonly CFN_STACK_NAME=FoodEvidently
readonly LAUNCH_START_AT="$(date -u +'%Y-%m-%dT%H:%M:%SZ')"

echo "CFN_TEMPLATE = ${CFN_TEMPLATE}"
echo "CFN_STACK_NAME = ${CFN_STACK_NAME}"
echo "LAUNCH_START_AT = ${LAUNCH_START_AT}"

aws cloudformation deploy \
--stack-name "${CFN_STACK_NAME}" \
--parameter-overrides LauchStartAt="${LAUNCH_START_AT}" \
--template-file "${CFN_TEMPLATE}" ${CHANGESET_OPTION}