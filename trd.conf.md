TRD.CONF(5) - File Formats Manual

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
