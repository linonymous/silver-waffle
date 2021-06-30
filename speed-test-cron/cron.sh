#!/bin/sh

# export PATH=$PATH:/home/linonymous/go/bin/speed-test-cron

/home/linonymous/go/bin/speed-test-cron

# cron to run every 10 minutes
# crontab -e
# */10 * * * * /home/linonymous/go/src/speed-test-cron/cron.sh >> /home/go/src/speed_test.log 2>&1