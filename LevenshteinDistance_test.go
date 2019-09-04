package main

import "testing"

func TestDist(t *testing.T) {
  
  var tests = []struct{
    word1 string;word2 string;result int
  }{
    {word1:"GAMBOL",word2:"GUMBO",result:2},
    {word1:"zeil",word2:"trials",result:4},
    {word1:"ала",word2:"алаб",result:1},
  }

  for _, test := range tests {
    total := dist(test.word1,test.word2)
    if total != test.result {
      t.Errorf("Levenshtein distance was incorrect, got: %d, want: %d.", total, test.result)
    }
  }
}