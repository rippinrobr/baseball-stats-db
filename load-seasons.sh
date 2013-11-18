#!/bin/bash

NODE=`which node`

echo ""
echo ""
echo "### Starting to load the seasons data ###"

$NODE load-seasons.js
$NODE load-seasons-managers.js
$NODE load-half-seasons.js
$NODE load-rosters.js
$NODE load-playoff-records.js

echo "### Seasons data loaded! ###"
echo ""
echo ""

