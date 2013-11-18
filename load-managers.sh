#!/bin/bash

NODE=`which node`

echo ""
echo ""
echo "### Starting to load the manager data ###"

$NODE load-managers.js
$NODE load-mgr-demog.js
$NODE load-half-mgrs.js
$NODE load-manager-awards.js

echo "### Managers data loaded! ###"
echo ""
echo ""

