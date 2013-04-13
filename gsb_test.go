package gsb_test

import (
	"testing"

	"github.com/balasanjay/gsb"
)

func TestListNetwork(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping network test in short mode")
	}

	t.Parallel()

	l, err := gsb.DefaultClient.List(nil)
	if err != nil {
		t.Errorf("Expecting nil error, got %v", err)
		t.FailNow()
	}

	if len(l) < 2 {
		t.Errorf("Expecting at least 2 lists, got %v lists", len(l))
		t.FailNow()
	}

	t.Logf("Got lists %v", l)
}

func TestListParse(t *testing.T) {
	t.Parallel()

	// TODO(sanjay): write this test
}

func TestListURL(t *testing.T) {
	t.Parallel()

	// TODO(sanjay): write this test
}
