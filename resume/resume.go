package resume

import (
	"github.com/gobuffalo/packr"
	yaml "gopkg.in/yaml.v1"
)

func Resume() *ResumeInfo {
	resumeData, err := packr.NewBox(".").Find("resume.yaml")
	if err != nil {
		panic(err)
	}

	var r ResumeInfo
	err = yaml.Unmarshal(resumeData, &r)
	if err != nil {
		panic(err)
	}

	return &r
}

type ResumeInfo struct {
	Name      string     `yaml:"name"`
	Info      Info       `yaml:"info"`
	Contact   Contact    `yaml:"contact"`
	Location  Location   `yaml:"location"`
	Social    []Social   `yaml:"social"`
	Skills    []Skill    `yaml:"skills"`
	Work      []Work     `yaml:"work"`
	Interests []Interest `yaml:"interests"`
}

type Info struct {
	Position string `yaml:"position"`
	Summary  string `yaml:"summary"`
}

type Contact struct {
	Email   string `yaml:"email"`
	Website string `yaml:"website"`
}

type Social struct {
	Network  string `yaml:"network"`
	Username string `yaml:"username"`
	URL      string `yaml:"url"`
}

type Skill struct {
	Name     string   `yaml:"name"`
	Level    string   `yaml:"level"`
	Keywords []string `yaml:"keywords"`
}

type Work struct {
	Company    Institution `yaml:"company"`
	Position   string      `yaml:"position"`
	Date       Date        `yaml:"date"`
	Keywords   []string    `yaml:"keywords"`
	Summary    string      `yaml:"summary"`
	Highlights []string    `yaml:"highlights"`
}

type Education struct {
	Institution Institution `yaml:"institution"`
	Area        string      `yaml:"area"`
	Type        string      `yaml:"type"`
	Date        Date        `yaml:"date"`
	Summary     string      `yaml:"summary"`
	Highlights  []string    `yaml:"highlights"`
}

type Institution struct {
	Name     string   `yaml:"name"`
	Website  string   `yaml:"website"`
	Location Location `yaml:"location"`
}

type Interest struct {
	Name     string   `yaml:"name"`
	Date     Date     `yaml:"date"`
	Keywords []string `yaml:"keywords"`
	Summary  string   `yaml:"summary"`
}

type Date struct {
	Start string `yaml:"start"`
	End   string `yaml:"end"`
}

type Location struct {
	City    string `yaml:"city"`
	Country string `yaml:"country"`
}
