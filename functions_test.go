package assignment2

import (
	"testing"
	"time"
)

func Test_DBInit(t *testing.T) {
	err := DBInit()
	if err != nil {
		t.Error(err)
	}
}

func Test_DBSave(t *testing.T) {
	webhook := Webhookreg{}

	webhook.Event = "commits"
	webhook.Time = time.Now()

	_, err := DBSave(&webhook)
	if err != nil {
		t.Error(err)
	}
}
