package api

import (
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"time"

	"github.com/pkg/errors"
)

const host = "https://www.stine.uni-hamburg.de"
const apiURL = host + "/scripts/mgrqispi.dll"
const origin = host
const referer = origin + "/"
const appName = "CampusNet"

const defaultTimeout = 15 * time.Second

var reRefresh = regexp.MustCompile(`ARGUMENTS=-N(\d+),-N(\d+)`)
var reFiletransfer = regexp.MustCompile(`href=["'](\/scripts\/filetransfer\.exe\?\S+)["']`)

// Account represents a STiNE account.
type Account struct {
	client  *http.Client
	session string
}

// NewAccount creates a new Account.
func NewAccount() Account {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{
		Timeout: defaultTimeout,
		Jar:     jar,
	}

	return Account{
		client: client,
	}
}

// Login starts a new session.
func (acc *Account) Login(user, pass string) error {
	res, err := acc.DoFormRequest(url.Values{
		"usrname":   {user},
		"pass":      {pass},
		"APPNAME":   {appName},
		"PRGNAME":   {"LOGINCHECK"},
		"ARGUMENTS": {"clino,usrname,pass,menuno,menu_type,browser,platform"},
		"clino":     {"000000000000001"},
		"menuno":    {"000000"}, // orig: 000265
		"menu_type": {"classic"},
		"browser":   {""},
		"platform":  {""},
	})
	if err != nil {
		return err
	}

	refresh := res.Header.Get("refresh")
	match := reRefresh.FindStringSubmatch(refresh)
	if len(match) < 3 {
		return errors.New("invalid refresh")
	}
	acc.session = match[1]

	return nil
}

// SchedulerExport exports the schedule of a given month or week as an .ics file.
// Date examples: "Y2020M06" (month), "Y2020W25" (week).
// Dates must be in the present or future.
func (acc *Account) SchedulerExport(date string) (string, error) {
	res, err := acc.DoFormRequest(url.Values{
		"APPNAME":   {appName},
		"PRGNAME":   {"SCHEDULER_EXPORT_START"},
		"ARGUMENTS": {"sessionno,menuid,date"},
		"sessionno": {acc.session},
		"menuid":    {"000000"}, // orig: 000365
		"date":      {date},
	})
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	match := reFiletransfer.FindStringSubmatch(string(html))
	if len(match) < 2 {
		return "", errors.New("no url found")
	}

	return host + match[1], nil
}
