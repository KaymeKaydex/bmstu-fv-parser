package fv_test

import (
	"testing"
	"time"

	"github.com/magiconair/properties/assert"

	"github.com/KaymeKaydex/bmstu-fv-parser.git/internal/pkg/clients/fv"
)

func TestRequestGetWorkingOut_String(t *testing.T) {
	testCase := &fv.RequestGetWorkingOut{
		Id:            3,
		Date:          time.Date(2001, 11, 12, 0, 0, 0, 0, time.Local),
		SecurityLSKey: "123",
	}

	assert.Equal(t, testCase.String(), "id=3&date=2001-11-12&security_ls_key=123")
}
