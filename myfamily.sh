#!/bin/bash

# Check if HERO_ID is set
if [ -z "$HERO_ID" ]; then
  echo "HERO_ID is not set"
  exit 1
fi

# Fetch data and extract relatives
curl -s https://01.gritlab.ax/assets/superhero/all.json \
| jq -r --arg id "$HERO_ID" '
  .[] 
  | select(.id == ($id | tonumber)) 
  | .connections.relatives
  | gsub("^\"|\"$"; "")
'
