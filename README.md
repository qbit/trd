# trd

TFTP Rewrite Daemon

## About

OpenBSD's TFTPD(8) supports request rewriting via a domain socket. This
app takes a config file in the format of "IP OP filename", where
filename is the file to be rewritten for a given IP.

