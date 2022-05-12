package envloader

import (
	"os"
	"strings"
)

// Load environment variables from the `.env` file and set them to the `os.Environ` variable.
// Returns a list of environment variables that were not found in the `.env` file.
// Returns an error if the `.env` file cannot be read.
//	Example:
// 		keys, err := envloader.LoadEnv()
func LoadEnv() ([]string, error) {
	// Read the file
	file, err := os.ReadFile(".env")
	if err != nil {
		return []string{}, err
	}

	// Split into lines
	lines := strings.Split(string(file), "\n")

	// Loop over lines
	for _, line := range lines {
		// Split line into key and value
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue
		}
		os.Setenv(parts[0], parts[1])
	}

	return lines, err
}
