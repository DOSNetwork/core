package utils

import (
	"context"
	"fmt"
	"os/exec"
	"sync"
)

func AllowConnection(ip, port string) error {
	args := []string{"allow", "from", fmt.Sprintf("%s", ip), "to", "any", "port", port}
	out, err := exec.Command("ufw", args...).Output()
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		return err
	}
	fmt.Printf("ufw: %s\n", out)
	return nil
}

func ResetConnList() error {
	out, err := exec.Command("ufw", "--force", "reset").Output()
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		return err
	}
	fmt.Printf("ufw: %s\n", out)
	out, err = exec.Command("ufw", "allow", "22/tcp").Output()
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		return err
	}
	fmt.Printf("ufw: %s\n", out)
	out, err = exec.Command("ufw", "allow", "7946").Output()
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		return err
	}
	fmt.Printf("ufw: %s\n", out)
	out, err = exec.Command("ufw", "--force", "enable").Output()
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		return err
	}
	fmt.Printf("ufw: %s\n", out)
	return nil
}

func ReportError(ctx context.Context, errc chan error, err error) {
	select {
	case <-ctx.Done():
	case errc <- err:
	}
	return
}

func ReportResult(ctx context.Context, outc chan interface{}, result interface{}) {
	select {
	case <-ctx.Done():
	case outc <- result:
	}
	return
}

func MergeErrors(ctx context.Context, cs ...<-chan error) chan error {
	var wg sync.WaitGroup
	// We must ensure that the output channel has the capacity to
	// hold as many errors
	// as there are error channels.
	// This will ensure that it never blocks, even
	// if WaitForPipeline returns early.
	out := make(chan error, len(cs))
	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls
	// wg.Done.
	output := func(c <-chan error) {
		for n := range c {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	// Start a goroutine to close out once all the output goroutines
	// are done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
