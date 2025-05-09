package life_test

import (
  "fmt"
  "suslik/life"
  "testing"
)

func TestInstantiate(t *testing.T) {
  b1 := life.NewLife(4,4)
  b2 := life.NewLife(4,4)
  // These should be equivalent
  if ! life.AreEqual(b1,b2) {
     t.Error("blank 4x4 boards aren't the same")
  }
  b3 := life.NewLife(5,4)
  b4 := life.NewLife(4,5)
  if life.AreEqual(b1,b3) {
    t.Error("different size boards are the same")
  }
  if life.AreEqual(b1,b4) {
    t.Error("different size boards are the same")
  }
  // Add in a block to b1, b2 and compare
  life.AddBlock(0,0,b1)
  if life.AreEqual(b1,b2) {
    t.Error("one board has a block, blank board is equivalent")
  }
  life.AddBlock(0,0,b2)
  if ! life.AreEqual(b1,b2) {
    t.Error("2 boards, same block, unequal")
  }
  b5 := life.NewLife(5,5)
  life.AddGlider(0, 0, b5, life.DownRight)
  b6 := life.CopyLife(b5)
  if ! life.AreEqual(b5,b6) {
    t.Error("Copied boards aren't the same")
  }
  // A glider takes 4 cycles to move 1 block down and 1 block across
  // On a 5x5 board, it will take 5 x 4 cycles to completely cycle
  for i := 0 ; i< 19 ; i++ {
    life.Cycle(b5)
    if life.AreEqual(b5,b6) {
      t.Error(fmt.Sprintf("Glider cycle %d has looped, should not", i))
    }
  }
  life.Cycle(b5)
  if ! life.AreEqual(b5,b6) {
    t.Error("Glider on 5x5 board did not cycle with period 20")
  }
  // Create a new board and add a (down,left) glider in 2 ways
  b7 := life.NewLife(8,8)
  life.AddGlider(2, 2, b7, life.DownLeft)
  b8 := life.NewLife(8,8)
  if life.AreEqual(b7, b8) {
     t.Error("Empty board and glider board is the same")
  }
  gl := [3]string {" * ", "*  ", "***" }
  life.AddEntity(2, 2, b8, gl[:], life.Neutral)
  if ! life.AreEqual(b7,b8) {
     t.Error("Manual and standard downleft glider are different")
  }
  // Replace with a reflected glider
  life.AddEntity(2, 2, b8, gl[:], life.YReflection)
  if life.AreEqual(b7,b8) {
     t.Error("Reflected and standard downleft glider are the same")
  }
}
