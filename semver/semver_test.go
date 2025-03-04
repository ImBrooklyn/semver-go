package semver

import (
    "fmt"
    "reflect"
    "testing"
)

type testData struct {
    input    string
    expected *Version
}

var validVersions = []*testData{
    // valid inputs
    {
        input:    "0.0.4",
        expected: &Version{Major: 0, Minor: 0, Patch: 4},
    },
    {
        input:    "1.2.3",
        expected: &Version{Major: 1, Minor: 2, Patch: 3},
    },
    {
        input:    "10.20.30",
        expected: &Version{Major: 10, Minor: 20, Patch: 30},
    },
    {
        input:    "1.1.2-prerelease+meta",
        expected: &Version{Major: 1, Minor: 1, Patch: 2, Prerelease: "prerelease", Metadata: "meta"},
    },
    {
        input:    "1.1.2+meta",
        expected: &Version{Major: 1, Minor: 1, Patch: 2, Metadata: "meta"},
    },
    {
        input:    "1.1.2+meta-valid",
        expected: &Version{Major: 1, Minor: 1, Patch: 2, Metadata: "meta-valid"},
    },
    {
        input:    "1.0.0-alpha",
        expected: &Version{Major: 1, Minor: 0, Patch: 0, Prerelease: "alpha"},
    },
    {
        input:    "1.0.0-beta",
        expected: &Version{Major: 1, Minor: 0, Patch: 0, Prerelease: "beta"},
    },
    {
        input:    "1.0.0-alpha.beta",
        expected: &Version{Major: 1, Minor: 0, Patch: 0, Prerelease: "alpha.beta"},
    },
    {
        input:    "1.0.0-alpha.beta.1",
        expected: &Version{Major: 1, Minor: 0, Patch: 0, Prerelease: "alpha.beta.1"},
    },
    {
        input:    "1.0.0-alpha.1",
        expected: &Version{Major: 1, Minor: 0, Patch: 0, Prerelease: "alpha.1"},
    },
    {
        input:    "1.0.0-alpha0.valid",
        expected: &Version{Major: 1, Minor: 0, Patch: 0, Prerelease: "alpha0.valid"},
    },
    {
        input:    "1.0.0-alpha.0valid",
        expected: &Version{Major: 1, Minor: 0, Patch: 0, Prerelease: "alpha.0valid"},
    },
    {
        input: "1.0.0-alpha-a.b-c-somethinglong+build.1-aef.1-its-okay",
        expected: &Version{
            Major:      1,
            Minor:      0,
            Patch:      0,
            Prerelease: "alpha-a.b-c-somethinglong",
            Metadata:   "build.1-aef.1-its-okay",
        },
    },
    {
        input:    "1.0.0-rc.1+build.1",
        expected: &Version{Major: 1, Minor: 0, Patch: 0, Prerelease: "rc.1", Metadata: "build.1"},
    },
    {
        input:    "1.0.0-rc.1+build.123",
        expected: &Version{Major: 1, Minor: 0, Patch: 0, Prerelease: "rc.1", Metadata: "build.123"},
    },
    {
        input:    "1.2.3-beta",
        expected: &Version{Major: 1, Minor: 2, Patch: 3, Prerelease: "beta"},
    },
    {
        input:    "10.2.3-DEV-SNAPSHOT",
        expected: &Version{Major: 10, Minor: 2, Patch: 3, Prerelease: "DEV-SNAPSHOT"},
    },
    {
        input:    "1.2.3-SNAPSHOT-123",
        expected: &Version{Major: 1, Minor: 2, Patch: 3, Prerelease: "SNAPSHOT-123"},
    },
    {
        input:    "1.0.0",
        expected: &Version{Major: 1, Minor: 0, Patch: 0},
    },
    {
        input:    "2.0.0",
        expected: &Version{Major: 2, Minor: 0, Patch: 0},
    },
    {
        input:    "1.1.7",
        expected: &Version{Major: 1, Minor: 1, Patch: 7},
    },
    {
        input:    "2.0.0+build.1848",
        expected: &Version{Major: 2, Minor: 0, Patch: 0, Metadata: "build.1848"},
    },
    {
        input:    "2.0.1-alpha.1227",
        expected: &Version{Major: 2, Minor: 0, Patch: 1, Prerelease: "alpha.1227"},
    },
    {
        input:    "1.0.0-alpha+beta",
        expected: &Version{Major: 1, Minor: 0, Patch: 0, Prerelease: "alpha", Metadata: "beta"},
    },
    {
        input: "1.2.3----RC-SNAPSHOT.12.9.1--.12+788",
        expected: &Version{
            Major:      1,
            Minor:      2,
            Patch:      3,
            Prerelease: "---RC-SNAPSHOT.12.9.1--.12",
            Metadata:   "788",
        },
    },
    {
        input: "1.2.3----R-S.12.9.1--.12+meta",
        expected: &Version{
            Major:      1,
            Minor:      2,
            Patch:      3,
            Prerelease: "---R-S.12.9.1--.12",
            Metadata:   "meta",
        },
    },
    {
        input: "1.2.3----RC-SNAPSHOT.12.9.1--.12",
        expected: &Version{
            Major:      1,
            Minor:      2,
            Patch:      3,
            Prerelease: "---RC-SNAPSHOT.12.9.1--.12",
        },
    },
    {
        input: "1.0.0+0.build.1-rc.10000aaa-kk-0.1",
        expected: &Version{
            Major:    1,
            Minor:    0,
            Patch:    0,
            Metadata: "0.build.1-rc.10000aaa-kk-0.1",
        },
    },
    {
        input:    "1.0.0-0A.is.legal",
        expected: &Version{Major: 1, Minor: 0, Patch: 0, Prerelease: "0A.is.legal"},
    },
}

