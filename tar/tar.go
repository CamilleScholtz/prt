package tar

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
)

var formats = regexp.MustCompile(`.*\.tar|.*\.tar\.gz|.*\.tar\.Z|.*\.tgz|.*\.tar\.bz2|.*\.tbz2|.*\.tar\.xz|.*\.txz|.*\.tar\.lzma|.*\.tar\.lz|.*\.zip|.*\.rpm|.*\.7z`)

func IsArchive(path string) bool {
	return formats.MatchString(path)
}

// Unpack unpacks an archive.
func Unpack(source, target string) error {
	cmd := exec.Command("bsdtar", "-p", "-o", "-C", target, "-xf", source)
	bb := new(bytes.Buffer)
	cmd.Stdout = bb

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("tar unpack %s: Something went wrong", source)
	}

	return nil
}
