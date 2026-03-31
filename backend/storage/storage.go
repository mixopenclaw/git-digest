package storage

// Storage defines an interface for artifact storage.
type Storage interface {
	SaveArtifact(path string, data []byte) error
	LoadArtifact(path string) ([]byte, error)
}
