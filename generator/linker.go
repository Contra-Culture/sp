package generator

type (
	loadOption bool
	Link       struct {
		Name           string
		Title          string
		Description    string
		Reference      string
		Representation string
	}
	NS struct {
		HTMLClasses []string
		Link        Link
		Children    Links
	}
	Links  map[string]*Link
	Linker struct {
		rootDir string
		links   Links
	}
)

const (
	Load     = loadOption(true)
	DontLoad = loadOption(false)
)

func NewLinker(root string, doLoad loadOption) (linker *Linker, err error) {
	linker.rootDir = root
	if doLoad {
		err = linker.load(root)
		return
	}
	linker.links = make(Links)
	return
}
func (l *Linker) Get(path []string) (link Link) {
	return
}
func (l *Linker) Make(path []string, name string) (err error) {
	return
}
func (l *Linker) Save() (err error) {
	return
}
func (l *Linker) load(filePath string) (err error) {
	return
}
