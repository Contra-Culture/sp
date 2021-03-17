package generator

import "github.com/Contra-Culture/sp/generator/asset"

type (
	Universe struct {
		rootPath  string
		templates Templates
		fragments Fragments
		queries   Queries
		store     *FragmentStore
	}
)

func (u *Universe) Update(asset *asset.Asset) {

}
