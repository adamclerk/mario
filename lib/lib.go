package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/andrew-d/go-termutil"
	"github.com/codegangsta/cli"
)

type Task interface {
	Process()
	Print()
}

type factory interface {
	Make(line string, context *cli.Context) Task
}

func Run(f factory, c *cli.Context) {

	if termutil.Isatty(os.Stdin.Fd()) {
		fmt.Println("No Piped Input Found!")
		cli.ShowAppHelp(c)
		return
	}

	var wg sync.WaitGroup

	in := make(chan Task)

	wg.Add(1)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- f.Make(s.Text(), c)
		}
		if s.Err() != nil {
			log.Fatalf("Error reading STDIN: %s", s.Err())
		}
		close(in)
		wg.Done()
	}()

	out := make(chan Task)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for t := range in {
				t.Process()
				out <- t
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for t := range out {
		t.Print()
	}
}

type Command interface {
	Add() cli.Command
}
