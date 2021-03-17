package generator

type (
	Universe struct {
		rootPath  string
		templates Templates
		fragments Fragments
		queries   Queries
		store     *FragmentStore
	}
)

func (u *Universe) Update(asset *Asset) {

}
