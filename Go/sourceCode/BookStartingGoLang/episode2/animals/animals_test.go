package animals

import (
  "testing"
)

func TestElephantFeed(t *testing.T) {
  expect := "Grass"
  actual := ElephantFeed()

  if expect != actual {
    t.Errorf("%s != %s", expect, actual)
  }
}

func TestMonky(t *testing.T) {
  expect := "Banana"
  actual := MonkeyFeed()

  if expect != actual {
    t.Errorf("%s != %s", expect, actual)
  }
}
