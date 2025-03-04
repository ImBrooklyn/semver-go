package semver

import (
    "bytes"
    "fmt"
    "regexp"
    "strconv"
)

var re = regexp.MustCompile(`^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<metadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

type Version struct {
    Major      int64
    Minor      int64
    Patch      int64
    Prerelease string
    Metadata   string
}

func (v *Version) String() string {
    var buffer bytes.Buffer

    _, _ = fmt.Fprintf(&buffer, "%d.%d.%d", v.Major, v.Minor, v.Patch)

    if v.Prerelease != "" {
        _, _ = fmt.Fprintf(&buffer, "-%s", v.Prerelease)
    }

    if v.Metadata != "" {
        _, _ = fmt.Fprintf(&buffer, "+%s", v.Metadata)
    }

    return buffer.String()
}

// Parse semantic version from string
func Parse(version string) (*Version, error) {
    if !re.MatchString(version) {
        return nil, fmt.Errorf("invalid semantic version: %s", version)
    }
    matched := re.FindStringSubmatch(version)
    major, err := strconv.ParseInt(getCapture("major", matched), 10, 64)
    if err != nil {
        return nil, err
    }
    minor, err := strconv.ParseInt(getCapture("minor", matched), 10, 64)
    if err != nil {
        return nil, err
    }
    patch, err := strconv.ParseInt(getCapture("patch", matched), 10, 64)
    if err != nil {
        return nil, err
    }
    prerelease := getCapture("prerelease", matched)
    metadata := getCapture("metadata", matched)

    return &Version{
        Major:      major,
        Minor:      minor,
        Patch:      patch,
        Prerelease: prerelease,
        Metadata:   metadata,
    }, nil
}

// getCapture gets name capture from matched result
// prevents index out of range
func getCapture(name string, matched []string) string {
    if len(matched) == 0 {
        return ""
    }
    index := re.SubexpIndex(name)
    if index != -1 && index < len(matched) {
        return matched[index]
    }
    return ""
}
