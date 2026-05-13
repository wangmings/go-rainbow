package rainbow

import "testing"

func TestCoreRainbowBehavior(t *testing.T) {
	Enabled(true)
	cases := map[string]string{
		Wrap("hello").Foreground(5).String():                 "\x1b[35mhello\x1b[0m",
		Wrap("hello").Foreground("red").String():             "\x1b[31mhello\x1b[0m",
		Wrap("hello").Colour("red").String():                 "\x1b[31mhello\x1b[0m",
		Wrap("hello").Color("midnightblue").String():         "\x1b[38;5;18mhello\x1b[0m",
		Wrap("hello").Midnightblue().String():                "\x1b[38;5;18mhello\x1b[0m",
		Wrap("hello").Foreground(255, 128, 64).String():      "\x1b[38;5;215mhello\x1b[0m",
		Wrap("hello").Background("#ff8040").String():         "\x1b[48;5;215mhello\x1b[0m",
		Wrap("hello").Underline().Bright().String():          "\x1b[4m\x1b[1mhello\x1b[0m",
		Wrap("hello").CrossOut().Strike().String():           "\x1b[9m\x1b[9mhello\x1b[0m",
		Wrap("hello").Red().Bright().Blue().Blink().String(): "\x1b[31m\x1b[1m\x1b[34m\x1b[5mhello\x1b[0m",
		Uncolor("\x1b[35mhello\x1b[0mm"):                     "hellom",
	}
	for got, want := range cases {
		if got != want {
			t.Fatalf("got %q want %q", got, want)
		}
	}
}

func TestDisableAndIndependentPainter(t *testing.T) {
	Enabled(false)
	if got := Wrap("hello").Red().Bright().String(); got != "hello" {
		t.Fatalf("disabled global painter returned %q", got)
	}
	if got := Wrap("hello").Color("not-a-real-color").Background("still-not-real").String(); got != "hello" {
		t.Fatalf("disabled painter should not parse colors, got %q", got)
	}
	p := Painter{Enabled: true}
	if got := p.Wrap("hello").Green().String(); got != "\x1b[32mhello\x1b[0m" {
		t.Fatalf("independent painter returned %q", got)
	}
	Enabled(true)
}

func TestNewPainterInheritsGlobalState(t *testing.T) {
	Enabled(false)
	disabled := New()
	if disabled.Enabled {
		t.Fatal("new painter should inherit disabled global state")
	}

	Enabled(true)
	enabled := New()
	if !enabled.Enabled {
		t.Fatal("new painter should inherit enabled global state")
	}

	enabled.Enabled = false
	if !Global.Enabled {
		t.Fatal("changing painter state should not change global state")
	}
}

func TestUncolorOnlyRemovesSGRSequences(t *testing.T) {
	if got := Uncolor("\x1b[35mhello\x1b[0mm"); got != "hellom" {
		t.Fatalf("uncolor got %q", got)
	}
	if got := Uncolor("\x1b[1Thello"); got != "\x1b[1Thello" {
		t.Fatalf("non-SGR escape should stay intact, got %q", got)
	}
}

func TestX11PaletteCoverage(t *testing.T) {
	if got := len(x11Colors); got != 145 {
		t.Fatalf("x11 palette has %d colors, want 145", got)
	}
}

func TestDefaultEnabledMatchesEnvironmentOverrides(t *testing.T) {
	t.Setenv("TERM", "dumb")
	t.Setenv("CLICOLOR_FORCE", "")
	if defaultEnabled() {
		t.Fatal("TERM=dumb should disable colors")
	}

	t.Setenv("CLICOLOR_FORCE", "1")
	if !defaultEnabled() {
		t.Fatal("CLICOLOR_FORCE=1 should enable colors")
	}
}
