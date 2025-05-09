package life

import (
  //"fmt"
  "image"
  "image/color"
  "math/rand"
)

// Track cell liveness as an int for ease of summing
type liferow []int
type neighborcount []int

type Lifeboard struct {
   max_x, max_y int
   alive []liferow
   neighbors []neighborcount
}

// Create a new x by y Lifeboard, initially blank
func NewLife(x, y int) *Lifeboard {
  l := Lifeboard { x, y, make([]liferow, y), make([]neighborcount, y) }
  for j := 0; j < y; j++ {
    l.alive[j] = make(liferow, x)
    l.neighbors[j] = make(neighborcount, x)
    for i := 0; i < x ; i++ {        
        l.alive[j][i] = 0
      }
  }
  return &l 
} 

// Copy a board
func CopyLife(l *Lifeboard) *Lifeboard {
  l2 := NewLife(l.max_x, l.max_y)
  for j := 0 ; j < l.max_y ; j++ {
    for i := 0 ; i < l.max_x ; i++ {
      l2.alive[j][i] = l.alive[j][i]
    }
  }
  return l2
}

// Randomize content of existing board.
// Threshold is % chance of a live cell
func Randomize(l *Lifeboard, threshold int) {
  for j := 0; j < l.max_y; j++ {
    for i := 0 ; i < l.max_x ; i++ {
      if rand.Intn(100) < threshold {
        l.alive[j][i] = 1
      } else {
        l.alive[j][i] = 0
      }
    } // i
  } // j
}

// Stringifier
func (l *Lifeboard) String() string {
   var s string
   for j := 0 ; j < l.max_y ; j++ {
      for i := 0 ; i < l.max_x ; i++ {
         if l.alive[j][i] > 0 {
            s = s + "*"
         } else {
            s = s + "."
         }
      }
      s = s + "\n"
   }
   return s
}

type rowcount []int

// Cycle a board
func Cycle(l * Lifeboard) {
  UpdateNeighbors(l)
  for thisrow := 0 ; thisrow < l.max_y ; thisrow++ {
    for thiscol := 0 ; thiscol < l.max_x ; thiscol++ {
       // Game of life rules
       if l.neighbors[thisrow][thiscol] < 2 {
         // Any live cell with fewer than two live neighbours dies, as if caused by under-population.
         l.alive[thisrow][thiscol] = 0
       } else if (l.alive[thisrow][thiscol] > 0 && l.neighbors[thisrow][thiscol] > 3) {
         // Any live cell with more than three live neighbours dies, as if by over-population.
         l.alive[thisrow][thiscol] = 0
       } else if l.neighbors[thisrow][thiscol] == 3 {
         // Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
         l.alive[thisrow][thiscol] = 1
       } 
       // Any live cell with two or three live neighbours lives on to the next generation.
    } // thiscol
  } // thisrow
}

func UpdateNeighbors(l *Lifeboard) {
  // Update the neighbor count matrix
  var above, below, leftcol, rightcol, neighbors int
  for thisrow := 0 ; thisrow < l.max_y; thisrow++ {
     above = (l.max_y + thisrow - 1) % l.max_y
     below = (thisrow + 1) % l.max_y
     for thiscol := 0; thiscol < l.max_x ; thiscol++ {
       // Manual count left to right, top to bottom
       neighbors = 0
       leftcol = (l.max_x + thiscol - 1) % l.max_x
       rightcol = (thiscol + 1) % l.max_x
       neighbors += l.alive[above][leftcol]
       neighbors += l.alive[above][thiscol]
       neighbors += l.alive[above][rightcol]
       neighbors += l.alive[thisrow][leftcol]
       neighbors += l.alive[thisrow][rightcol]
       neighbors += l.alive[below][leftcol]
       neighbors += l.alive[below][thiscol]
       neighbors += l.alive[below][rightcol]
       l.neighbors[thisrow][thiscol] = neighbors
    } // thiscol
  } // thisrow
}


// Life entities

