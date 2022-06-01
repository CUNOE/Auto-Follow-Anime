#!/bin/bash
apt-get update
apt-get install cron -y
sed -i 's/required/sufficient/' /etc/pam.d/cron
echo "0 */1 * * * root /root/app" >> /etc/crontab
/etc/init.d/cron restart