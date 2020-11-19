#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
args="-confdir $CURDIR/conf"
exec $CURDIR/industry_identification_center $args