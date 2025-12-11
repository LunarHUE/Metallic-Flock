package k3s

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func StartK3sServer(ctx context.Context) error {
	const serviceName = "k3s.service"

	fmt.Printf("Starting %s...\n", serviceName)

	// We use systemctl to start the service.
	// This relies on the NixOS config having written the unit file correctly.
	cmd := exec.CommandContext(ctx, "systemctl", "start", serviceName)

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to start k3s server: %s: %w", string(output), err)
	}

	// Optional: Verify it is actually running
	if err := waitForServiceState(ctx, serviceName, "active"); err != nil {
		return fmt.Errorf("service started but did not become active: %w", err)
	}

	fmt.Println("K3s server started successfully.")
	return nil
}

func waitForServiceState(ctx context.Context, serviceName, targetState string) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			// Check is-active. It returns exit code 0 if active, non-zero otherwise.
			// We act purely on the string output for clearer "activating" states.
			cmd := exec.CommandContext(ctx, "systemctl", "is-active", serviceName)
			out, _ := cmd.Output() // Ignore error, as "inactive" returns error code 3
			state := string(out)

			// Clean up newline
			if len(state) > 0 && state[len(state)-1] == '\n' {
				state = state[:len(state)-1]
			}

			if state == targetState {
				return nil
			}

			// If it failed, stop waiting
			if state == "failed" {
				return fmt.Errorf("service entered failed state")
			}
		}
	}
}
