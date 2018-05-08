#!/bin/sh

pkill auth-oathd > /dev/null 2>&1 || true
nohup sh -c '/usr/local/bin/auth-oathd -config /etc/openvpn/users.json -socket /tmp/openvpn.sock -user nobody -group nogroup' >> /var/log/auth-oathd.log &
