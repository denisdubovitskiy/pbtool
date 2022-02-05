package pathmanager

import "path/filepath"

type Manager struct {
	baseDir     string
	outputDir   string
	externalDir string
}

func New(cache, output, external string) *Manager {
	return &Manager{
		baseDir:     cache,
		outputDir:   output,
		externalDir: external,
	}
}

func (s *Manager) FormatCachePath(path ...string) string {
	args := append([]string{}, s.baseDir)
	return filepath.Join(append(args, path...)...)
}

func (s *Manager) FormatOutputPath(path ...string) string {
	args := append([]string{}, s.outputDir)
	return filepath.Join(append(args, path...)...)
}

func (s *Manager) FormatExternalPath(path ...string) string {
	args := append([]string{}, s.externalDir)
	return filepath.Join(append(args, path...)...)
}
