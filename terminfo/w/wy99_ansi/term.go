// Generated automatically.  DO NOT HAND-EDIT.

package wy99_ansi

import "github.com/tinywolf3/tcell/v2/terminfo"

func init() {

	// Wyse WY-99GT in ansi mode (int'l PC keyboard)
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:         "wy99-ansi",
		Columns:      80,
		Lines:        25,
		Bell:         "\a",
		Clear:        "\x1b[H\x1b[J$<200>",
		ShowCursor:   "\x1b[34h\x1b[?25h",
		HideCursor:   "\x1b[?25l",
		AttrOff:      "\x1b[m\x0f\x1b[\"q",
		Underline:    "\x1b[4m",
		Bold:         "\x1b[1m",
		Dim:          "\x1b[2m",
		Blink:        "\x1b[5m",
		Reverse:      "\x1b[7m",
		EnterKeypad:  "\x1b[?1h",
		ExitKeypad:   "\x1b[?1l",
		PadChar:      "\x00",
		AltChars:     "``aaffggjjkkllmmnnooqqssttuuvvwwxx{{||}}~~",
		EnterAcs:     "\x0e",
		ExitAcs:      "\x0f",
		EnableAcs:    "\x1b)0",
		SetCursor:    "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:  "\b$<1>",
		CursorUp1:    "\x1bM",
		KeyUp:        "\x1bOA",
		KeyDown:      "\x1bOB",
		KeyRight:     "\x1bOC",
		KeyLeft:      "\x1bOD",
		KeyBackspace: "\b",
		KeyF1:        "\x1bOP",
		KeyF2:        "\x1bOQ",
		KeyF3:        "\x1bOR",
		KeyF4:        "\x1bOS",
		KeyF5:        "\x1b[M",
		KeyF6:        "\x1b[17~",
		KeyF7:        "\x1b[18~",
		KeyF8:        "\x1b[19~",
		KeyF9:        "\x1b[20~",
		KeyF10:       "\x1b[21~",
		KeyF11:       "\x1b[23~",
		KeyF12:       "\x1b[24~",
		KeyF17:       "\x1b[K",
		KeyF18:       "\x1b[31~",
		KeyF19:       "\x1b[32~",
		KeyF20:       "\x1b[33~",
		KeyF21:       "\x1b[34~",
		KeyF22:       "\x1b[35~",
		KeyF23:       "\x1b[1~",
		KeyF24:       "\x1b[2~",
		KeyBacktab:   "\x1b[z",
		AutoMargin:   true,
	})

	// Wyse WY-99GT in ansi mode (US PC keyboard)
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:         "wy99a-ansi",
		Columns:      80,
		Lines:        25,
		Bell:         "\a",
		Clear:        "\x1b[H\x1b[J$<200>",
		ShowCursor:   "\x1b[34h\x1b[?25h",
		HideCursor:   "\x1b[?25l",
		AttrOff:      "\x1b[m\x0f\x1b[\"q",
		Underline:    "\x1b[4m",
		Bold:         "\x1b[1m",
		Dim:          "\x1b[2m",
		Blink:        "\x1b[5m",
		Reverse:      "\x1b[7m",
		EnterKeypad:  "\x1b[?1h",
		ExitKeypad:   "\x1b[?1l",
		PadChar:      "\x00",
		AltChars:     "``aaffggjjkkllmmnnooqqssttuuvvwwxx{{||}}~~",
		EnterAcs:     "\x0e",
		ExitAcs:      "\x0f",
		EnableAcs:    "\x1b)0",
		SetCursor:    "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:  "\b$<1>",
		CursorUp1:    "\x1bM",
		KeyUp:        "\x1bOA",
		KeyDown:      "\x1bOB",
		KeyRight:     "\x1bOC",
		KeyLeft:      "\x1bOD",
		KeyBackspace: "\b",
		KeyF1:        "\x1bOP",
		KeyF2:        "\x1bOQ",
		KeyF3:        "\x1bOR",
		KeyF4:        "\x1bOS",
		KeyF5:        "\x1b[M",
		KeyF6:        "\x1b[17~",
		KeyF7:        "\x1b[18~",
		KeyF8:        "\x1b[19~",
		KeyF9:        "\x1b[20~",
		KeyF10:       "\x1b[21~",
		KeyF11:       "\x1b[23~",
		KeyF12:       "\x1b[24~",
		KeyF17:       "\x1b[K",
		KeyF18:       "\x1b[31~",
		KeyF19:       "\x1b[32~",
		KeyF20:       "\x1b[33~",
		KeyF21:       "\x1b[34~",
		KeyF22:       "\x1b[35~",
		KeyF23:       "\x1b[1~",
		KeyF24:       "\x1b[2~",
		KeyBacktab:   "\x1b[z",
		AutoMargin:   true,
	})
}
