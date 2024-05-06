#! /bin/bash

id=$HERO_ID

curl -s https://01.gritlab.ax/assets/superhero/all.json | jq ' .[] | select( .id == '$id') | .connections | .relatives' -r