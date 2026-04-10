#!/bin/sh

curl -s https://01.gritlab.ax/assets/superhero/all.json \
  | jq -r --arg id "$id" '.[$id|tonumber].connections.relatives'
