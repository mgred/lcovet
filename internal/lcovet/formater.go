package lcovet

import (
	"io"
	"text/template"

	"github.com/mgred/lcovet/internal/output"
)

func NewFormatter(r *[]Record) *Formatter {
	temp := template.New("output")
	return &Formatter{records: r, template: temp}
}

type Formatter struct {
	records  *[]Record
	template *template.Template
}

func (f *Formatter) Simple(o io.Writer) error {
	return f.execute(o, output.PrintTemplate)
}

func (f *Formatter) Html(o io.Writer) error {
	return f.execute(o, output.PrintTemplate)
}

func (f *Formatter) execute(w io.Writer, t string) error {
	templ, err := f.template.Parse(t)
	if err != nil {
		return err
	}
	if err = templ.Execute(w, *f.records); err != nil {
		return err
	}
	return nil
}
