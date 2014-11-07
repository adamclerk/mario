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

type task interface {
	process()
	print()
}

type factory interface {
	make(line string, context *cli.Context) task
}

func Run(f factory, c *cli.Context) {

	if termutil.Isatty(os.Stdin.Fd()) {
		fmt.Println("No Piped Input Found!")
		cli.ShowAppHelp(c)
		return
	}

	var wg sync.WaitGroup

	in := make(chan task)

	wg.Add(1)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- f.make(s.Text(), c)
		}
		if s.Err() != nil {
			log.Fatalf("Error reading STDIN: %s", s.Err())
		}
		close(in)
		wg.Done()
	}()

	out := make(chan task)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for t := range in {
				t.process()
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
		t.print()
	}
}

type Command interface {
	Add() cli.Command
}
