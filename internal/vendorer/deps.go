package vendorer

type Cloner interface {
	Clone(repo, dst string) error
}

type PathFormatter interface {
	FormatCachePath(path ...string) string
	FormatOutputPath(path ...string) string
	FormatExternalPath(path ...string) string
}

type FileSystem interface {
	CreateDirectory(dir string) error
	CopyFile(src, dst string) error
}

type Finder interface {
	FindImportedProtofilesInFile(path string) ([]string, error)
	FindProtoFiles(dirs []string) []string
}

type DependencyManager interface {
	EnqueueDependencies(deps ...string)
	IsEmpty() bool
	Enqueue(dep string)
	Dequeue() string
}
