#!/bin/bash

curl -s https://01.gritlab.ax/assets/superhero/all.json |
jq -r --arg id "$HERO_ID" '
  .[]
  | select(.id == ($id | tonumber))
  | .connections.relatives
' | while IFS= read -r line; do
    printf "%b\n" "$line"
done
