#!/bin/bash

curr=$(ps aux | grep XMU-Data | wc -l)
if [[ $curr == 1 ]]; then
    datetime=$(date "+%Y-%m-%d %H:%M:%S")
    echo "$datetime - Daemon Down... Rebooting" >> ~/cron.out
    nohup ~/go/bin/XMU-Data2 &
fi
