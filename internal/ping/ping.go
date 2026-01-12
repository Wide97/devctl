package ping

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Check esegue una richiesta GET a baseURL + "/ping".
// Restituisce nil se il servizio risponde 200 OK, altrimenti errore.
func Check(ctx context.Context, baseURL string) error {
	client := &http.Client{
		Timeout: 3 * time.Second, // timeout breve per ping rapido
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+"/ping", nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("service returned status %d", resp.StatusCode)
	}

	return nil
}
