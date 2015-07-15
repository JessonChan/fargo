// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fargo

import (
	"testing"

	"net/http"

	"github.com/astaxie/beego/context"
)

func TestGetInt_1(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	i := &context.BeegoInput{Params: map[string]string{"age": "40"}, Request: r}
	ctx := &context.Context{Input: i}
	ctrlr := Controller{Ctx: ctx}

	//	val, _ := ctrlr.GetInt("age")
	//	if val != 40 {
	//		t.Fail()
	//	}
	val, err := ctrlr.GetInt("age2", 19)
	t.Log(val, err)
	if val != 19 {
		t.Fail()
	}
	//Output: int
}
