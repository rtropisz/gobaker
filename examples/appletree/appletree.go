package main

import (
	"log"
	"strings"
	"time"

	"github.com/rtropisz/gobaker"
)

const (
	size        = 1024
	lowName     = "./AppleTree_lowpoly.obj"
	highName    = "./AppleTree.obj"
	highPlyName = "./AppleTree.ply"
)

func main() {
	scene := gobaker.NewScene(size)
	log.Printf("Starting")
	start := time.Now()
	scene.Lowpoly.ReadOBJ(lowName, false)
	log.Printf("Readed lowpoly mesh.. %s", time.Since(start))
	scene.Highpoly.ReadOBJ(highName, true)
	scene.Highpoly.ReadPLY(highPlyName)

	log.Printf("Readed highpoly mesh.. %s", time.Since(start))
	scene.Bake()
	log.Printf("Baking...")
	log.Printf("Finished baking: %s", time.Since(start))
	scene.BakedDiffuse.SaveImage(strings.TrimSuffix(lowName, ".obj") + "_diff.png")
	scene.BakedID.SaveImage(strings.TrimSuffix(lowName, ".obj") + "_id.png")
	log.Printf("Finished saving images: %s", time.Since(start))
}
