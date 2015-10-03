// +build ignore

// Copyright 2015 The TCell Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This command is used to generate suitable configuration files in either
// go syntax or in JSON.  It defaults to JSON output on stdout.  If no
// term values are specified on the command line, then $TERM is used.
//
// Usage is like this:
//
// mkinfo [-go file.go] [-json file.json] [-quiet] [-nofatal] [<term>...]
//
// -go       specifiles Go output into the named file.  Use - for stdout.
// -json     specifies JSON output in the named file.  Use - for stdout
// -nofatal  indicates that errors loading definitions should not be fatal
//

package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gdamore/tcell"
)

// #include <curses.h>
// #include <term.h>
// #cgo LDFLAGS: -lcurses
//
// void noenv() {
//	use_env(FALSE);
// }
//
// char *tigetstr_good(char *name) {
//	char *r;
//	r = tigetstr(name);
//	if (r == (char *)-1) {
//		r = NULL;
//	}
//	return (r);
// }
import "C"

func tigetnum(s string) int {
	n := C.tigetnum(C.CString(s))
	return int(n)
}

func tigetflag(s string) bool {
	n := C.tigetflag(C.CString(s))
	return n != 0
}

func tigetstr(s string) string {
	// NB: If the string is invalid, we'll get back -1, which causes
	// no end of grief.  So make sure your capability strings are correct!
	cs := C.tigetstr_good(C.CString(s))
	if cs == nil {
		return ""
	}
	return C.GoString(cs)
}