func AddBlock(x int, y int, l *Lifeboard) {
  // 2 x 2 block with top left at x,y
  x = (x + l.max_x) % l.max_x
  y = (y + l.max_y) % l.max_y
  x1 := (x + 1) % l.max_x
  y1 := (y + 1) % l.max_y
  l.alive[y][x] = 1
  l.alive[y][x1] = 1
  l.alive[y1][x] = 1
  l.alive[y1][x1] = 1
}

func AddBeehive(x int, y int, l *Lifeboard) {
  // 4 x 3 block with top left at x,y
  x = (x + l.max_x) % l.max_x
  y = (y + l.max_y) % l.max_y
  x1 := (x + 1) % l.max_x
  x2 := (x + 2) % l.max_x
  x3 := (x + 3) % l.max_x
  y1 := (y + 1) % l.max_y
  y2 := (y + 2) % l.max_y
  l.alive[y][x] = 0
  l.alive[y][x1] = 1
  l.alive[y][x2] = 1
  l.alive[y][x3] = 0
  l.alive[y1][x] = 1
  l.alive[y1][x1] = 0
  l.alive[y1][x2] = 0
  l.alive[y1][x3] = 1
  l.alive[y2][x] = 0
  l.alive[y2][x1] = 1
  l.alive[y2][x2] = 1
  l.alive[y2][x3] = 0
}

func AddBoat(x int, y int, l *Lifeboard) {
  // 3x3 block with top left at x,y
  x = (x + l.max_x) % l.max_x
  y = (y + l.max_y) % l.max_y
  x1 := (x + 1) % l.max_x
  x2 := (x + 2) % l.max_x
  y1 := (y + 1) % l.max_y
  y2 := (y + 2) % l.max_y
  l.alive[y][x] = 1
  l.alive[y][x1] = 1
  l.alive[y][x2] = 0
  l.alive[y1][x] = 1
  l.alive[y1][x1] = 0
  l.alive[y1][x2] = 1
  l.alive[y2][x] = 0
  l.alive[y2][x1] = 1
  l.alive[y2][x2] = 0
}

func AddBlinker(x int, y int, l *Lifeboard) {
  // 3x3 block with top left at x,y
  x = (x + l.max_x) % l.max_x
  y = (y + l.max_y) % l.max_y
  x1 := (x + 1) % l.max_x
  x2 := (x + 2) % l.max_x
  y1 := (y + 1) % l.max_y
  y2 := (y + 2) % l.max_y
  l.alive[y][x] = 0
  l.alive[y][x1] = 1
  l.alive[y][x2] = 0
  l.alive[y1][x] = 0
  l.alive[y1][x1] = 1
  l.alive[y1][x2] = 0
  l.alive[y2][x] = 0
  l.alive[y2][x1] = 1
  l.alive[y2][x2] = 0
}

func AddToad(x int, y int, l *Lifeboard) {
  // 4x2 block with top left at x,y
  x = (x + l.max_x) % l.max_x
  y = (y + l.max_y) % l.max_y
  x1 := (x + 1) % l.max_x
  x2 := (x + 2) % l.max_x
  x3 := (x + 3) % l.max_x
  y1 := (y + 1) % l.max_y
  l.alive[y][x] = 0
  l.alive[y][x1] = 1
  l.alive[y][x2] = 1
  l.alive[y][x3] = 1
  l.alive[y1][x] = 1
  l.alive[y1][x1] = 1
  l.alive[y1][x2] = 1
  l.alive[y1][x3] = 0
}

// Orientations
type DiagOrientation int
const (
  DownRight = iota
  DownLeft = iota
  UpRight = iota
  UpLeft = iota
)
type HorizOrientation int
const (
  Neutral = iota
  XReflection = iota
  YReflection = iota
  XYReflection = iota
)

