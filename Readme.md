for i in *.mkv; do  ffmpeg -i  "$i" -c:v libx264 -c:a aac  "${i%.*}.mp4";  done

for i in *.mkv; do  ffmpeg -i  "$i" -c:v libx264 -ar 44100 -c:a aac "${i%.*}.mp4";  done

for i in *.mkv; do  HandBrakeCLI -Z "Chromecast 1080p30 Surround" -i "$i" -o "${i%.*}.mp4";  done