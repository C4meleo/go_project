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
	//reads all files located in /proc
	files, err := ioutil.ReadDir("/proc/")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		//all files whose name begins with a number
		if file.IsDir() && unicode.IsDigit(rune(file.Name()[0])) {
			proc.PID = file.Name()
			//read the file exe relative to the file contained in the variable file.Name()
			target, _ := os.Readlink("/proc/" + file.Name() + "/exe")
			if len(target) > 0 {
				proc.Pwd = target
			} else {
				proc.Pwd = "Need privileges"
			}
			//read the file cwd relative to the file contained in the variable file.Name()
			target2, _ := os.Readlink("/proc/" + file.Name() + "/cwd")
			if len(target2) > 0 {
				proc.Cwd = target2
			} else {
				proc.Cwd = "Need privileges"
			}
			//read the file smaps relative to the file name contained in the variable file.Name()
			f, err := os.Open("/proc/" + file.Name() + "/smaps")
			//condition error if we don't have the permission to read the file
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
			//append to procs any new data added to proc
			procs = append(procs, proc)
		}
	}
	return procs
}

func pid() string {
	var res string
	//procs is a structure table
	var procs []Procs
	//append to procs all new data
	procs = all_procs(procs)
	//print all data
	for _, proc := range procs {
		res += " PID:\t\t" + proc.PID + "\n PWD:\t\t" + proc.Pwd + "\n CWD:\t\t" + proc.Cwd + "\n" + proc.Size + "\n" + proc.Rss + "\n" + proc.Pss + "\n" + proc.Swap + "\n------\n"
	}
	return res
}

func main() {
	fmt.Println(pid())
}
