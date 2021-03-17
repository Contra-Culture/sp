package generator

type (
	FragmentStore struct {
		rootPath string
	}
)

func NewFragmentStore(rootPath string) (store *FragmentStore) {
	store.rootPath = rootPath
	return
}

func (s *FragmentStore) Save(path string, content string) (err error) {
	return
}
