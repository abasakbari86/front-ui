#!/bin/sh

# Start fail2ban
[ $FRONTUI_ENABLE_FAIL2BAN == "true" ] && fail2ban-client -x start

# Run front-ui
exec /app/front-ui
