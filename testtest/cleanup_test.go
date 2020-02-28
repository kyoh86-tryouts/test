package testtest

import "testing"

func TestCleanup(t *testing.T) {
	t.Run("subtest", func(t *testing.T) {
		t.Cleanup(func() {
			t.Log(1)
		})
	})
	t.Cleanup(func() {
		t.Log(2)
	})
	t.Run("subtest-2", func(t *testing.T) {
		t.Cleanup(func() {
			t.Log(3)
		})
	})
}
