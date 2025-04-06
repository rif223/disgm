package store

// TokenStore defines an interface for storing and loading tokens.
//
// The implementing types should provide the actual logic for storing the tokens,
// whether it be in-memory, in a file, or in a database.
type TokenStore interface {

	// Store saves the provided map of tokens.
	//
	// Parameters:
	//   - tokens: map[string]string – A map where keys are guild ids
	//     and values are the associated tokens (e.g., generated tokens).
	//
	// Returns:
	//   - error: An error, if any occurs during the storage process.
	//     It should return nil if the storage is successful.
	Store(tokens map[string]string) error

	// Load retrieves a map of command tokens.
	//
	// Returns:
	//   - tokens: map[string]string – A map where keys are guild ids and
	//     values are the associated tokens.
	//   - error: An error, if any occurs during the loading process.
	//     It should return nil if the loading is successful.
	Load() (tokens map[string]string, err error)
}
