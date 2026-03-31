package scanner

// Scanner defines the minimal scanner interface used by the backend.
// Implementations should return findings for a given path and options.

type Finding struct {
	File    string                 `json:"file"`
	Message string                 `json:"message"`
	Meta    map[string]interface{} `json:"meta,omitempty"`
}

// Scanner is the interface implemented by concrete scanners.
type Scanner interface {
	Scan(path string, opts map[string]string) ([]Finding, error)
}

// NewScanner returns a simple no-op scanner (placeholder).
func NewScanner() Scanner {
	return &noopScanner{}
}

type noopScanner struct{}

func (n *noopScanner) Scan(path string, opts map[string]string) ([]Finding, error) {
	// TODO: real scanning logic
	return []Finding{}, nil
}
