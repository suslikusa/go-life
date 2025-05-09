package main

import (
   "fmt"
   "image"
   "image/gif"
   "os"
   "suslik/life"
   "time"
)


func main() {
   // A board, blank
   const x, y = 120, 100
   l := life.NewLife(x, y)
   // Static
   life.AddBlock(0,0,l)
   life.AddBlock(20,0,l)
   life.AddBlock(0,10,l)
   life.AddBeehive(10,10,l)
   life.AddBeehive(30,80,l)
   life.AddBoat(20,10,l)
   life.AddBoat(30,15,l)
   // Dynamic
   life.AddBlinker(10,15,l)
   life.AddToad(30,10,l)
   life.AddGlider(10, 0, l, life.DownRight)
   life.AddGlider(70, 0, l, life.DownRight)
   life.AddGlider(120, 0, l, life.DownLeft)
   life.AddGlider(20, 30, l, life.UpLeft)
   life.AddGlider(40, 28, l, life.UpRight)
   life.AddGliderGun(20,50, l, life.Neutral)
   life.AddGliderGun(70,40, l, life.YReflection)
 
   var imgs [] *image.Paletted
   var delays []int

   for c := 0 ; c < 1000 ; c++ {
     fmt.Printf("Cycle: %d\n", c)
     life.Cycle(l)
     //fmt.Println(l)
     // Generate an img with 6x6 pixels per cell
     imgs = append(imgs, life.Image(6,6,l))
     delays = append(delays, 0)
     time.Sleep(30*time.Millisecond)
   }
   // Turn the image into an animated GIF and write to file
   imgFile, err := os.Create("lifeimg.gif")
   if err != nil {
     fmt.Println("Failed to create PGIF file")
   } else {
     defer imgFile.Close()
     gif.EncodeAll(imgFile, &gif.GIF {
       Image: imgs,
       Delay: delays,
     })
   }
}

