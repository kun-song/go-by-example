package main

import (
	"fmt"
	"net"
	"net/url"
)

/**
 * `net/url`：URL 解析
 */
func main() {
	p := fmt.Println

	// schema, authentication info, host, port, path, query params, query fragment
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	p(u) // postgres://user:pass@host.com:5432/path?k=v#f

	p(u.Scheme) // postgres

	// url.User = username + password
	p(u.User)            // user:pass
	p(u.User.Username()) // user
	pa, _ := u.User.Password()
	p(pa) // pass

	// Host = hostname + port，使用 net.SplitHostPort 分开两者
	p(u.Host) // host.com:5432
	hostname, port, _ := net.SplitHostPort(u.Host)
	p(hostname, port) // host.com 5432

	p(u.Path)     // /path
	p(u.RawQuery) // k=v
	p(u.Fragment) // f

	m, _ := url.ParseQuery(u.RawQuery)
	p(m)         // map[k:[v]]
	p(m["k"])    // [v]
	p(m["k"][0]) // v
}
