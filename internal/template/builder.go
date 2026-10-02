package template

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type CommandFunc func(source, output string) []string

/*
func build(dir string, output string) (string, error) {

	buildID := randHex(8)
	mutex := randHex(16)

	ldflags := fmt.Sprintf(
		`-s -w -X main.BuildID=%s -X main.MutexName=%s`,
		buildID, mutex,
	)

	cmd := exec.Command("go", "build",
		"-ldflags", ldflags,
		"-trimpath",
		"-o", output,
		dir,
	)

	stdout, err := cmd.CombinedOutput()
	return string(stdout), err
}
*/

func randHex(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func BuildGoImplant(source, output string) []string {
	buildID := randHex(8)
	mutex := randHex(16)

	ldflags := fmt.Sprintf(
		`-s -w -X main.BuildID=%s -X main.MutexName=%s`,
		buildID, mutex,
	)

	return []string{"go", "build",
		"-ldflags", ldflags,
		"-trimpath",
		"-o", output,
		source}
}
