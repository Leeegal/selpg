package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Selpg_args struct {
	start_page  int
	end_page    int
	in_filename string
	page_len    int
	page_type   int

	print_dest string
}

type sp_args Selpg_args

func main() {

	sa := Selpg_args{-1, -1, "\\0", 72, 0, "\\0"}
	ac := 0
	//proname := os.Args[0]

	for i, _ := range os.Args[1:] {
		//fmt.Printf("Argument %d is %s \n", i+1, a)
		ac = i + 1
		//fmt.Printf("%v \n", ac)
	}

	var s1 string
	var s2 string
	var argno int
	var i int
	max := 1<<32 - 1

	//----------------
	if ac < 3 {
		fmt.Fprintln(os.Stderr, "not enough arguments")
		os.Exit(1)
	}
	//-----------------
	s1 = os.Args[1]
	if s1[0:2] != "-s" {
		fmt.Fprintln(os.Stderr, "1st arg should be -sstart_page")
		os.Exit(2)
	}

	i, err := strconv.Atoi(s1[2:])
	if err != nil {
		panic(err)
	}

	if i < 1 || i > max {
		fmt.Fprintln(os.Stderr, "invalid start page")
		os.Exit(3)
	}
	sa.start_page = i
	//-----------------
	s2 = os.Args[2]
	if s2[0:2] != "-e" {
		fmt.Fprintln(os.Stderr, "2nd arg should be -eend_page")
		os.Exit(4)
	}
	i, err = strconv.Atoi(s2[2:])
	if err != nil {
		panic(err)
	}
	if i < 1 || i > max || i < sa.start_page {
		fmt.Fprintln(os.Stderr, "invalid end page")
		os.Exit(5)
	}
	sa.end_page = i
	//-------------------
	argno = 3
	for argno <= ac-1 && os.Args[argno][0] == '-' {
		s1 = os.Args[argno]

		switch s1[1:2] {
		//-------------
		case "l":
			s2 = s1[2:]
			i, err = strconv.Atoi(s2)
			if err != nil {
				panic(err)
			}

			if i < 1 || i > max {
				fmt.Fprintln(os.Stderr, "invalid page length")
				os.Exit(6)
			}
			argno++
			sa.page_len = i
			//--------------
		case "f":
			if s1[0:2] != "-f" {
				fmt.Fprintln(os.Stderr, "option should be \"-f\"")
				os.Exit(7)
			}
			argno++
			sa.page_type = 2
			//---------------
		case "d":
			s2 = s1[2:]
			if len(s2) < 1 {
				fmt.Fprintln(os.Stderr, "-d option requires a printer destination")
				os.Exit(8)
			}
			argno++
			sa.print_dest = s2

			//-----------------
		default:
			fmt.Fprintln(os.Stderr, "unknown option ")
			os.Exit(9)
		}
	}
	//----------------------------------
	if sa.print_dest[0:2] != "\\0" {

		s1 = "lp -d" + sa.print_dest
		cmd := exec.Command(s1)
		//fmt.Println(cmd.Args)

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Start()
		r := bufio.NewReader(stdout)
		line, _, err := r.ReadLine()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Println(line)
	}
	//------------------------------------
	if argno <= ac {

		var pagen = 1
		var col = 0
		sa.in_filename = os.Args[argno]
		fin, err := os.Open(sa.in_filename)
		defer fin.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, "cannot find or open the file")
			os.Exit(10)
		}
		if sa.in_filename[0:1] == "\\0" {
			fin = os.Stdin
		} else {
			buf := bufio.NewReader(fin)
			for {
				line, err := buf.ReadString('\n')
				line = strings.TrimSpace(line)
				col++

				if sa.page_type == 0 {
					if pagen >= sa.start_page && pagen <= sa.end_page {
						fmt.Println(line)
					}
					if col == sa.page_len {
						pagen++
						col = 0
					}

					if err != nil {
						if err == io.EOF {
							break
						}
					}
				} else {
					if col == sa.page_len || line == "\\f" {
						pagen++
						col = 0
					}
					if line != "\\f" && pagen >= sa.start_page && pagen <= sa.end_page {
						fmt.Println(line)
					}
					if err != nil {
						if err == io.EOF {
							break
						}
					}
				}
			}
			//fmt.Println(sa)
			if pagen < sa.end_page {
				fmt.Fprintln(os.Stderr, "the end_page is more than the total pages")
			}
		}
	}

	//------------------------------------------------------------------
}
