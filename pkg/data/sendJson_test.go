// Copyright (c) 2021 Circutor S.A. All rights reserved.

package data_test

import (
	"net/http/httptest"
	"testing"

	"github.com/circutor/common-library/pkg/data"
)

func TestFailSendJson(t *testing.T) {
	t.Parallel()

	d := data.NewData()

	d.SendJSON(httptest.NewRecorder(), map[string]interface{}{"foo": make(chan int)})
}