var invalidVersions = []string{
    "1",
    "1.2",
    "1.2.3-0123",
    "1.2.3-0123.0123",
    "1.1.2+.123",
    "+invalid",
    "-invalid",
    "-invalid+invalid",
    "-invalid.01",
    "alpha",
    "alpha.beta",
    "alpha.beta.1",
    "alpha.1",
    "alpha+beta",
    "alpha_beta",
    "alpha.",
    "alpha..",
    "beta",
    "1.0.0-alpha_beta",
    "-alpha.",
    "1.0.0-alpha..",
    "1.0.0-alpha..1",
    "1.0.0-alpha...1",
    "1.0.0-alpha....1",
    "1.0.0-alpha.....1",
    "1.0.0-alpha......1",
    "1.0.0-alpha.......1",
    "01.1.1",
    "1.01.1",
    "1.1.01",
    "1.2",
    "1.2.3.DEV",
    "1.2-SNAPSHOT",
    "1.2.31.2.3----RC-SNAPSHOT.12.09.1--..12+788",
    "1.2-RC-SNAPSHOT",
    "-1.0.3-gamma+b7718",
    "+justmeta",
    "9.8.7+meta+meta",
    "9.8.7-whatever+meta+meta",
    "99999999999999999999999.999999999999999999.99999999999999999",
    "99999999999999999999999.999999999999999999.99999999999999999----RC-SNAPSHOT.12.09.1--------------------------------..12",
}

func TestParse(t *testing.T) {
    type args struct {
        version string
    }

    type testcase struct {
        name    string
        args    args
        want    *Version
        wantErr bool
    }

    var tests []*testcase
    for i, v := range validVersions {
        tests = append(tests, &testcase{
            name:    fmt.Sprintf("ValidTestcase-%d", i),
            args:    args{version: v.input},
            want:    v.expected,
            wantErr: false,
        })
    }

    for i, v := range invalidVersions {
        tests = append(tests, &testcase{
            name:    fmt.Sprintf("InvalidTestcase-%d", i),
            args:    args{version: v},
            want:    nil,
            wantErr: true,
        })
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Parse(tt.args.version)
            if (err != nil) != tt.wantErr {
                t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Parse() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestVersion_String(t *testing.T) {
    type fields struct {
        Major      int64
        Minor      int64
        Patch      int64
        Prerelease string
        Metadata   string
    }

    type testcase struct {
        name   string
        fields fields
        want   string
    }

    var tests []*testcase

    for i, v := range validVersions {
        tests = append(tests, &testcase{
            name: fmt.Sprintf("Testcase-%d", i),
            fields: fields{
                Major:      v.expected.Major,
                Minor:      v.expected.Minor,
                Patch:      v.expected.Patch,
                Prerelease: v.expected.Prerelease,
                Metadata:   v.expected.Metadata,
            },
            want: v.input,
        })
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            v := &Version{
                Major:      tt.fields.Major,
                Minor:      tt.fields.Minor,
                Patch:      tt.fields.Patch,
                Prerelease: tt.fields.Prerelease,
                Metadata:   tt.fields.Metadata,
            }
            if got := v.String(); got != tt.want {
                t.Errorf("String() = %v, want %v", got, tt.want)
            }
        })
    }
}
