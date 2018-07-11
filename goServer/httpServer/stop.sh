#! /bin/bash

gamename="httpServer.linux"

kill $(ps aux | grep ${gamename}$ | grep daemon.sh | grep -v grep | awk '{print $2}') 2>/dev/null
killall -15 $gamename
