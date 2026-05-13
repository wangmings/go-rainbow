package rainbow

import "os"

var Global = Painter{Enabled: defaultEnabled()}

func Wrap(v any) Text         { return Global.Wrap(v) }
func Enabled(v bool)          { Global.Enabled = v }
func Uncolor(s string) string { return stripANSI(s) }
func New() Painter            { return Painter{Enabled: Global.Enabled} }

func defaultEnabled() bool {
	if os.Getenv("CLICOLOR_FORCE") == "1" {
		return true
	}
	if os.Getenv("TERM") == "dumb" {
		return false
	}
	return isTTY(os.Stdout) && isTTY(os.Stderr)
}

func isTTY(file *os.File) bool {
	info, err := file.Stat()
	return err == nil && info.Mode()&os.ModeCharDevice != 0
}
