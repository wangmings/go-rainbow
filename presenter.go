package rainbow

import "fmt"

type Painter struct{ Enabled bool }

type Text struct {
	value   string
	enabled bool
}

func (p Painter) Wrap(v any) Text       { return Text{value: fmt.Sprint(v), enabled: p.Enabled} }
func (t Text) String() string           { return t.value }
func (t Text) Bold() Text               { return t.Bright() }
func (t Text) Dark() Text               { return t.Faint() }
func (t Text) Strike() Text             { return t.CrossOut() }
func (t Text) FG(values ...any) Text    { return t.Color(values...) }
func (t Text) Foreground(v ...any) Text { return t.Color(v...) }
func (t Text) Colour(v ...any) Text     { return t.Color(v...) }
func (t Text) BG(values ...any) Text    { return t.Background(values...) }
func (t Text) Reset() Text              { return t.wrapCode("0") }
func (t Text) Bright() Text             { return t.wrapCode("1") }
func (t Text) Faint() Text              { return t.wrapCode("2") }
func (t Text) Italic() Text             { return t.wrapCode("3") }
func (t Text) Underline() Text          { return t.wrapCode("4") }
func (t Text) Blink() Text              { return t.wrapCode("5") }
func (t Text) Inverse() Text            { return t.wrapCode("7") }
func (t Text) Hide() Text               { return t.wrapCode("8") }
func (t Text) CrossOut() Text           { return t.wrapCode("9") }
func (t Text) Black() Text              { return t.Color("black") }
func (t Text) Red() Text                { return t.Color("red") }
func (t Text) Green() Text              { return t.Color("green") }
func (t Text) Yellow() Text             { return t.Color("yellow") }
func (t Text) Blue() Text               { return t.Color("blue") }
func (t Text) Magenta() Text            { return t.Color("magenta") }
func (t Text) Cyan() Text               { return t.Color("cyan") }
func (t Text) White() Text              { return t.Color("white") }
func (t Text) Default() Text            { return t.Color("default") }

func (t Text) Color(values ...any) Text {
	if !t.enabled {
		return t
	}
	code, err := colorCode(false, values...)
	if err != nil {
		panic(err)
	}
	return t.wrapCode(code)
}

func (t Text) Background(values ...any) Text {
	if !t.enabled {
		return t
	}
	code, err := colorCode(true, values...)
	if err != nil {
		panic(err)
	}
	return t.wrapCode(code)
}

func (t Text) wrapCode(code string) Text {
	if !t.enabled {
		return t
	}
	t.value = wrapANSI(t.value, code)
	return t
}
