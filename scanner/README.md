# Scanner

This package provides a pluggable scanner interface for the project.

The scanner should expose a Scan(path, opts) -> []Finding signature, where a Finding contains a file, message, and optional metadata. The initial implementation included here is a no-op placeholder useful for integration and testing.

Implementers:
- Provide concrete scanner implementations under scanner/** and wire them into the backend services.
- Ensure results are serializable to JSON if stored in the results table.
