package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

/*
	"github.com/robfig/cron/v3"
	소스코드를 보니 Job들은 Gorutine으로 실행되는데 한번 테스트!
*/

func main() {
	c := cron.New(cron.WithLocation(time.UTC))

	var entryId cron.EntryID
	var entryId2 cron.EntryID

	entryId, _ = c.AddFunc("* * * * *", func() {
		curr := c.Entry(entryId).Prev.UTC()
		fmt.Printf("%d : %s\n", entryId, curr)
		time.Sleep(time.Minute*3)
		fmt.Printf("%d : %s\n", entryId, time.Now())
	})

	entryId2, _ = c.AddFunc("* * * * *", func() {
		curr := c.Entry(entryId2).Prev.UTC()
		fmt.Printf("%d : %s\n", entryId2, curr)
		time.Sleep(time.Minute*3)
		fmt.Printf("%d : %s\n", entryId2, time.Now())
	})

	c.Run()
}
