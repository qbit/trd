package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"

	"golang.org/x/sys/unix"
)

var debug bool
var version string

// Req is a representation of our tftp requests
type Req struct {
	ip net.IP
	op string
	fn string
}

func (r *Req) parse(s string) error {
	parts := strings.Split(s, " ")
	if len(parts) != 3 {
		err := errors.New("Invalid string")
		return err
	}

	ip := net.ParseIP(parts[0])
	if ip == nil {
		err := errors.New("Invalid IP")
		return err
	}

	r.ip = ip
	r.op = parts[1]
	r.fn = parts[2]

	return nil
}

// Rewrite represent one of our re-write lines.
type Rewrite struct {
	ip    net.IP
	match string
	fn    string
}

// Rewrites are a grouping of lines
type Rewrites []Rewrite

func (r *Rewrite) parse(s string) error {
	parts := strings.Split(s, " ")
	if len(parts) != 3 {
		err := fmt.Errorf("Invalid string: '%v'", parts)
		return err
	}

	ip := net.ParseIP(parts[0])
	if ip == nil {
		err := errors.New("Invalid IP")
		return err
	}

	r.ip = ip
	r.match = parts[1]
	r.fn = parts[2]

	return nil
}

func makeRewrites(path string) (*Rewrites, error) {
	var rw Rewrites

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var r Rewrite
		err := r.parse(string(scanner.Text()))
		if err != nil {
			log.Fatal(err)
		}
		rw = append(rw, r)
	}

	return &rw, nil
}

func dbg(s string) {
	if debug {
		log.Printf(s)
	}
}

func main() {
	var cfile = flag.String("c", "/etc/trd.conf", "path to rewrites file.")
	var sock = flag.String("s", "/tmp/trd.sock", "path to socket.")
	var root = flag.String("r", "/var/tftpd", "path to chroot to.")
	var dUser = flag.String("u", "_tftpd", "user to drop privs to")
	var ver = flag.Bool("v", false, "print version and exit")
	flag.BoolVar(&debug, "debug", false, "causes decomer to print debug info")
	flag.Parse()

	if *ver {
		fmt.Println(version)
		os.Exit(0)
	}

	os.Remove(*sock)

	ln, err := net.Listen("unix", *sock)
	if err != nil {
		log.Fatal(err)
	}
	dbg(fmt.Sprintf("listening to: %s", *sock))

	defer ln.Close()

	rws, err := makeRewrites(*cfile)
	if err != nil {
		log.Fatal(err)
	}

	u, err := user.Lookup(*dUser)
	if err != nil {
		log.Fatal(err)
	}

	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		log.Fatal(err)
	}

	gid, err := strconv.Atoi(u.Gid)
	if err != nil {
		log.Fatal(err)
	}

	err = syscall.Chroot(*root)
	if err != nil {
		log.Fatal(err)
	}
	dbg(fmt.Sprintf("chrooting to: %s", *root))

	err = syscall.Setgid(gid)
	if err != nil {
		log.Fatal(err)
	}
	dbg(fmt.Sprintf("setgid to: %s", u.Gid))

	err = syscall.Setuid(uid)
	if err != nil {
		log.Fatal(err)
	}
	dbg(fmt.Sprintf("setuid to: %s", u.Uid))

	unix.PledgePromises("stdio rpath unix")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			var r Req
			line := string(scanner.Text())
			err := r.parse(line)
			if err != nil {
				log.Println("Spitting back")
				fmt.Fprintf(conn, "%s\n", line)
				break
			}

			actionString := r.fn
			for _, rw := range *rws {
				if r.ip.Equal(rw.ip) {
					if rw.match == r.fn {
						log.Printf("rewriting '%s' as '%s' for %s\n", r.fn, rw.fn, rw.ip)
						actionString = rw.fn
					}
				}
			}
			fmt.Fprintf(conn, "%s\n", actionString)
		}
	}
}
