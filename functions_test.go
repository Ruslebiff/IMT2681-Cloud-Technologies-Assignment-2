package assignment2

import "testing"

func Test_DBInit(t *testing.T) {
	err := DBInit()
	if err != nil {
		t.Error(err)
	}
}
