/*
Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mysqlctl

import (
	"reflect"
	"testing"

	"github.com/youtube/vitess/go/mysql"
)

func TestMysql56VersionMatch(t *testing.T) {
	table := map[string]bool{
		"10.0.13-MariaDB-1~precise-log": false,
		"5.1.63-google-log":             false,
		"5.6.24-log":                    true,
	}
	for input, want := range table {
		if got := (&mysql56{}).VersionMatch(input); got != want {
			t.Errorf("(&mysql56{}).VersionMatch(%#v) = %v, want %v", input, got, want)
		}
	}
}

func TestMysql56MakeBinlogEvent(t *testing.T) {
	input := []byte{1, 2, 3}
	want := mysql.NewMysql56BinlogEvent([]byte{1, 2, 3})
	if got := (&mysql56{}).MakeBinlogEvent(input); !reflect.DeepEqual(got, want) {
		t.Errorf("(&mysql56{}).MakeBinlogEvent(%#v) = %#v, want %#v", input, got, want)
	}
}
