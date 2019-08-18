package main

import "testing"

func TestDist(t *testing.T) {
    total := dist("GAMBOL","GUMBO")
    if total != 2 {
       t.Errorf("Levenshtein distance was incorrect, got: %d, want: %d.", total, 2)
    }
    //"Zeil","trials" 4
     total = dist("zeil","trials")
    if total != 4 {
       t.Errorf("Levenshtein distance was incorrect, got: %d, want: %d.", total, 4)
    }
    //"ала","алаб"
     total = dist("ала","алаб")
    if total != 1 {
       t.Errorf("Levenshtein distance was incorrect, got: %d, want: %d.", total, 1)
    }
}