package preview

import (
	"testing"

	"github.com/k-kohey/axe/internal/preview/analysis"
)

func TestSharedIndexCacheFreshTracksStaleness(t *testing.T) {
	c := newSharedIndexCache(nil)
	if c.Fresh() != nil {
		t.Fatal("empty cache must not be fresh")
	}

	loaded := &analysis.IndexStoreCache{}
	c.Set(loaded)
	if c.Fresh() != loaded {
		t.Fatal("cache must be fresh right after Set")
	}

	c.MarkStale()
	if c.Fresh() != nil {
		t.Fatal("stale cache must not be reported fresh")
	}
	if c.Get() != loaded {
		t.Fatal("stale cache must stay readable via Get")
	}

	reloaded := &analysis.IndexStoreCache{}
	c.Set(reloaded)
	if c.Fresh() != reloaded {
		t.Fatal("Set must clear the stale flag")
	}
}
