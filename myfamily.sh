
#!/bin/bash

if [ -z "$HERO_ID" ]; then
  echo "HERO_ID is not set"
  exit 1
fi

curl -s https://01.gritlab.ax/assets/superhero/all.json \
| jq -r --arg id "$HERO_ID" '
  .[]
  | select(.id == ($id | tonumber))
  | .connections.relatives
  | gsub("^\"|\"$"; "")
' | tr '\n' '\\n'