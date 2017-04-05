TRD(8) - System Manager's Manual

# NAME

**trd** - TFTP Rewrite Daemon

# SYNOPSIS

**trd**
\[**-csu**]
\[**-c**&nbsp;\[config]]
\[**-s**&nbsp;\[socket]]
\[**-u**&nbsp;\[user]]
\[**-debug**]

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

> Path to config file.
> Default is
> */etc/trd.conf*.

**-s** *socket*

> Specifies the path to the socket
> **trd**
> should create.
> This needs to be passed to
> tftpd(1)
> via the
> **-r**
> option.

**-u** *user*

> User
> **trd**
> should drop privlidges to.
> Default is \_tftpd.

**-debug**

> Tells
> **trd**
> to print debug messages.

# SEE ALSO

tftpd(1),
trd.conf(5)

# HISTORY

The first version of
**trd**
was released in May of 2016.

# AUTHORS

**trd**
was written by
Aaron Bieber &lt;[aaron@bolddaemon.com](mailto:aaron@bolddaemon.com)&gt;.

OpenBSD 6.1 - May 1, 2016
