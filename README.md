TRD(8) - System Manager's Manual

# NAME

**trd** - TFTP Rewrite Daemon

# SYNOPSIS

**trd**
\[**-c**&nbsp;*string*]
\[**-debug**]
\[**-s**&nbsp;*string*]
\[**-u**&nbsp;*string*]

# DESCRIPTION

OpenBSD's
tftpd(1)
supports on-the-fly rewriting of file paths based on IP address.
It does this by sending the tftp request lines over a socket that is
specified via the
**-r**
option.

**trd**
creates a socket and expects
tftpd(1)
to send lines in the format of "IP OP filename".

Hosts are matched by IP and requested filename.

The options are as follows:

**-c** *config*

>Path of to config file.
>Default is
>*/etc/trd.conf*.

**-debug**

>Tells
>**trd**
>to print debug messages.

**-s** *socket*

>Specifies the path to the socket
>**trd**
>should create.
>This needs to be passed to
>tftpd(1)
>via the
>**-r**
>option.

**-u** *user*

>User
>**trd**
>should drop privlidges to.
>Default is \_tftpd.

# SEE ALSO

tftpd(1),
trd.conf(5)

OpenBSD 6.1 - May 1, 2016
---
ACME-CLIENT.CONF(5) - File Formats Manual

# NAME

**trd.conf** - trd configuration file

# DESCRIPTION

The
**trd.conf**
config file is split into three columns, delimited by a single space:

**IP**

>Source IP address of the request.

**Requested File**

>File that the source IP is requesting.

**Rewrite File**

>File that the host should use instead of the requested file.

For example:

	10.33.0.12 /etc/boot.conf /i386/snapshots/boot.conf
	10.33.0.12 /bsd /i386/snapshots/bsd.rd

# FILES

*/etc/trd.conf*

>trd(8)
>configuration file

# SEE ALSO

trd(8)

OpenBSD 6.1 - May 1, 2016
