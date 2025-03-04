// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package semver

import (
    "fmt"
    "testing"
)

type fixture struct {
    GreaterVersion string
    LessVersion    string
}

var fixtures = []fixture{
    {"0.0.0", "0.0.0-foo"},
    {"0.0.1", "0.0.0"},
    {"1.0.0", "0.9.9"},
    {"0.10.0", "0.9.0"},
    {"0.99.0", "0.10.0"},
    {"2.0.0", "1.2.3"},
    {"0.0.0", "0.0.0-foo"},
    {"0.0.1", "0.0.0"},
    {"1.0.0", "0.9.9"},
    {"0.10.0", "0.9.0"},
    {"0.99.0", "0.10.0"},
    {"2.0.0", "1.2.3"},
    {"0.0.0", "0.0.0-foo"},
    {"0.0.1", "0.0.0"},
    {"1.0.0", "0.9.9"},
    {"0.10.0", "0.9.0"},
    {"0.99.0", "0.10.0"},
    {"2.0.0", "1.2.3"},
    {"1.2.3", "1.2.3-asdf"},
    {"1.2.3", "1.2.3-4"},
    {"1.2.3", "1.2.3-4-foo"},
    {"1.2.3-5-foo", "1.2.3-5"},
    {"1.2.3-5", "1.2.3-4"},
    {"1.2.3-5-foo", "1.2.3-5-Foo"},
    {"3.0.0", "2.7.2+asdf"},
    {"3.0.0+foobar", "2.7.2"},
    {"1.2.3-a.10", "1.2.3-a.5"},
    {"1.2.3-a.b", "1.2.3-a.5"},
    {"1.2.3-a.b", "1.2.3-a"},
    {"1.2.3-a.b.c.10.d.5", "1.2.3-a.b.c.5.d.100"},
    {"1.0.0", "1.0.0-rc.1"},
    {"1.0.0-rc.2", "1.0.0-rc.1"},
    {"1.0.0-rc.1", "1.0.0-beta.11"},
    {"1.0.0-beta.11", "1.0.0-beta.2"},
    {"1.0.0-beta.2", "1.0.0-beta"},
    {"1.0.0-beta", "1.0.0-alpha.beta"},
    {"1.0.0-alpha.beta", "1.0.0-alpha.1"},
    {"1.0.0-alpha.1", "1.0.0-alpha"},
    {"1.2.3-rc.1-1-1hash", "1.2.3-rc.2"},
}

func TestCompare(t *testing.T) {
    for i, f := range fixtures {
        t.Run(fmt.Sprintf("Testcase-%d", i), func(t *testing.T) {
            gt, err := Parse(f.GreaterVersion)
            if err != nil {
                t.Fatalf("Parse(%q): %v", f.GreaterVersion, err)
            }
            lt, err := Parse(f.LessVersion)
            if err != nil {
                t.Fatalf("Parse(%q): %v", f.LessVersion, err)
            }

            if gt.LessThan(lt) {
                t.Errorf("%s should not be less than %s", gt, lt)
            }
            if gt.Equals(lt) {
                t.Errorf("%s should not be equal to %s", gt, lt)
            }
            if gt.Compare(lt) <= 0 {
                t.Errorf("%s should be greater than %s", gt, lt)
            }
            if !lt.LessThan(gt) {
                fmt.Println("!lt.LessThan(gt):", !lt.LessThan(gt))
                t.Errorf("%s should be less than %s", lt, gt)
            }
            if !lt.Equals(lt) {
                t.Errorf("%s should be equal to %s", lt, lt)
            }

            if !gt.Equals(gt) {
                t.Errorf("%s should be equal to %s", gt, gt)
            }

            if lt.Compare(gt) > 0 {
                t.Errorf("%s should not be greater than %s", lt, gt)
            }

        })
    }
}
