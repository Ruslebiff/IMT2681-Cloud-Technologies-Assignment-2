package assignment2

import (
	"testing"
	"time"
)

var Testid = ""

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

	var err error
	Testid, err = DBSave(&webhook)
	if err != nil {
		t.Error(err)
	}
}

func Test_DBReadid(t *testing.T) {
	webhook, err := DBReadid(Testid)
	if err != nil {
		t.Error(err)
	}

	if webhook.ID != Testid {
		t.Error(err)
	}
}

func Test_DBDelete(t *testing.T) {
	err := DBDelete(Testid)
	if err != nil {
		t.Error(err)
	}
}

func Test_DBReadall(t *testing.T) {
	var webhooks []Webhookreg    // temp storage for all webhooks from db
	webhooks, err := DBReadall() // read from db into webhooks array
	if err != nil {
		t.Error(err)
	}

	if webhooks == nil {
		t.Error(err)
	}
}

/*
func Test_DBClose(t *testing.T) {
	err := DBClose()
	if err != nil {
		t.Error(err)
	}
}
*/
