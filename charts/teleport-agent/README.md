# Teleport daemon agent

This chart presents how to use a script inside the pod in order to configure teleport daemon on a host machine.
The script was successfully tested on Ubuntu 20.04 machine.

### Known issues
Installing on systems with GNU libc older than 2.18 (e.g. Centos 7) requires usage of dedicated teleport binary. The script does not support it for now.
