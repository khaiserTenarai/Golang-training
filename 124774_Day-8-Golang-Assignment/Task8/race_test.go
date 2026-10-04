package race_detection

import "testing"

func TestCounter(t *testing.T) {
	got := Counter()

	t.Logf("Counter value: %d", got)
}
