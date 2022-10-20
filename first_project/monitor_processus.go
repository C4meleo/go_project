package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode"
)

type Procs struct {
	PID  string
	Pwd  string
	Cwd  string
	Pss  string
	Rss  string
	Swap string
	Size string
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
			} else {
				proc.Pwd = "Need privileges"
			}
			target2, _ := os.Readlink("/proc/" + file.Name() + "/cwd")
			if len(target2) > 0 {
				proc.Cwd = target2
			} else {
				proc.Cwd = "Need privileges"
			}
			f, err := os.Open("/proc/" + file.Name() + "/smaps")
			//condition error
			if err != nil {
				proc.Size = "Size: \t\tNeed privileges"
				proc.Rss = "Rss: \t\tNeed privileges"
				proc.Pss = "Pss: \t\tNeed privileges"
				proc.Swap = "Swap: \t\tNeed privileges"
			}
			//defer to close the file
			defer f.Close()
			//read the file and affect each data to her structure variable and diplay them
			scanner := bufio.NewScanner(f)
			var line int = 0
			for scanner.Scan() {
				switch line {
				case 1:
					proc.Size = scanner.Text()
				case 4:
					proc.Rss = scanner.Text()
				case 5:
					proc.Pss = scanner.Text()
				case 18:
					proc.Swap = scanner.Text()
				}
				line++
			}
			procs = append(procs, proc)
		}
	}
	return procs
}

func print_all() {
	var procs []Procs
	procs = all_procs(procs)
	for _, proc := range procs {
		fmt.Println(" PID:\t\t", proc.PID, "\n PWD:\t\t", proc.Pwd, "\n CWD:\t\t", proc.Cwd, "\n", proc.Size, "\n", proc.Rss, "\n", proc.Pss, "\n", proc.Swap, "\n------")
	}
}

func main() {
	print_all()
}
