package fv

import (
	"fmt"
	"time"
)

type RequestGetWorkingOut struct {
	Id            int
	Date          time.Time
	SecurityLSKey string
}

func (r RequestGetWorkingOut) String() string {
	formatDate := r.Date.Format("2006-01-02")

	return fmt.Sprintf("id=%d&date=%s&security_ls_key=%s", r.Id, formatDate, r.SecurityLSKey)
}
