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
	"encoding/json"

	"github.com/astaxie/beego/context"
)

// GlobalDocAPI store the swagger api documents
var GlobalDocAPI map[string]interface{}

func init() {
	if EnableDocs {
		GlobalDocAPI = make(map[string]interface{})
	}
}

func serverDocs(ctx *context.Context) {
	var obj interface{}
	if splat := ctx.Input.Param(":splat"); splat == "" {
		obj = GlobalDocAPI["Root"]
	} else {
		if v, ok := GlobalDocAPI[splat]; ok {
			obj = v
		}
	}
	if obj != nil {
		bt, err := json.Marshal(obj)
		if err != nil {
			ctx.Output.SetStatus(504)
			return
		}
		ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
		ctx.Output.Body(bt)
		return
	}
	ctx.Output.SetStatus(404)
}
