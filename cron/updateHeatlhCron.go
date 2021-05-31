package cron

import (
	"context"
	"fmt"
	"load-balancer/pool"

	"github.com/robfig/cron/v3"
)

func UpdateHealthCron() {
	c := cron.New()
	c.AddFunc("@every 5s", func() {
		fmt.Println("running cron after 5 sec")
		ctx := context.Background()
		pool := pool.GetServerPool()
		for _, server := range pool.Servers {
			err := GetHealth(ctx, server.URL)
			if err != nil {
				fmt.Printf("server is down: %s\n", server.URL.Host)
				server.SetHealth(true)
			} else {
				server.SetHealth(false)
			}
		}
	})
	c.Start()
}
