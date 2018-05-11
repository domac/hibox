#!/bin/bash
sed -i '/agent/d' /etc/rc.d/rc.local
echo "cd /apps/svr/domac/src/agent && sh shutdown.sh && sh startup.sh" >> /etc/rc.d/rc.local
ldconfig