# !/bin/sh

if [ $# -ne 1 ]; then
        echo "Usage: $0 file"
        exit 1;
fi

while : 
do
        echo "$(date +"%Y-%m-%d %H:%M:%S")   $1 was starting."
        $PWD/$1
        sleep 10
done
