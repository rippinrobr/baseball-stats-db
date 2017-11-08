#!/bin/bash

NODE=`which node`

echo ""
echo ""
echo "### Starting to load the seasons data ###"

$NODE load-teams.js
$NODE load-teams-managers.js
$NODE load-half-seasons.js
$NODE load-playoff-records.js

echo "### Seasons data loaded! ###"
echo ""
echo ""

