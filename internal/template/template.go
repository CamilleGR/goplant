package template

import (
	"embed"
	"fmt"
	"path"
)

//go:embed templates/*.tmpl
var fsys embed.FS

type Template struct {
	name     string
	fileName string
	desc     string
	code     string
	compiler CommandFunc
}

// GetCode retourne le code source du template
func (t *Template) GetCode() string {
	return t.code
}

// GetCompiler retourne la fonction compiler du template
func (t *Template) GetCompiler() CommandFunc {
	return t.compiler
}

func (t *Template) GetName() string {
	return t.name
}

func (t *Template) GetFileName() string {
	return t.fileName
}

func (t *Template) GetDesc() string {
	return t.desc
}

func newTemplate(name string, fileName string, desc string, compiler CommandFunc) Template {
	b, err := fsys.ReadFile(path.Join("templates", fileName))
	if err != nil {
		panic(fmt.Sprintf("templates: lecture de %q: %v", fileName, err))
	}
	return Template{name: name, fileName: fileName, code: string(b), desc: desc, compiler: compiler}
}

var available = []Template{
	newTemplate(
		"go",
		"go_implant.go.tmpl",
		"Simple Go executable compiled with random flag",
		BuildGoImplant,
	),
}

func GetAvailableTemplates() []Template { return available }

func ListTemplateNames() []string {
	names := make([]string, len(available))
	for i, tmpl := range available {
		names[i] = tmpl.name
	}
	return names
}

func GetTemplate(name string) (*Template, error) {
	for i, tmpl := range available {
		if tmpl.name == name {
			return &available[i], nil
		}
	}
	return nil, fmt.Errorf("template not found: %s", name)
}
