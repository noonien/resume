package template

import (
	"html/template"
	"io/ioutil"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/gobuffalo/packr"
	"github.com/russross/blackfriday"
)

func HTML() (*template.Template, error) {
	var t *template.Template
	err := packr.NewBox("html").Walk(func(name string, f packr.File) error {
		if !strings.HasSuffix(name, ".tpl.html") {
			return nil
		}
		name = strings.TrimSuffix(name, ".tpl.html")

		if t == nil {
			t = template.New(name).Funcs(sprig.FuncMap()).Funcs(template.FuncMap{
				"md": func(in string) template.HTML {
					return template.HTML(
						blackfriday.Run([]byte(in)),
					)
				},
			})
		} else {
			t = t.New(name)
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		s := string(b)

		_, err = t.Parse(s)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	t = t.Lookup("resume")
	return t, nil
}
