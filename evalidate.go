package evalidate

import (
	"net"
	"net/smtp"
	"regexp"
	"strings"
	"time"
)

const forceDisconnectAfter = time.Second * 5

var (
	emailRegexp = regexp.MustCompile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
)

func ValidateFormat(email string) error {
	if !emailRegexp.MatchString(email) {
		return ErrBadFormat
	}
	return nil
}

func ValidateHost(email string) error {
	_, host := split(email)
	_, err := net.LookupMX(host)
	if err != nil {
		return ErrUnresolvableHost(err)
	}

	// client, err := DialTimeout(fmt.Sprintf("%s:%d", mx[0].Host, 25), forceDisconnectAfter)
	// if err != nil {
	// 	return ErrTimeout(err)
	// }
	// defer client.Close()

	// err = client.Hello("gmail.com")
	// if err != nil {
	// 	return ErrNotReachable(err)
	// }
	// err = client.Mail("abishekkumar92@gmail.com")
	// if err != nil {
	// 	return ErrClientEmail(err)
	// }
	// err = client.Rcpt(email)
	// if err != nil {
	// 	return ErrClientRct(err)
	// }
	return nil
}

func DialTimeout(addr string, timeout time.Duration) (*smtp.Client, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}

	t := time.AfterFunc(timeout, func() { conn.Close() })
	defer t.Stop()

	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func split(email string) (account, host string) {
	i := strings.LastIndexByte(email, '@')
	account = email[:i]
	host = email[i+1:]
	return
}
