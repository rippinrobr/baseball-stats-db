#!/bin/bash

NODE=`which node`

echo ""
echo ""
echo "### Starting to load the player data ###"

$NODE load-plyr-demog.js
$NODE load-allstars.js
$NODE load-appearances.js
$NODE load-batting.js
$NODE load-fielding.js
$NODE load-pitching.js
$NODE load-players-awards.js
$NODE load-positions.js
$NODE load-postseason-batting.js
$NODE load-postseason-fielding.js
$NODE load-postseason-pitching.js
$NODE load-salaries.js
$NODE create-player-stat-indexes.js

echo "### Player data loaded! ###"
echo ""
echo ""

