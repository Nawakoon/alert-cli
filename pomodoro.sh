#/bin/bash

timer_file="/tmp/pomodoro.txt"
wait_time=1500 # 25 minutes

if [ ! -f $timer_file ]; then
    echo "Starting pomodoro"
    now=$(date +%s)
    echo $(($now + $wait_time)) > $timer_file &&
    sleep $wait_time &&
    rm $timer_file &&
    osascript -e 'display dialog "Pomodoro done" with title "Time for a break"' &
else
    timer_end=$(cat $timer_file)
    timer_now=$(date +%s)
    time_left=$(($timer_end - $timer_now))
    minutes_left=$(($time_left / 60))
    seconds_left=$(($time_left % 60))
    echo "$minutes_left:$seconds_left min left"
fi

