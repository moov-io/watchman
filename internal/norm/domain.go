package norm

import (
	"net"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// DomainKeys returns every DNS name from the prepared host down to eTLD+1.
// Public suffixes (co.uk, ru), IP addresses, and unparseable hosts are omitted
// so a query of "ru" does not match every .ru listing.
func DomainKeys(input string) []string {
	host := Website(input)
	if host == "" {
		return nil
	}
	if i := strings.LastIndex(host, ":"); i >= 0 && !strings.Contains(host, "]") {
		host = host[:i]
	}
	host = strings.Trim(host, ".")
	if host == "" {
		return nil
	}
	if ip := net.ParseIP(strings.Trim(host, "[]")); ip != nil {
		return nil
	}

	etld1, err := publicsuffix.EffectiveTLDPlusOne(host)
	if err != nil || etld1 == "" || !strings.HasSuffix(host, etld1) {
		return nil
	}

	var out []string
	for {
		out = append(out, host)
		if host == etld1 {
			break
		}
		dot := strings.IndexByte(host, '.')
		if dot < 0 {
			break
		}
		host = host[dot+1:]
	}
	return out
}

// DomainKeysFromEmail walks the host of an email address. Consumer mail
// providers (gmail, yahoo, mail.ru, …) are skipped so those hosts are not
// treated as an entity domain.
func DomainKeysFromEmail(email string) []string {
	email = strings.ToLower(strings.TrimSpace(email))
	at := strings.LastIndex(email, "@")
	if at < 0 || at == len(email)-1 {
		return nil
	}
	keys := DomainKeys(email[at+1:])
	if len(keys) == 0 {
		return nil
	}
	if _, skip := publicEmailHosts[keys[len(keys)-1]]; skip {
		return nil
	}
	return keys
}

// publicEmailHosts is keyed by eTLD+1. It is only applied when extracting a
// domain from an email address, not from website= or domain= values.
var publicEmailHosts = map[string]struct{}{
	"gmail.com": {}, "googlemail.com": {}, "gmail.co.uk": {},
	"yahoo.com": {}, "yahoo.co.uk": {}, "yahoo.co.jp": {}, "yahoo.fr": {},
	"yahoo.de": {}, "yahoo.it": {}, "yahoo.es": {}, "yahoo.com.br": {},
	"yahoo.com.mx": {}, "yahoo.co.in": {}, "ymail.com": {}, "rocketmail.com": {},
	"hotmail.com": {}, "hotmail.co.uk": {}, "hotmail.fr": {}, "hotmail.it": {},
	"outlook.com": {}, "outlook.co.uk": {}, "live.com": {}, "msn.com": {},
	"icloud.com": {}, "me.com": {}, "mac.com": {},
	"aol.com": {},
	"mail.ru": {}, "inbox.ru": {}, "list.ru": {}, "bk.ru": {}, "rambler.ru": {},
	"yandex.ru": {}, "yandex.com": {},
	"qq.com": {}, "163.com": {}, "126.com": {}, "sina.com": {}, "sohu.com": {},
	"proton.me": {}, "protonmail.com": {}, "pm.me": {},
	"gmx.com": {}, "gmx.de": {}, "gmx.net": {}, "web.de": {}, "t-online.de": {},
	"zoho.com": {}, "mail.com": {}, "inbox.com": {}, "fastmail.com": {},
	"hey.com": {}, "tutanota.com": {}, "tuta.com": {},
	"naver.com": {}, "daum.net": {}, "rediffmail.com": {},
	"libero.it": {}, "virgilio.it": {},
	"orange.fr": {}, "wanadoo.fr": {}, "free.fr": {}, "laposte.net": {},
	"seznam.cz": {}, "wp.pl": {}, "o2.pl": {}, "ukr.net": {},
	"comcast.net": {}, "att.net": {}, "sbcglobal.net": {}, "verizon.net": {},
	"bellsouth.net": {}, "charter.net": {},
	"btinternet.com": {}, "sky.com": {}, "virginmedia.com": {},
}
