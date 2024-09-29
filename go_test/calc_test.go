package calc

import (
    "testing"
)

func TestMax(t *testing.T) {
    got := Max(1, 2)
    want := 2
    if got != want {
        t.Errorf("Max(1, 2) == %d, want %d", got, want)
    }
}