// This program is used to collect data from the system's terminfo library,
// and write it into Go source code.  That is, we maintain our terminfo
// capabilities encoded in the program.  It should never need to be run by
// an end user, but developers can use this to add codes for additional
// terminal types.
func getinfo(name string) (*tcell.Terminfo, error) {
	rsn := C.int(0)
	C.noenv()
	rv, _ := C.setupterm(C.CString(name), 1, &rsn)
	if rv == C.ERR {
		switch rsn {
		case 1:
			return nil, errors.New("hardcopy terminal")
		case 0:
			return nil, errors.New("terminal definition not found")
		case -1:
			return nil, errors.New("terminfo database missing")
		default:
			return nil, errors.New("setupterm failed (other)")
		}
	}
	t := &tcell.Terminfo{}
	t.Name = name
	t.Colors = tigetnum("colors")
	t.Columns = tigetnum("cols")
	t.Lines = tigetnum("lines")
	t.Bell = tigetstr("bel")
	t.Clear = tigetstr("clear")
	t.EnterCA = tigetstr("smcup")
	t.ExitCA = tigetstr("rmcup")
	t.ShowCursor = tigetstr("cnorm")
	t.HideCursor = tigetstr("civis")
	t.AttrOff = tigetstr("sgr0")
	t.Underline = tigetstr("smul")
	t.Bold = tigetstr("bold")
	t.Blink = tigetstr("blink")
	t.Dim = tigetstr("dim")
	t.Reverse = tigetstr("rev")
	t.EnterKeypad = tigetstr("smkx")
	t.ExitKeypad = tigetstr("rmkx")
	t.SetFg = tigetstr("setaf")
	t.SetBg = tigetstr("setab")
	t.SetCursor = tigetstr("cup")
	t.CursorBack1 = tigetstr("cub1")
	t.CursorUp1 = tigetstr("cuu1")
	t.KeyF1 = tigetstr("kf1")
	t.KeyF2 = tigetstr("kf2")
	t.KeyF3 = tigetstr("kf3")
	t.KeyF4 = tigetstr("kf4")
	t.KeyF5 = tigetstr("kf5")
	t.KeyF6 = tigetstr("kf6")
	t.KeyF7 = tigetstr("kf7")
	t.KeyF8 = tigetstr("kf8")
	t.KeyF9 = tigetstr("kf9")
	t.KeyF10 = tigetstr("kf10")
	t.KeyF11 = tigetstr("kf11")
	t.KeyF12 = tigetstr("kf12")
	t.KeyF13 = tigetstr("kf13")
	t.KeyF14 = tigetstr("kf14")
	t.KeyF15 = tigetstr("kf15")
	t.KeyF16 = tigetstr("kf16")
	t.KeyF17 = tigetstr("kf17")
	t.KeyF18 = tigetstr("kf18")
	t.KeyF19 = tigetstr("kf19")
	t.KeyF20 = tigetstr("kf20")
	t.KeyF21 = tigetstr("kf21")
	t.KeyF22 = tigetstr("kf22")
	t.KeyF23 = tigetstr("kf23")
	t.KeyF24 = tigetstr("kf24")
	t.KeyF25 = tigetstr("kf25")
	t.KeyF26 = tigetstr("kf26")
	t.KeyF27 = tigetstr("kf27")
	t.KeyF28 = tigetstr("kf28")
	t.KeyF29 = tigetstr("kf29")
	t.KeyF30 = tigetstr("kf30")
	t.KeyF31 = tigetstr("kf31")
	t.KeyF32 = tigetstr("kf32")
	t.KeyF33 = tigetstr("kf33")
	t.KeyF34 = tigetstr("kf34")
	t.KeyF35 = tigetstr("kf35")
	t.KeyF36 = tigetstr("kf36")
	t.KeyF37 = tigetstr("kf37")
	t.KeyF38 = tigetstr("kf38")
	t.KeyF39 = tigetstr("kf39")
	t.KeyF40 = tigetstr("kf40")
	t.KeyF41 = tigetstr("kf41")
	t.KeyF42 = tigetstr("kf42")
	t.KeyF43 = tigetstr("kf43")
	t.KeyF44 = tigetstr("kf44")
	t.KeyF45 = tigetstr("kf45")
	t.KeyF46 = tigetstr("kf46")
	t.KeyF47 = tigetstr("kf47")
	t.KeyF48 = tigetstr("kf48")
	t.KeyF49 = tigetstr("kf49")
	t.KeyF50 = tigetstr("kf50")
	t.KeyF51 = tigetstr("kf51")
	t.KeyF52 = tigetstr("kf52")
	t.KeyF53 = tigetstr("kf53")
	t.KeyF54 = tigetstr("kf54")
	t.KeyF55 = tigetstr("kf55")
	t.KeyF56 = tigetstr("kf56")
	t.KeyF57 = tigetstr("kf57")
	t.KeyF58 = tigetstr("kf58")
	t.KeyF59 = tigetstr("kf59")
	t.KeyF60 = tigetstr("kf60")
	t.KeyF61 = tigetstr("kf61")
	t.KeyF62 = tigetstr("kf62")
	t.KeyF63 = tigetstr("kf63")
	t.KeyF64 = tigetstr("kf64")
	t.KeyInsert = tigetstr("kich1")
	t.KeyDelete = tigetstr("kdch1")
	t.KeyBackspace = tigetstr("kbs")
	t.KeyHome = tigetstr("khome")
	t.KeyEnd = tigetstr("kend")
	t.KeyUp = tigetstr("kcuu1")
	t.KeyDown = tigetstr("kcud1")
	t.KeyRight = tigetstr("kcuf1")
	t.KeyLeft = tigetstr("kcub1")
	t.KeyPgDn = tigetstr("knp")
	t.KeyPgUp = tigetstr("kpp")
	t.KeyBacktab = tigetstr("kcbt")
	t.KeyExit = tigetstr("kext")
	t.KeyCancel = tigetstr("kcan")
	t.KeyPrint = tigetstr("kprt")
	t.KeyHelp = tigetstr("khlp")
	t.KeyClear = tigetstr("kclr")
	t.Mouse = tigetstr("kmous")
	// If the kmous entry is present, then we need to record the
	// the codes to enter and exit mouse mode.  Sadly, this is not
	// part of the terminfo databases anywhere that I've found, but
	// is an extension.  The escape codes are documented in the XTerm
	// manual, and all terminals that have kmous are expected to
	// use these same codes.
	if t.Mouse != "" {
		// we anticipate that all xterm mouse tracking compatible
		// terminals understand mouse tracking (1000), but we hope
		// that those that don't understand any-event tracking (1003)
		// will at least ignore it.  Likewise we hope that terminals
		// that don't understand SGR reporting (1006) just ignore it.
		t.EnterMouse = "\x1b[?1000h\x1b[?1003h\x1b[?1006h"
		t.ExitMouse = "\x1b[?1006l\x1b[?1003l\x1b[?1000l"
	}
	// We only support colors in ANSI 8 or 256 color mode.
	if t.Colors < 8 || t.SetFg == "" {
		t.Colors = 0
	}
	if t.SetCursor == "" {
		return nil, errors.New("terminal not cursor addressable")
	}

	// For padding, we lookup the pad char.  If that isn't present,
	// and npc is *not* set, then we assume a null byte.
	t.PadChar = tigetstr("pad")
	if t.PadChar == "" {
		if !tigetflag("npc") {
			t.PadChar = "\u0000"
		}
	}

	return t, nil
}

