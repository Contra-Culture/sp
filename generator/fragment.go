package generator

import "io"

type (
	TemplateName       string
	TemplateArgName    string
	TemplateHelperName string
	TemplateArgs       map[TemplateArgName]interface{}
	TemplateHelper     func(TemplateArgs, io.Writer)
	TemplateHelpers    map[TemplateHelperName]TemplateHelper
	TemplateFn         func(args TemplateArgs, helpers TemplateHelpers, writer io.Writer)
	Template           struct {
		name TemplateName
		fn   TemplateFn
	}
	Templates    map[TemplateName]*Template
	FragmentName string
	Fragments    map[FragmentName]Fragment
	Fragment     struct {
		isPage    bool
		name      FragmentName
		template  *Template
		queries   Queries
		fragments Fragments
	}
	PreparationErrorHandler func(t *Template, err error)
)

var (
	NoQueries   Queries
	NoFragments Fragments
)

func NewPage(name string, template *Template, queries Queries, fragments Fragments) (page *Fragment) {
	page.isPage = true
	page.name = FragmentName(name)
	page.template = template
	page.queries = queries
	page.fragments = fragments
	return
}

func NewFragment(name string, template *Template, queries Queries, fragments Fragments) (fragment *Fragment) {
	fragment.isPage = false
	fragment.name = FragmentName(name)
	fragment.template = template
	fragment.queries = queries
	fragment.fragments = fragments
	return
}

func (f *Fragment) Prepare(dirPath string, args TemplateArgs, helpers TemplateHelpers, store *FragmentStore, preparationErrorHandler PreparationErrorHandler) {
	_, err := f.template.Render(args, helpers)

	if err != nil {
		preparationErrorHandler(f.template, err)
		return
	}
	if f.isPage {

	} else {

	}
}

func NewTemplate(name string, fn TemplateFn) (template *Template) {
	template.name = TemplateName(name)
	template.fn = fn
	return
}

func (t *Template) Render(args TemplateArgs, helpers TemplateHelpers) (result string, err error) {
	return
}