func AddGlider(x int, y int, l *Lifeboard, o DiagOrientation) {
  // 3x3 block with top left at x,y 
  // Flying down and right
  x = (x + l.max_x) % l.max_x
  y = (y + l.max_y) % l.max_y
  x1 := (x + 1) % l.max_x
  x2 := (x + 2) % l.max_x
  y1 := (y + 1) % l.max_y
  y2 := (y + 2) % l.max_y
  // Handle orientation
  if o == DownLeft || o == UpLeft {
    // Reflect in X axis
    x2, x = x, x2     
  }
  if o == UpLeft || o == UpRight {
    // Reflect in Y axis
    y2, y = y, y2
  }
  l.alive[y][x] = 0
  l.alive[y][x1] = 1
  l.alive[y][x2] = 0
  l.alive[y1][x] = 0
  l.alive[y1][x1] = 0
  l.alive[y1][x2] = 1
  l.alive[y2][x] = 1
  l.alive[y2][x1] = 1
  l.alive[y2][x2] = 1
}

// Arbitrary entities represented as lists of strings, "*" is a live cell
func AddEntity(x int, y int, l *Lifeboard, ent []string, o HorizOrientation) {
   var thisrow, thiscol int
   for e := 0 ; e < len(ent) ; e++ {
     if (o == Neutral || o == XReflection) {
       thisrow = (y + e) % l.max_y
     } else {
       thisrow = (y + len(ent) - e) % l.max_y
     }
     for index, runeval := range ent[e] {
       if (o == Neutral || o == YReflection) {
         thiscol = (x + index) % l.max_x
       } else {
         thiscol = (x + len(ent[e]) - index) % l.max_x
       }
       if runeval == 0x2a { // '*'
         l.alive[thisrow][thiscol] = 1
       } else {
         l.alive[thisrow][thiscol] = 0
       }
     }
   }
}

func AddGliderGun(x int, y int, l *Lifeboard, o HorizOrientation) {
  // Gosper glider gun is 38 x 11, including a blank single char border
  gg := [11]string{
    //01234567890123456789012345678901234567
     "                                      ",
     "                         *            ",
     "                       * *            ",
     "             **      **            ** ",
     "            *   *    **            ** ",
     " **        *     *   **               ",
     " **        *   * **    * *            ",
     "           *     *       *            ",
     "            *   *                     ",
     "             **                       ",
     "                                      ",
   }
  AddEntity(x, y, l, gg[:], o)
}

// Comparison

func AreEqual(l1 *Lifeboard, l2 *Lifeboard) bool {
  if l1.max_x != l2.max_x {
    return false
  }
  if l1.max_y != l2.max_y {
    return false
  }
  // Sizes the same, compare contents
  for j := 0; j < l1.max_y ; j++ {
    for i := 0; i < l1.max_x ; i++ {
      if l1.alive[j][i] != l2.alive[j][i] {
        return false
      }
    }
  }
  // Everything the same
  return true
}

// Image construction

func Image(scale_x int, scale_y int, l *Lifeboard) *image.Paletted {
   // Return an RGBA image representing the Life board, with 
   // each Life cell occupying scale_x by scale_y pixels
   r := image.Rectangle { image.Point {0, 0}, image.Point {scale_x * l.max_x , scale_y * l.max_y} } 
   p := []color.Color{
        color.RGBA{0x00, 0x00, 0x00, 0xff},
        color.RGBA{0x00, 0x00, 0xff, 0xff},
        color.RGBA{0x00, 0xff, 0x00, 0xff},
        color.RGBA{0x00, 0xff, 0xff, 0xff},
        color.RGBA{0xff, 0x00, 0x00, 0xff},
        color.RGBA{0xff, 0x00, 0xff, 0xff},
        color.RGBA{0xff, 0xff, 0x00, 0xff},
        color.RGBA{0xff, 0xff, 0xff, 0xff},
   }
   img := image.NewPaletted(r, p)
   c := color.NRGBA { 0xff, 0x40, 0x20, 0x80 }
   for j := 0 ; j < l.max_y ; j++ {
     for i := 0; i < l.max_x ; i++ {
       if l.alive[j][i] > 0 {
         for sy := 0 ; sy < scale_y ; sy++ {
           for sx := 0 ; sx < scale_x ; sx++ {
             img.Set( i * scale_x + sx, j * scale_y + sy, c)
             //fmt.Printf("setting %d,%d\n", i*scale_x + sx, j*scale_y + sy)
           } // sx
         } // sy
       } // l.alive
     } // i
  } // j
  return img
}