func dotGoAddInt(w io.Writer, n string, i int) {
	if i == 0 {
		// initialized to 0, ignore
		return
	}
	fmt.Fprintf(w, "		%-13s %d,\n", n+":", i)
}
func dotGoAddStr(w io.Writer, n string, s string) {
	if s == "" {
		return
	}
	fmt.Fprintf(w, "		%-13s %q,\n", n+":", s)
}

func dotGoAddArr(w io.Writer, n string, a []string) {
	if len(a) == 0 {
		return
	}
	fmt.Fprintf(w, "		%-13s []string{ ", n+":")
	did := false
	for _, b := range a {
		if did {
			fmt.Fprint(w, ", ")
		}
		did = true
		fmt.Fprintf(w, "%q", b)
	}
	fmt.Fprintln(w, " },")
}

func dotGoHeader(w io.Writer) {
	fmt.Fprintf(w, "// Generated by %s (%s/%s) on %s.\n",
		os.Args[0],
		runtime.GOOS, runtime.GOARCH,
		time.Now().Format(time.UnixDate))
	fmt.Fprintln(w, "// DO NOT HAND-EDIT")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "package tcell")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "func init() {")
}

func dotGoTrailer(w io.Writer) {
	fmt.Fprintln(w, "}")
}

func dotGoInfo(w io.Writer, t *tcell.Terminfo) {
	fmt.Fprintln(w, "	AddTerminfo(&Terminfo{")
	dotGoAddStr(w, "Name", t.Name)
	dotGoAddArr(w, "Aliases", t.Aliases)
	dotGoAddInt(w, "Columns", t.Columns)
	dotGoAddInt(w, "Lines", t.Lines)
	dotGoAddInt(w, "Colors", t.Colors)
	dotGoAddStr(w, "Bell", t.Bell)
	dotGoAddStr(w, "Clear", t.Clear)
	dotGoAddStr(w, "EnterCA", t.EnterCA)
	dotGoAddStr(w, "ExitCA", t.ExitCA)
	dotGoAddStr(w, "ShowCursor", t.ShowCursor)
	dotGoAddStr(w, "HideCursor", t.HideCursor)
	dotGoAddStr(w, "AttrOff", t.AttrOff)
	dotGoAddStr(w, "Underline", t.Underline)
	dotGoAddStr(w, "Bold", t.Bold)
	dotGoAddStr(w, "Dim", t.Dim)
	dotGoAddStr(w, "Blink", t.Blink)
	dotGoAddStr(w, "Reverse", t.Reverse)
	dotGoAddStr(w, "EnterKeypad", t.EnterKeypad)
	dotGoAddStr(w, "ExitKeypad", t.ExitKeypad)
	dotGoAddStr(w, "SetFg", t.SetFg)
	dotGoAddStr(w, "SetBg", t.SetBg)
	dotGoAddStr(w, "PadChar", t.PadChar)
	dotGoAddStr(w, "Mouse", t.Mouse)
	dotGoAddStr(w, "EnterMouse", t.EnterMouse)
	dotGoAddStr(w, "ExitMouse", t.ExitMouse)
	dotGoAddStr(w, "SetCursor", t.SetCursor)
	dotGoAddStr(w, "CursorBack1", t.CursorBack1)
	dotGoAddStr(w, "CursorUp1", t.CursorUp1)
	dotGoAddStr(w, "KeyUp", t.KeyUp)
	dotGoAddStr(w, "KeyDown", t.KeyDown)
	dotGoAddStr(w, "KeyRight", t.KeyRight)
	dotGoAddStr(w, "KeyLeft", t.KeyLeft)
	dotGoAddStr(w, "KeyInsert", t.KeyInsert)
	dotGoAddStr(w, "KeyDelete", t.KeyDelete)
	dotGoAddStr(w, "KeyBackspace", t.KeyBackspace)
	dotGoAddStr(w, "KeyHome", t.KeyHome)
	dotGoAddStr(w, "KeyEnd", t.KeyEnd)
	dotGoAddStr(w, "KeyPgUp", t.KeyPgUp)
	dotGoAddStr(w, "KeyPgDn", t.KeyPgDn)
	dotGoAddStr(w, "KeyF1", t.KeyF1)
	dotGoAddStr(w, "KeyF2", t.KeyF2)
	dotGoAddStr(w, "KeyF3", t.KeyF3)
	dotGoAddStr(w, "KeyF4", t.KeyF4)
	dotGoAddStr(w, "KeyF5", t.KeyF5)
	dotGoAddStr(w, "KeyF6", t.KeyF6)
	dotGoAddStr(w, "KeyF7", t.KeyF7)
	dotGoAddStr(w, "KeyF8", t.KeyF8)
	dotGoAddStr(w, "KeyF9", t.KeyF9)
	dotGoAddStr(w, "KeyF10", t.KeyF10)
	dotGoAddStr(w, "KeyF11", t.KeyF11)
	dotGoAddStr(w, "KeyF12", t.KeyF12)
	dotGoAddStr(w, "KeyF13", t.KeyF13)
	dotGoAddStr(w, "KeyF14", t.KeyF14)
	dotGoAddStr(w, "KeyF15", t.KeyF15)
	dotGoAddStr(w, "KeyF16", t.KeyF16)
	dotGoAddStr(w, "KeyF17", t.KeyF17)
	dotGoAddStr(w, "KeyF18", t.KeyF18)
	dotGoAddStr(w, "KeyF19", t.KeyF19)
	dotGoAddStr(w, "KeyF20", t.KeyF20)
	dotGoAddStr(w, "KeyF21", t.KeyF21)
	dotGoAddStr(w, "KeyF22", t.KeyF22)
	dotGoAddStr(w, "KeyF23", t.KeyF23)
	dotGoAddStr(w, "KeyF24", t.KeyF24)
	dotGoAddStr(w, "KeyF25", t.KeyF25)
	dotGoAddStr(w, "KeyF26", t.KeyF26)
	dotGoAddStr(w, "KeyF27", t.KeyF27)
	dotGoAddStr(w, "KeyF28", t.KeyF28)
	dotGoAddStr(w, "KeyF29", t.KeyF29)
	dotGoAddStr(w, "KeyF30", t.KeyF30)
	dotGoAddStr(w, "KeyF31", t.KeyF31)
	dotGoAddStr(w, "KeyF32", t.KeyF32)
	dotGoAddStr(w, "KeyF33", t.KeyF33)
	dotGoAddStr(w, "KeyF34", t.KeyF34)
	dotGoAddStr(w, "KeyF35", t.KeyF35)
	dotGoAddStr(w, "KeyF36", t.KeyF36)
	dotGoAddStr(w, "KeyF37", t.KeyF37)
	dotGoAddStr(w, "KeyF38", t.KeyF38)
	dotGoAddStr(w, "KeyF39", t.KeyF39)
	dotGoAddStr(w, "KeyF40", t.KeyF40)
	dotGoAddStr(w, "KeyF41", t.KeyF41)
	dotGoAddStr(w, "KeyF42", t.KeyF42)
	dotGoAddStr(w, "KeyF43", t.KeyF43)
	dotGoAddStr(w, "KeyF44", t.KeyF44)
	dotGoAddStr(w, "KeyF45", t.KeyF45)
	dotGoAddStr(w, "KeyF46", t.KeyF46)
	dotGoAddStr(w, "KeyF47", t.KeyF47)
	dotGoAddStr(w, "KeyF48", t.KeyF48)
	dotGoAddStr(w, "KeyF49", t.KeyF49)
	dotGoAddStr(w, "KeyF50", t.KeyF50)
	dotGoAddStr(w, "KeyF51", t.KeyF51)
	dotGoAddStr(w, "KeyF52", t.KeyF52)
	dotGoAddStr(w, "KeyF53", t.KeyF53)
	dotGoAddStr(w, "KeyF54", t.KeyF54)
	dotGoAddStr(w, "KeyF55", t.KeyF55)
	dotGoAddStr(w, "KeyF56", t.KeyF56)
	dotGoAddStr(w, "KeyF57", t.KeyF57)
	dotGoAddStr(w, "KeyF58", t.KeyF58)
	dotGoAddStr(w, "KeyF59", t.KeyF59)
	dotGoAddStr(w, "KeyF60", t.KeyF60)
	dotGoAddStr(w, "KeyF61", t.KeyF61)
	dotGoAddStr(w, "KeyF62", t.KeyF62)
	dotGoAddStr(w, "KeyF63", t.KeyF63)
	dotGoAddStr(w, "KeyF64", t.KeyF64)
	dotGoAddStr(w, "KeyCancel", t.KeyCancel)
	dotGoAddStr(w, "KeyPrint", t.KeyPrint)
	dotGoAddStr(w, "KeyExit", t.KeyExit)
	dotGoAddStr(w, "KeyHelp", t.KeyHelp)
	dotGoAddStr(w, "KeyClear", t.KeyClear)
	dotGoAddStr(w, "KeyBacktab", t.KeyBacktab)
	fmt.Fprintln(w, "	})")
}

