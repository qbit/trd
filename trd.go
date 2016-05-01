package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var debug bool

type Req struct {
	ip net.IP
	op string
	fn string
}

type Rewrites []Req

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

func makeRewrites(path string) (*Rewrites, error) {
	var rw Rewrites

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var r Req
		r.parse(string(scanner.Text()))
		rw = append(rw, r)
	}

	return &rw, nil
}

func main() {
	var cfile = flag.String("c", "/etc/rewrites.conf", "Path to rewrites file")
	var sock = flag.String("s", "/tmp/trd.sock", "Path to socket")
	flag.BoolVar(&debug, "debug", false, "causes decomer to print debug info")
	flag.Parse()

	os.Remove(*sock)

	ln, err := net.Listen("unix", *sock)
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	rws, err := makeRewrites(*cfile)
	if err != nil {
		log.Fatal(err)
	}

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
				io.Copy(conn, conn)
			}

			for _, rw := range *rws {
				if rw.ip.Equal(r.ip) {
					r.fn = rw.fn
				}
			}

			fmt.Fprintf(conn, "%s %s %s\n", r.ip, r.op, r.fn)
		}
	}
}
