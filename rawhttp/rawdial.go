package rawhttp

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/url"
	"strings"
)

// Dial connects to the named server.
// If there's no scheme specified or the scheme is "http," then default to
// port 80 and no TLS. If the scheme is "https" default to port 443 and use
// TLS. If the URL contains an explicit port number, override the scheme;
// command line options can override anything.
func Dial(serverspec string, forceTLS bool) (net.Conn, error) {
	host, port, useTLS, err := getHostPort(serverspec, forceTLS)
	if err != nil {
		return nil, err
	}
	if useTLS {
		return tls.Dial("tcp", host+":"+port,
			&tls.Config{InsecureSkipVerify: true, ServerName: host})
	} else {
		return net.Dial("tcp", host+":"+port)
	}
}

// getHostPort contains the logic for turning a user-entered server spec
// string into something that works with the GoLang Dial() function.
//
// If the servername looks like a plain name, then leave it alone other than
// ensuring there's a port number (defaulted to 80). If the servername looks
// like a URL, parse it out fully and return just the argument needed for the
// Dial function.
func getHostPort(serverspec string, forceTLS bool) (string, string, bool, error) {
	// Lazy check: there's no requirement that a URL have a slash after the
	// scheme, but all HTTP, HTTPS, and FTP urls do. So if ":/" is found,
	// assume a simple hostname with optional port number, defaulted to 80.
	if strings.Index(serverspec, ":/") == -1 {
		i := strings.Index(serverspec, ":")
		if i == len(serverspec)-1 {
			return "", "", false, fmt.Errorf("malformed port [%s]", serverspec)
		}
		if i == -1 {
			return serverspec, "80", forceTLS, nil
		} else {
			return serverspec[0:i], serverspec[i+1:], forceTLS, nil
		}
	} else {
		u, err := url.Parse(serverspec)
		if err != nil {
			return "", "", false, err
		}

		useTLS := (u.Scheme == "https" || forceTLS)

		i := strings.Index(u.Host, ":")
		if i == len(u.Host)-1 {
			return "", "", false, fmt.Errorf("malformed port: [%s]", u.Host)
		}
		if i == -1 {
			if useTLS {
				return u.Host, "443", useTLS, nil
			} else {
				return u.Host, "80", useTLS, nil
			}
		}
		return u.Host[0:i], u.Host[i+1:], useTLS, nil
	}
}
