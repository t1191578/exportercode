package main

import (
  "testing"
  "os"
  "io/ioutil"
)

func TestMemoryMetrics(t *testing.T) {
  // Read the input data from a file
  file, err := os.Open("test_data/sinfo_Memory.txt")
  if err != nil { t.Fatalf("Can not open test data: %v", err) }
  data, err := ioutil.ReadAll(file)
  t.Logf("%+v", ParseMemoryMetrics(data))
}

func TestMemorysGetMetrics(t *testing.T) {
  t.Logf("%+v", MemoryGetMetrics())
}
