// Copyright 2020 FastWeGo
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

package marketcode_test

import (
	"fmt"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/marketcode"
)

func ExampleApplyCode() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := marketcode.ApplyCode(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleApplyCodeQuery() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := marketcode.ApplyCodeQuery(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleApplyCodeDownload() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := marketcode.ApplyCodeDownload(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCodeActive() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := marketcode.CodeActive(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCodeActiveQuery() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := marketcode.CodeActiveQuery(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleTicketToCode() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := marketcode.TicketToCode(ctx, payload)

	fmt.Println(resp, err)
}
