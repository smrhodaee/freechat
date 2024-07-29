package configs

import (
	"app/services/ratelimit"
	"time"
)

var ChaptaRatelimit = &ratelimit.Rule{
	Name: "chapta",
	Items: []ratelimit.Item{
		{
			Count:  10,
			Expire: time.Minute,
		},
	},
}
