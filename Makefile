VERSION =	`git describe --tags 2>/dev/null || git log -n 1 --format="%h"`
PREFIX ?=	/usr/local
MAN =		trd.8 trd.conf.5
MANDIR ?=	${PREFIX}/man/man

lint:
	mandoc -T lint trd.8
	mandoc -T lint trd.conf.5

trd.8: lint
	mandoc -T lint trd.8
	mandoc -T ascii trd.8 >$@

trd.conf.5: lint
	mandoc -T lint trd.conf.5
	mandoc -T ascii trd.conf.5 >$@

README.md: lint
	mandoc -T markdown trd.8 >$@
	mandoc -T markdown trd.conf.5 > trd.conf.md

build:
	go build -ldflags "-X main.version=${VERSION}"

realinstall: build
	install -d $(DESTDIR)$(PREFIX)/bin
	install -m 755 trd $(DESTDIR)$(PREFIX)/bin/trd

.include <bsd.prog.mk>
