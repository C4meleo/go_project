package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode"
)

type Procs struct {
	PID string
	Pwd string
	Cwd string
}

func all_procs(procs []Procs) []Procs {
	var proc Procs
	files, err := ioutil.ReadDir("/proc/")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() && unicode.IsDigit(rune(file.Name()[0])) {
			proc.PID = file.Name()
			target, _ := os.Readlink("/proc/" + file.Name() + "/exe")
			if len(target) > 0 {
				proc.Pwd = target
			}
			target2, _ := os.Readlink("/proc/" + file.Name() + "/cwd")
			if len(target2) > 0 {
				proc.Cwd = target2
			}
			procs = append(procs, proc)
		}
	}
	return procs
}

func print_all() {
	var procs []Procs
	fmt.Println("PID\tEXE\tCWD\t% of use")
	procs = all_procs(procs)
	for _, proc := range procs {
		fmt.Println(proc.PID, proc.Pwd, proc.Cwd)
	}
}

func main() {
	print_all()
}