func main() {
	gofile := ""
	jsonfile := ""
	nofatal := false
	quiet := false

	flag.StringVar(&gofile, "go", "", "generate go source in named file")
	flag.StringVar(&jsonfile, "json", "", "generate json in named file")
	flag.BoolVar(&nofatal, "nofatal", false, "errors are not fatal")
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")
	flag.Parse()
	var e error
	js := []byte{}

	args := flag.Args()
	if len(args) == 0 {
		args = []string{os.Getenv("TERM")}
	}

	tdata := make(map[string]*tcell.Terminfo)
	adata := make(map[string]string)
	for _, term := range args {
		if arr := strings.SplitN(term, "=", 2); len(arr) == 2 {
			adata[arr[0]] = arr[1]
		} else if t, e := getinfo(term); e != nil {
			if !quiet {
				fmt.Fprintf(os.Stderr,
					"Failed loading %s: %v\n", term, e)
			}
			if !nofatal {
				os.Exit(1)
			}
		} else {
			tdata[t.Name] = t
		}
	}
	for alias, canon := range adata {
		if t, ok := tdata[canon]; ok {
			t.Aliases = append(t.Aliases, alias)
		} else {
			if !quiet {
				fmt.Fprintf(os.Stderr,
					"Alias %s missing canonical %s\n",
					alias, canon)
			}
			if !nofatal {
				os.Exit(1)
			}
		}
	}

	if gofile != "" {
		w := os.Stdout
		if gofile != "-" {
			if w, e = os.Create(gofile); e != nil {
				fmt.Fprintf(os.Stderr, "Failed: %v", e)
				os.Exit(1)
			}
		}
		dotGoHeader(w)
		for _, term := range args {
			if t := tdata[term]; t != nil {
				dotGoInfo(w, t)
			}
		}
		dotGoTrailer(w)
		if w != os.Stdout {
			w.Close()
		}
	} else if jsonfile != "" {
		w := os.Stdout
		if jsonfile != "-" {
			if w, e = os.Create(jsonfile); e != nil {
				fmt.Fprintf(os.Stderr, "Failed: %v", e)
			}
		}
		for _, term := range args {
			if t := tdata[term]; t != nil {
				js, e = json.Marshal(t)
				fmt.Fprintln(w, string(js))
			}
			// arguably if there is more than one term, this
			// should be a javascript array, but that's not how
			// we load it.  We marshal objects one at a time from
			// the file.
		}
		if e != nil {
			fmt.Fprintf(os.Stderr, "Failed: %v", e)
			os.Exit(1)
		}
	} else {
		for _, term := range args {
			if t := tdata[term]; t != nil {
				js, e := json.Marshal(tdata[term])
				if e != nil {
					fmt.Fprintf(os.Stderr, "Failed: %v", e)
					os.Exit(1)
				}
				fmt.Fprintln(os.Stdout, string(js))
			}
		}
	}
}