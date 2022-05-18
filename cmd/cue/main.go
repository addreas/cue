// Copyright 2018 The CUE Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"cuelang.org/go/cmd/cue/cmd"
)

func main() {
	fc, err := os.Create("/tmp/cpupprof")
	if err != nil {
		log.Fatal(err)
	}
	fm, err := os.Create("/tmp/mempprof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(fc)
	res := cmd.Main()
	pprof.StopCPUProfile()
	runtime.GC()
	if err := pprof.WriteHeapProfile(fm); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
	fc.Close()
	fm.Close()
	os.Exit(res)
}
