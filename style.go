package terminal

import "strings"

var emptyStyle = style{}

type style struct {
	fgColor     string
	bgColor     string
	otherColors []string
	asString    string
}

func (s *style) string() string {
	if s.asString != "" || s.empty() {
		return s.asString
	}

	var styles []string
	if s.fgColor != "" {
		styles = append(styles, s.fgColor)
	}
	if s.bgColor != "" {
		styles = append(styles, s.bgColor)
	}
	styles = append(styles, s.otherColors...)
	s.asString = strings.Join(styles, " ")
	return s.asString
}

func (s *style) empty() bool {
	return s.fgColor == "" && s.bgColor == "" && len(s.otherColors) == 0
}

func (s *style) removeOther(r string) {
	// Must be a better way ..
	var removed []string

	for _, s := range s.otherColors {
		if s != r {
			removed = append(removed, s)
		}
	}
	s.otherColors = removed
}

func (s *style) color(colors []string) *style {
	if len(colors) == 0 {
		return s
	}

	if len(colors) == 1 && colors[0] == "0" {
		// Shortcut for full style reset
		return &emptyStyle
	}

	s = &style{fgColor: s.fgColor, bgColor: s.bgColor, otherColors: s.otherColors}

	if len(colors) >= 2 {
		if colors[0] == "38" && colors[1] == "5" {
			// Extended set foreground x-term color
			s.fgColor = "term-fgx" + colors[2]
			return s
		}

		// Extended set background x-term color
		if colors[0] == "48" && colors[1] == "5" {
			s.bgColor = "term-bgx" + colors[2]
			return s
		}
	}

	for _, cc := range colors {
		// If multiple colors are defined, i.e. \e[30;42m\e then loop through each
		// one, and assign it to s.fgColor or s.bgColor
		switch cc {
		case "0":
			// Reset all styles - don't use &emptyStyle here as we could end up adding colours
			// in this same action.
			s = &style{}
		case "21", "22":
			s.removeOther("term-fg1")
			s.removeOther("term-fg2")
			// Turn off italic
		case "23":
			s.removeOther("term-fg3")
			// Turn off underline
		case "24":
			s.removeOther("term-fg4")
			// Turn off crossed-out
		case "29":
			s.removeOther("term-fg9")
		case "39":
			s.fgColor = ""
		case "49":
			s.bgColor = ""
			// 30–37, then it's a foreground color
		case "30", "31", "32", "33", "34", "35", "36", "37":
			s.fgColor = "term-fg" + cc
			// 40–47, then it's a background color.
		case "40", "41", "42", "43", "44", "45", "46", "47":
			s.bgColor = "term-bg" + cc
			// 90-97 is like the regular fg color, but high intensity
		case "90", "91", "92", "93", "94", "95", "96", "97":
			s.fgColor = "term-fgi" + cc
			// 100-107 is like the regular bg color, but high intensity
		case "100", "101", "102", "103", "104", "105", "106", "107":
			s.fgColor = "term-bgi" + cc
			// 1-9 random other styles
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			s.otherColors = append(s.otherColors, "term-fg"+cc)
		}
	}
	return s
}
