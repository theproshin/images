#!/bin/bash

ARGS=("$@")
PROFILE=""

for ((i=0; i<$#; i++)); do
  if [[ "${ARGS[$i]}" == "-profile" ]]; then
    PROFILE="${ARGS[$((i+1))]}"
    break
  fi
done

if [ -n "$PROFILE" ]; then
  echo "[firefox-profile] injecting user.js into $PROFILE" >&2
  cp /opt/firefox-user.js "$PROFILE/user.js"
fi

exec /usr/bin/firefox-real "${ARGS[@]}"