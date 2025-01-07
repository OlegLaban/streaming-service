#!/bin/bash
while :
do
	ffmpeg -re -i $(./main "$1") -f lavfi -i anullsrc -c:v copy -c:a copy -f flv "rtmp://nginx/hls/live-$1"
    if (disaster-condition)
    then
        break       	   #Abandon the loop.
    fi
done