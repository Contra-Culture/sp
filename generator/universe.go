package generator

import "github.com/Contra-Culture/sp/generator/asset"

type (
	Universe struct {
		rootPath        string
		templates       Templates
		fragments       Fragments
		queries         Queries
		assetAttributes asset.AssetAttributeTypes
		assetTypes      []asset.AssetType
		store           *FragmentStore
	}
)

func New() (univ *Universe, err error) {
	return
}

func (u *Universe) Update(asset *asset.Asset) {

}

func (u *Universe) AddTemplate() {

}
func (u *Universe) AddFragment() {

}
func (u *Universe) AddQuery() {

}
func (u *Universe) AddAssetType() {

}
