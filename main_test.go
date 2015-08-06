package main

import (
  "testing"
)

func TestTruth(t *testing.T) {
  if t == nil {
    t.Fatalf("truth fails")
  }
}
