package cron

import (
        "time"
        "github.com/domeos/sender/g"
)
func UpdateApiConfig() {

        for {
                g.UpdateApiConfig()
                time.Sleep(time.Millisecond * 60000)
        }
}
