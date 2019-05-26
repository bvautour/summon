package hue

import (
	"fmt"
	"strconv"
	"strings"
)

// Hue ...
type Hue struct {
	attr []Attr
}

// Attr ...
type Attr int

const esc = "\x1b"

// RESET ...
const RESET = 0

// Foreground
const (
	B Attr = iota + 30 // black
	R                  // red
	G                  // green
	Y                  // yellow
	U                  // blue
	M                  // magenta
	C                  // cyan
	W                  // white
)

// New ...
func New(colors ...Attr) *Hue {
	h := &Hue{attr: make([]Attr, 0)}
	h.attr = append(h.attr, colors...)
	return h
}

func (h *Hue) seq(s string) string {
	return fmt.Sprintf("%s[%sm", esc, s)
}

func (h *Hue) wrap(s string) string {
	return h.colorize() + s + h.reset()
}

func (h *Hue) colorize() string {
	format := make([]string, len(h.attr))
	for i, v := range h.attr {
		format[i] = strconv.Itoa(int(v))
	}
	return h.seq(strings.Join(format, ";"))
}

func (h *Hue) reset() string {
	return h.seq(string(RESET))
}

// Black ...
func Black(s string) string {
	h := New(B)
	return h.wrap(s)
}

// Red ...
func Red(s string) string {
	h := New(R)
	return h.wrap(s)
}

// Green ...
func Green(s string) string {
	h := New(G)
	return h.wrap(s)
}

// Yellow ...
func Yellow(s string) string {
	h := New(Y)
	return h.wrap(s)
}

// Blue ...
func Blue(s string) string {
	h := New(U)
	return h.wrap(s)
}

// Magenta ...
func Magenta(s string) string {
	h := New(M)
	return h.wrap(s)
}

// Cyan ...
func Cyan(s string) string {
	h := New(C)
	return h.wrap(s)
}

// White ...
func White(s string) string {
	h := New(W)
	return h.wrap(s)
}
