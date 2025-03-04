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
    "strconv"
    "strings"
)

// Compare tests if v is less than, equal to, or greater than other,
// returning -1, 0, or +1 respectively.
func (v *Version) Compare(other *Version) int {
    if cmp := compareNumbers(
        []int64{v.Major, v.Minor, v.Patch},
        []int64{other.Major, other.Minor, other.Patch},
    ); cmp != 0 {
        return cmp
    }

    return comparePrerelease(v.Prerelease, other.Prerelease)
}

// Equals tests if v equal to other
func (v *Version) Equals(other *Version) bool {
    return v.Compare(other) == 0
}

// LessThan tests if v is less than other
func (v *Version) LessThan(other *Version) bool {
    return v.Compare(other) < 0
}

// GreaterThan tests if v is greater than other
func (v *Version) GreaterThan(other *Version) bool {
    return v.Compare(other) > 0
}

// compareNumbers recursively compares number parts of lhs and rhs
func compareNumbers(lhs, rhs []int64) int {
    if len(lhs) == 0 {
        return 0
    }

    l := lhs[0]
    r := rhs[0]

    if l > r {
        return 1
    } else if l < r {
        return -1
    }

    return compareNumbers(lhs[1:], rhs[1:])
}

// comparePrerelease recursively compares prerelease of lhs and rhs
func comparePrerelease(lhs, rhs string) int {
    if len(lhs) == 0 && len(rhs) > 0 {
        return 1
    } else if len(lhs) > 0 && len(rhs) == 0 {
        return -1
    }

    return comparePrereleaseParts(
        strings.Split(lhs, "."),
        strings.Split(rhs, "."),
    )
}

// comparePrereleaseParts recursively compares all parts of prerelease
// ref: github.com/coreos/go-semver
func comparePrereleaseParts(versionA, versionB []string) int {
    // A larger set of pre-release fields has a higher precedence than a smaller set,
    // if all the preceding identifiers are equal.
    if len(versionA) == 0 {
        if len(versionB) > 0 {
            return -1
        }
        return 0
    } else if len(versionB) == 0 {
        // We're longer than versionB so return 1.
        return 1
    }

    a := versionA[0]
    b := versionB[0]

    aInt := false
    bInt := false

    aI, err := strconv.Atoi(versionA[0])
    if err == nil {
        aInt = true
    }

    bI, err := strconv.Atoi(versionB[0])
    if err == nil {
        bInt = true
    }

    // Numeric identifiers always have lower precedence than non-numeric identifiers.
    if aInt && !bInt {
        return -1
    } else if !aInt && bInt {
        return 1
    }

    // Handle Integer Comparison
    // Equivalent to aInt && bInt, because aInt == bInt
    if aInt {
        if aI > bI {
            return 1
        } else if aI < bI {
            return -1
        }
    }

    // Handle String Comparison
    if a > b {
        return 1
    } else if a < b {
        return -1
    }

    return comparePrereleaseParts(versionA[1:], versionB[1:])
}
