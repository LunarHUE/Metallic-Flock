package k3s

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// CreateJoinToken generates a new cluster join token by calling the local k3s instance.
//
// Parameters:
//   - description: A human-readable note for auditing (e.g., "worker-node-01").
//   - ttl: The Time-To-Live for the token. If 0, it defaults to K3s default (usually 24h).
//
// Returns:
//   - The generated token string (trimmed).
//   - An error if the k3s command fails or is not found.
func CreateJoinToken(description string, ttl time.Duration) (string, error) {
	binPath, err := exec.LookPath("k3s")
	if err != nil {
		return "", fmt.Errorf("k3s binary not found in PATH: %w", err)
	}

	// Command: k3s token create --description "..." --ttl "..."
	args := []string{"token", "create"}

	if description != "" {
		args = append(args, "--description", description)
	}

	if ttl > 0 {
		// Format duration properly (e.g., "1h0m0s")
		args = append(args, "--ttl", ttl.String())
	}

	// 3. Prepare the command execution
	cmd := exec.Command(binPath, args...)

	// Capture Standard Output and Standard Error for debugging
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// 4. Run the command
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to create token (stderr: %s): %w", strings.TrimSpace(stderr.String()), err)
	}

	// 5. Clean up the output
	// The CLI usually returns the token with a trailing newline
	token := strings.TrimSpace(stdout.String())
	if token == "" {
		return "", fmt.Errorf("k3s command succeeded but returned empty token")
	}

	return token, nil
}
