package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	group, _ := errgroup.WithContext(context.Background())
	for i := 0; i < 10; i++ {
		index := i
		group.Go(func() error {
			fmt.Printf("start to execute the %d gorouting\n", index)
			time.Sleep(time.Duration(index) * time.Second)
			if index%2 == 0 {
				return fmt.Errorf("somethind has failed on grouting %d", index)
			}
			fmt.Println("gorouting:%d end \n", index)
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("main over")
}
