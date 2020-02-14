package evalidate_test

import (
	"testing"

	"github.com/kumarabd/evalidate"
)

var (
	samples = []struct {
		mail    string
		format  bool
		account bool //host+user
	}{
		{mail: "abishekkumar92@gmail.com", format: true, account: true},
		{mail: "support@g2mail.com", format: true, account: false},
	}
)

func TestValidateHost(t *testing.T) {
	for _, s := range samples {
		if !s.format {
			continue
		}

		err := evalidate.ValidateHost(s.mail)
		if err != nil && s.account == true {
			t.Errorf(`"%s" => host error: "%v"`, s.mail, err)
		}
		if err == nil && s.account == false {
			t.Errorf(`"%s" => account error`, s.mail)
		}
	}
}

func TestValidateFormat(t *testing.T) {
	for _, s := range samples {
		err := evalidate.ValidateFormat(s.mail)
		if err != nil && s.format == true {
			t.Errorf(`"%s" => validate error: "%v"`, s.mail, err)
		}
		if err == nil && s.format == false {
			t.Errorf(`"%s" => format error`, s.mail)
		}
	}
}
