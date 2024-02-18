package cli

import (
	"bufio"
	"os"
)

// Read input information from STDIN
func ReadFromSttin() string {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		return scanner.Text()

	}
	return ""
}
