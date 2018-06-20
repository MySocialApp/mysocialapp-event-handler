package msaeventmodules

import (
	"bytes"
	wk "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"gopkg.in/gomail.v2"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"os"
	"crypto/md5"
	"encoding/hex"
)

type EmailModule struct {
	SmtpHost     string         `yaml:"smtp-host"`
	SmtpPort     int            `yaml:"smtp-port"`
	SmtpTls      bool           `yaml:"smtp-tls"`
	SmtpLogin    string         `yaml:"smtp-login"`
	SmtpPassword string         `yaml:"smtp-password"`
	Dialer       *gomail.Dialer `yaml:"-"`
}

// possible concurrent write, todo improve
func (e *EmailModule) GetDialer() *gomail.Dialer {
	if e.Dialer != nil {
		return e.Dialer
	}
	e.Dialer = gomail.NewDialer(e.SmtpHost, e.SmtpPort, e.SmtpLogin, e.SmtpPassword)
	e.Dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return e.Dialer
}

func (e *EmailModule) Prepare() *Email {
	return &Email{
		emailModule: e,
	}
}

type Email struct {
	emailModule *EmailModule
	from        string
	to          []string
	cc          []string
	ccc         []string
	subject     string
	body        string
	files       []EmailFile
	htmlBody    bool
}

func (e *Email) From(email string) *Email {
	e.from = email
	return e
}

func (e *Email) Subject(subject string) *Email {
	e.subject = subject
	return e
}

func (e *Email) Body(body string, isHtml bool) *Email {
	e.body = body
	e.htmlBody = isHtml
	return e

}

func (e *Email) AddTo(to string) *Email {
	if e.to == nil {
		e.to = []string{to}
		return e
	}
	e.to = append(e.to, to)
	return e
}

func (e *Email) AddCc(cc string) *Email {
	if e.cc == nil {
		e.cc = []string{cc}
		return e
	}
	e.cc = append(e.cc, cc)
	return e
}

func (e *Email) AddCcc(ccc string) *Email {
	if e.ccc == nil {
		e.ccc = []string{ccc}
		return e
	}
	e.ccc = append(e.ccc, ccc)
	return e
}

func (e *Email) AddFile(buff *bytes.Buffer, filename string) *Email {
	if e.files == nil {
		e.files = []EmailFile{{Content:buff, Filename: filename}}
		return e
	}
	e.files = append(e.files, EmailFile{Content:buff, Filename: filename})
	return e
}

func (e *Email) AddHtmlToPdfFile(buff *bytes.Buffer, filename string) *Email {
	pdf, err := Html2pdf(buff)
	if err != nil {
		log.Printf("fail to generate pdf: %s", err.Error())
		return e
	}
	e.AddFile(pdf, filename)
	return e
}

func (e *Email) Send() error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", e.from)
	msg.SetHeader("Subject", e.subject)
	if e.htmlBody {
		msg.SetBody("text/html", e.body)
	}
	if !e.htmlBody {
		msg.SetBody("text", e.body)
	}
	if e.to != nil {
		r := []string{}
		for _, t := range e.to {
			r = append(r, t)
		}
		msg.SetHeader("To", r...)
	}
	if e.cc != nil {
		r := []string{}
		for _, c := range e.cc {
			r = append(r, c)
		}
		msg.SetHeader("CC", r...)
	}
	if e.ccc != nil {
		r := []string{}
		for _, c := range e.ccc {
			r = append(r, c)
		}
		msg.SetHeader("CCC", r...)
	}
	filename := ""
	if e.files != nil && len(e.files) > 0 {
		for _, file := range e.files {
			// TODO improve using real filename
			filename = fmt.Sprintf("/tmp/%s_%s", file.Hash(), file.Filename)
			if err := ioutil.WriteFile(filename, file.Content.Bytes(), 0755); err != nil {
				return err
			}
			msg.Attach(filename)
		}
	}
	defer os.Remove(filename)
	return e.emailModule.GetDialer().DialAndSend(msg)
}

type EmailFile struct {
	Content *bytes.Buffer
	Filename string
}

func (e *EmailFile) Hash() string {
	hasher := md5.New()
	hasher.Write(e.Content.Bytes())
	hasher.Write([]byte(e.Filename))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Html2pdf(buff *bytes.Buffer) (*bytes.Buffer, error) {
	pdfg, err := wk.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wk.OrientationPortrait)
	pdfg.Grayscale.Set(true)

	// Create a new input page from an URL
	page := wk.NewPageReader(buff)

	// Set options for this page
	//page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(1)

	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	if err = pdfg.Create(); err != nil {
		return nil, err
	}
	// return buffer content
	return pdfg.Buffer(), nil
}
