VERSION =	`git describe --tags 2>/dev/null || git log -n 1 --format="%h"`
PREFIX ?=	/usr/local
MAN =		trd.8 trd.conf.5
MANDIR ?=	${PREFIX}/man/man

clean:
	rm -f trd ${MAN}

lint:
	mandoc -T lint man/trd.8
	mandoc -T lint man/trd.conf.5

trd.8: lint
	mandoc -T ascii man/trd.8 >$@

trd.conf.5: lint
	mandoc -T ascii man/trd.conf.5 >$@

trd.conf.md: lint
	mandoc -T markdown man/trd.conf.5 > $@

README.md: lint
	mandoc -T markdown man/trd.8 >$@

build:
	go build -ldflags "-X main.version=${VERSION}"

realinstall: build
	install -d $(DESTDIR)$(PREFIX)/bin
	install -m 755 trd $(DESTDIR)$(PREFIX)/bin/trd

.include <bsd.prog.mk>
