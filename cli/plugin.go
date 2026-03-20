package main

// Plugin is the minimal plugin interface.
type Plugin interface {
	Name() string
	Run() error
}

// DiscoverPlugins is a stub for discovering plugins in a directory.
func DiscoverPlugins(dir string) []Plugin { _ = dir; return nil }
