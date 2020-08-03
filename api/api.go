package api

import (
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

const host = "https://www.stine.uni-hamburg.de"
const apiURL = host + "/scripts/mgrqispi.dll"
const origin = host
const referer = origin + "/"
const appName = "CampusNet"

const defaultTimeout = 15 * time.Second

// SkipNone means: do not skip (default)
const SkipNone = ""

// SkipNext means: skip to next day/week
const SkipNext = "N"

// SkipPrev means: skip to previous day/week
const SkipPrev = "P"

// ScheduleDay means: show day view
const ScheduleDay = "0"

// ScheduleWeek means: show week view
const ScheduleWeek = "1"

var reRefresh = regexp.MustCompile(`ARGUMENTS=-N(\d+),-N(\d+)`)
var reFiletransfer = regexp.MustCompile(`href=["'](\/scripts\/filetransfer\.exe\?\S+)["']`)
var reSpace = regexp.MustCompile(`\s+`)
var reNewLine = regexp.MustCompile(`\s*[\r\n]\s*`)
var reBr = regexp.MustCompile(`<br\S*/?>`)

var cnscCookieURL = url.URL{
	Scheme: "https",
	Host:   "www.stine.uni-hamburg.de",
	Path:   "/scripts",
}

// Account represents a STiNE account.
type Account struct {
	client  *http.Client
	session string
}

// Schedule represents a schedule for a date.
// Date is formatted by STiNE, e.g. as "Mo, 13. Jul. 2020".
// Entries consist of: Course ID, Name, Teachers, Time (Start/End), Room.
type Schedule struct {
	Date    string
	Entries [][]string
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

// Session returns the current session ID and cookie.
func (acc *Account) Session() (string, string) {
	return acc.session, acc.client.Jar.Cookies(&cnscCookieURL)[0].Value
}

// SetSession allows you to reuse a session ID and cookie.
func (acc *Account) SetSession(id, cnsc string) {
	acc.session = id
	acc.client.Jar.SetCookies(&cnscCookieURL, []*http.Cookie{
		{
			Name:  "cnsc",
			Value: cnsc,
		},
	})
}

// SessionValid checks if the current session is valid.
func (acc *Account) SessionValid() error {
	res, err := acc.DoFormRequest(url.Values{
		"APPNAME":   {appName},
		"PRGNAME":   {"EXTERNALPAGES"},
		"ARGUMENTS": {"-N" + acc.session},
	})
	if err != nil {
		return err
	}

	defer res.Body.Close()

	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if strings.Contains(string(html), "<h1>Timeout!</h1>") {
		return errors.New("timeout")
	}
	if strings.Contains(string(html), "<h1>Zugang verweigert</h1>") {
		return errors.New("invalid")
	}

	return nil
}

// Scheduler returns an array of schedules.
// If given, date must be formatted as: DD.MM.YYYY; default is current date.
// Skip can be empty, "N" (next) or "P" (previous).
// View can be "0" (day) or "1" (week).
func (acc *Account) Scheduler(date, skip, view string) ([]Schedule, error) {
	var schedules []Schedule

	res, err := acc.DoFormRequest(url.Values{
		"APPNAME":   {appName},
		"PRGNAME":   {"SCHEDULERPRINT"},
		"ARGUMENTS": {"sessionno,menuid,date,skip,view"},
		"sessionno": {acc.session},
		"menuid":    {"000000"}, // orig: 000267
		"date":      {date},
		"skip":      {skip},
		"view":      {view},
	})
	if err != nil {
		return schedules, err
	}

	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return schedules, err
	}

	scrapeTableBody(document, func(i int, tr *goquery.Selection) {
		tr.Find("td").Each(func(j int, td *goquery.Selection) {
			if td.HasClass("tbhead") {
				if colspan, ok := td.Attr("colspan"); ok && colspan == "100%" {
					schedules = append(schedules, Schedule{
						Date: scrapeText(td),
					})
				}
				return
			} else if !td.HasClass("tbdata") { // tbsubhead, etc. --> don't care
				// skip
				return
			}

			s := len(schedules)
			if s == 0 {
				// missing schedule header --> scheduler is probably empty
				return
			}
			schedule := &schedules[s-1]

			if j == 0 {
				schedule.Entries = append(schedule.Entries, []string{strings.SplitN(scrapeText(td), "\n", 2)[0]})
			} else if j < 5 {
				e := len(schedule.Entries)
				if e == 0 {
					// invalid layout
					return
				}
				entries := &schedule.Entries[e-1]
				*entries = append(*entries, strings.SplitN(scrapeText(td), "\n", 2)[0])
			}
		})
	})

	return schedules, nil
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

// Exams returns an array of exams, each with: ID, Name, Type, Date.
// semesterID can be empty (for the current semester), or and ID such as "099999904632582" (SoSe20), "999" (all)
func (acc *Account) Exams(semesterID string) ([][]string, error) {
	var exams [][]string

	res, err := acc.DoFormRequest(url.Values{
		"APPNAME":   {appName},
		"PRGNAME":   {"MYEXAMS"},
		"ARGUMENTS": {"sessionno,menuid,semester"},
		"sessionno": {acc.session},
		"menuid":    {"000000"}, // orig: 000316
		"semester":  {semesterID},
	})
	if err != nil {
		return exams, err
	}

	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return exams, err
	}

	scrapeTableBody(document, func(i int, tr *goquery.Selection) {
		var exam []string
		tr.Find("td").Each(func(j int, td *goquery.Selection) {
			if j < 4 {
				exam = append(exam, strings.SplitN(scrapeText(td), "\n", 2)[0])
			}
		})
		exams = append(exams, exam)
	})

	return exams, nil
}

// ExamResults returns an array of exam results, each with: ID, Name, Date, Grade, Grade text.
// semesterID can be empty (for the current semester), or and ID such as "099999904632582" (SoSe20), "999" (all)
func (acc *Account) ExamResults(semesterID string) ([][]string, error) {
	var exams [][]string

	res, err := acc.DoFormRequest(url.Values{
		"APPNAME":   {appName},
		"PRGNAME":   {"EXAMRESULTS"},
		"ARGUMENTS": {"sessionno,menuid,semester"},
		"sessionno": {acc.session},
		"menuid":    {"000000"}, // orig: 000316
		"semester":  {semesterID},
	})
	if err != nil {
		return exams, err
	}

	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return exams, err
	}

	scrapeTableBody(document, func(i int, tr *goquery.Selection) {
		var exam []string
		tr.Find("td").Each(func(j int, td *goquery.Selection) {
			if j == 0 {
				text := strings.SplitN(scrapeText(td), "\n", 2)[0]
				IDandName := strings.SplitN(text, "\u00a0 \u00a0", 2)
				if len(IDandName) == 1 {
					exam = append(exam, "")
				}
				exam = append(exam, IDandName...)
			} else if j < 4 {
				exam = append(exam, strings.SplitN(scrapeText(td), "\n", 2)[0])
			}
		})
		exams = append(exams, exam)
	})

	return exams, nil
}
