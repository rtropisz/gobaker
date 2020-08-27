package main

import (
	"flag"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/pkg/profile"
	"github.com/rtropisz/gobaker"
)

func main() {

	var (
		size            = flag.Int("s", 1024, "size of the output images in pixels")
		lowName         = flag.String("l", "", "path to lowpoly mesh")
		highName        = flag.String("h", "", "path to highpoly mesh")
		highPLYName     = flag.String("hp", "", "path to highpoly PLY mesh")
		cpuProfiling    = flag.Bool("cpuP", false, "turn on cpu profiling")
		memProfiling    = flag.Bool("memP", false, "turn on memory profiling")
		tracecProfiling = flag.Bool("traceP", false, "turn on trace profiling")
		useHalfCPU      = flag.Bool("useHalfCPU", true, "use half of available CPU cores, if set to false all you have")
	)
	flag.Parse()

	//Profiling
	if *cpuProfiling {
		defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	}
	if *memProfiling {
		defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	}
	if *tracecProfiling {
		defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	}

	workers := runtime.NumCPU()
	if *useHalfCPU {
		workers = runtime.NumCPU() / 2
	}
	scene := gobaker.NewScene(*size)

	log.Printf("Starting")
	start := time.Now()

	err := scene.Lowpoly.ReadOBJ(*lowName, false)
	if err != nil {
		log.Fatal(err)
	}

	err = scene.Highpoly.ReadOBJ(*highName, true)
	if err != nil {
		log.Fatal(err)
	}

	err = scene.Highpoly.ReadPLY(*highPLYName)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Started baking in %dx%d resolution", *size, *size)
	scene.Bake(workers)
	scene.BakedDiffuse.SaveImage(strings.TrimSuffix(*lowName, ".obj") + "_diff.png")
	scene.BakedID.SaveImage(strings.TrimSuffix(*lowName, ".obj") + "_id.png")
	// scene.BakedNormal.SaveImage(strings.TrimSuffix(*lowName, ".obj") + "_nrm.png")
	// scene.BakedObjectNormal.SaveImage(strings.TrimSuffix(*lowName, ".obj") + "_obj_nrm.png")
	log.Printf("Program finished in: %s", time.Since(start))
}
