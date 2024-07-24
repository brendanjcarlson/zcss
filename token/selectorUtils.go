package token

import (
	"strings"
	"unicode"
)

func IsCssSelectorStart(char rune) bool {
	return unicode.IsLetter(char) ||
		char == '*' ||
		char == '.' ||
		char == '#' ||
		char == '[' ||
		char == ':'
}

func IsCssSelector(value string) (subkind Subkind, ok bool) {
	if subkind, ok = IsUniversalSelector(value); ok {
		return
	}
	if subkind, ok = IsClassSelector(value); ok {
		return
	}
	if subkind, ok = IsIdSelector(value); ok {
		return
	}
	if subkind, ok = IsCssElementSelector(value); ok {
		return
	}
	if subkind, ok = IsCssPseudoClassSelector(value); ok {
		return
	}
	if subkind, ok = IsCssPseudoElementSelector(value); ok {
		return
	}
	if subkind, ok = IsInheritSelector(value); ok {
		return
	}
	return COMBINATOR_SELECTOR, true
}

// See https://developer.mozilla.org/en-US/docs/Web/HTML/Element
var CssElementSelectors = map[string]Subkind{
	// Main root
	"html": ELEMENT_SELECTOR,
	// Document metadata
	"base":  ELEMENT_SELECTOR,
	"head":  ELEMENT_SELECTOR,
	"link":  ELEMENT_SELECTOR,
	"meta":  ELEMENT_SELECTOR,
	"style": ELEMENT_SELECTOR,
	"title": ELEMENT_SELECTOR,
	// Sectioning root
	"body": ELEMENT_SELECTOR,
	// Content sectioning
	"address": ELEMENT_SELECTOR,
	"article": ELEMENT_SELECTOR,
	"aside":   ELEMENT_SELECTOR,
	"footer":  ELEMENT_SELECTOR,
	"header":  ELEMENT_SELECTOR,
	"h1":      ELEMENT_SELECTOR,
	"h2":      ELEMENT_SELECTOR,
	"h3":      ELEMENT_SELECTOR,
	"h4":      ELEMENT_SELECTOR,
	"h5":      ELEMENT_SELECTOR,
	"h6":      ELEMENT_SELECTOR,
	"hgroup":  ELEMENT_SELECTOR,
	"main":    ELEMENT_SELECTOR,
	"nav":     ELEMENT_SELECTOR,
	"section": ELEMENT_SELECTOR,
	"search":  ELEMENT_SELECTOR,
	// Text content
	"blockquote": ELEMENT_SELECTOR,
	"dd":         ELEMENT_SELECTOR,
	"div":        ELEMENT_SELECTOR,
	"dl":         ELEMENT_SELECTOR,
	"dt":         ELEMENT_SELECTOR,
	"figcaption": ELEMENT_SELECTOR,
	"hr":         ELEMENT_SELECTOR,
	"li":         ELEMENT_SELECTOR,
	"menu":       ELEMENT_SELECTOR,
	"ol":         ELEMENT_SELECTOR,
	"p":          ELEMENT_SELECTOR,
	"pre":        ELEMENT_SELECTOR,
	"ul":         ELEMENT_SELECTOR,
	// Inline text semantics
	"a":      ELEMENT_SELECTOR,
	"abbr":   ELEMENT_SELECTOR,
	"b":      ELEMENT_SELECTOR,
	"bdi":    ELEMENT_SELECTOR,
	"bdo":    ELEMENT_SELECTOR,
	"br":     ELEMENT_SELECTOR,
	"cite":   ELEMENT_SELECTOR,
	"code":   ELEMENT_SELECTOR,
	"data":   ELEMENT_SELECTOR,
	"dfn":    ELEMENT_SELECTOR,
	"em":     ELEMENT_SELECTOR,
	"i":      ELEMENT_SELECTOR,
	"kbd":    ELEMENT_SELECTOR,
	"mark":   ELEMENT_SELECTOR,
	"q":      ELEMENT_SELECTOR,
	"rp":     ELEMENT_SELECTOR,
	"rt":     ELEMENT_SELECTOR,
	"ruby":   ELEMENT_SELECTOR,
	"s":      ELEMENT_SELECTOR,
	"samp":   ELEMENT_SELECTOR,
	"small":  ELEMENT_SELECTOR,
	"span":   ELEMENT_SELECTOR,
	"strong": ELEMENT_SELECTOR,
	"sub":    ELEMENT_SELECTOR,
	"sup":    ELEMENT_SELECTOR,
	"time":   ELEMENT_SELECTOR,
	"u":      ELEMENT_SELECTOR,
	"var":    ELEMENT_SELECTOR,
	"wbr":    ELEMENT_SELECTOR,
	// Image and multimedia
	"area":  ELEMENT_SELECTOR,
	"audio": ELEMENT_SELECTOR,
	"img":   ELEMENT_SELECTOR,
	"map":   ELEMENT_SELECTOR,
	"track": ELEMENT_SELECTOR,
	"video": ELEMENT_SELECTOR,
	// Embedded content
	"embed":   ELEMENT_SELECTOR,
	"iframe":  ELEMENT_SELECTOR,
	"object":  ELEMENT_SELECTOR,
	"picture": ELEMENT_SELECTOR,
	"portal":  ELEMENT_SELECTOR,
	"source":  ELEMENT_SELECTOR,
	// SVG and MathML
	"svg":  ELEMENT_SELECTOR,
	"math": ELEMENT_SELECTOR,
	// Scripting
	"canvas":   ELEMENT_SELECTOR,
	"noscript": ELEMENT_SELECTOR,
	"script":   ELEMENT_SELECTOR,
	// Demarcating edits
	"del": ELEMENT_SELECTOR,
	"ins": ELEMENT_SELECTOR,
	// Table content
	"caption":  ELEMENT_SELECTOR,
	"col":      ELEMENT_SELECTOR,
	"colgroup": ELEMENT_SELECTOR,
	"table":    ELEMENT_SELECTOR,
	"tbody":    ELEMENT_SELECTOR,
	"td":       ELEMENT_SELECTOR,
	"tfoot":    ELEMENT_SELECTOR,
	"th":       ELEMENT_SELECTOR,
	"thead":    ELEMENT_SELECTOR,
	"tr":       ELEMENT_SELECTOR,
	// Forms
	"button":   ELEMENT_SELECTOR,
	"datalist": ELEMENT_SELECTOR,
	"fieldset": ELEMENT_SELECTOR,
	"form":     ELEMENT_SELECTOR,
	"input":    ELEMENT_SELECTOR,
	"label":    ELEMENT_SELECTOR,
	"legend":   ELEMENT_SELECTOR,
	"meter":    ELEMENT_SELECTOR,
	"optgroup": ELEMENT_SELECTOR,
	"option":   ELEMENT_SELECTOR,
	"output":   ELEMENT_SELECTOR,
	"progress": ELEMENT_SELECTOR,
	"select":   ELEMENT_SELECTOR,
	"textarea": ELEMENT_SELECTOR,
	// Interactive elements
	"details": ELEMENT_SELECTOR,
	"dialog":  ELEMENT_SELECTOR,
	"summary": ELEMENT_SELECTOR,
	// Web Components
	"slot":     ELEMENT_SELECTOR,
	"template": ELEMENT_SELECTOR,
	// Obsolete and deprecated elements
	"acronym":   ELEMENT_SELECTOR,
	"big":       ELEMENT_SELECTOR,
	"center":    ELEMENT_SELECTOR,
	"content":   ELEMENT_SELECTOR,
	"dir":       ELEMENT_SELECTOR,
	"font":      ELEMENT_SELECTOR,
	"frame":     ELEMENT_SELECTOR,
	"frameset":  ELEMENT_SELECTOR,
	"image":     ELEMENT_SELECTOR,
	"marquee":   ELEMENT_SELECTOR,
	"menuitem":  ELEMENT_SELECTOR,
	"nobr":      ELEMENT_SELECTOR,
	"noembed":   ELEMENT_SELECTOR,
	"noframes":  ELEMENT_SELECTOR,
	"param":     ELEMENT_SELECTOR,
	"plaintext": ELEMENT_SELECTOR,
	"rb":        ELEMENT_SELECTOR,
	"rtc":       ELEMENT_SELECTOR,
	"shadow":    ELEMENT_SELECTOR,
	"strike":    ELEMENT_SELECTOR,
	"tt":        ELEMENT_SELECTOR,
	"xmp":       ELEMENT_SELECTOR,
}

func IsCssElementSelector(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssElementSelectors[value]
	return
}

var CssPseudoClassSelectors = map[string]Subkind{
	// Element display
	":fullscreen":         PSEUDO_CLASS_SELECTOR,
	":modal":              PSEUDO_CLASS_SELECTOR,
	":picture-in-picture": PSEUDO_CLASS_SELECTOR,
	// Input
	":autofill":          PSEUDO_CLASS_SELECTOR,
	":enabled":           PSEUDO_CLASS_SELECTOR,
	":disabled":          PSEUDO_CLASS_SELECTOR,
	":read-write":        PSEUDO_CLASS_SELECTOR,
	":placeholder-shown": PSEUDO_CLASS_SELECTOR,
	":default":           PSEUDO_CLASS_SELECTOR,
	":checked":           PSEUDO_CLASS_SELECTOR,
	":indeterminate":     PSEUDO_CLASS_SELECTOR,
	":blank":             PSEUDO_CLASS_SELECTOR,
	":valid":             PSEUDO_CLASS_SELECTOR,
	":invalid":           PSEUDO_CLASS_SELECTOR,
	":in-range":          PSEUDO_CLASS_SELECTOR,
	":out-of-range":      PSEUDO_CLASS_SELECTOR,
	":required":          PSEUDO_CLASS_SELECTOR,
	":optional":          PSEUDO_CLASS_SELECTOR,
	":user-valid":        PSEUDO_CLASS_SELECTOR,
	":user-invalid":      PSEUDO_CLASS_SELECTOR,
	// Linguistic
	":dir":  PSEUDO_CLASS_SELECTOR,
	":lang": PSEUDO_CLASS_SELECTOR,
	// Location
	":any-link":      PSEUDO_CLASS_SELECTOR,
	":link":          PSEUDO_CLASS_SELECTOR,
	":visited":       PSEUDO_CLASS_SELECTOR,
	":local-link":    PSEUDO_CLASS_SELECTOR,
	":target":        PSEUDO_CLASS_SELECTOR,
	":target-within": PSEUDO_CLASS_SELECTOR,
	":scope":         PSEUDO_CLASS_SELECTOR,
	// Resource state
	":playing": PSEUDO_CLASS_SELECTOR,
	":paused":  PSEUDO_CLASS_SELECTOR,
	// Time dimensional
	":current": PSEUDO_CLASS_SELECTOR,
	":past":    PSEUDO_CLASS_SELECTOR,
	":future":  PSEUDO_CLASS_SELECTOR,
	// Tree-structural
	":root":             PSEUDO_CLASS_SELECTOR,
	":empty":            PSEUDO_CLASS_SELECTOR,
	":nth-child":        PSEUDO_CLASS_SELECTOR,
	":nth-last-child":   PSEUDO_CLASS_SELECTOR,
	":first-child":      PSEUDO_CLASS_SELECTOR,
	":last-child":       PSEUDO_CLASS_SELECTOR,
	":only-child":       PSEUDO_CLASS_SELECTOR,
	":nth-of-type":      PSEUDO_CLASS_SELECTOR,
	":nth-last-of-type": PSEUDO_CLASS_SELECTOR,
	":first-of-type":    PSEUDO_CLASS_SELECTOR,
	":last-of-type":     PSEUDO_CLASS_SELECTOR,
	":only-of-type":     PSEUDO_CLASS_SELECTOR,
	// User action
	":hover":         PSEUDO_CLASS_SELECTOR,
	":active":        PSEUDO_CLASS_SELECTOR,
	":focus":         PSEUDO_CLASS_SELECTOR,
	":focus-visible": PSEUDO_CLASS_SELECTOR,
	":focus-within":  PSEUDO_CLASS_SELECTOR,
	// Functional
	":is":    PSEUDO_CLASS_SELECTOR,
	":not":   PSEUDO_CLASS_SELECTOR,
	":where": PSEUDO_CLASS_SELECTOR,
	":has":   PSEUDO_CLASS_SELECTOR,
}

func IsCssPseudoClassSelector(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssPseudoClassSelectors[value]
	return
}

var CssPseudoElementSelectors = map[string]Subkind{
	"::after":  PSEUDO_ELEMENT_SELECTOR,
	"::before": PSEUDO_ELEMENT_SELECTOR,
}

func IsCssPseudoElementSelector(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssPseudoClassSelectors[value]
	return
}

func IsClassSelector(literal string) (subkind Subkind, ok bool) {
	if len(literal) == 0 {
		return SUBKIND_NONE, false
	}
	if literal[0] == '.' && strings.Count(literal, ".") == 1 {
		return CLASS_SELECTOR, true
	}
	return SUBKIND_NONE, false
}

func IsIdSelector(literal string) (subkind Subkind, ok bool) {
	if len(literal) == 0 {
		return SUBKIND_NONE, false
	}
	if literal[0] == '#' && strings.Count(literal, "#") == 1 {
		return ID_SELECTOR, true
	}
	return SUBKIND_NONE, false
}

func IsUniversalSelector(literal string) (subkind Subkind, ok bool) {
	if len(literal) == 0 {
		return SUBKIND_NONE, false
	}
	if literal == "*" {
		return UNIVERSAL_SELECTOR, true
	}
	return SUBKIND_NONE, false
}

func IsAttributeSelector(literal string) (subkind Subkind, ok bool) {
	if len(literal) == 0 {
		return SUBKIND_NONE, false
	}
	if literal[0] == '[' && literal[len(literal)-1] == ']' {
		return ATTRIBUTE_SELECTOR, true
	}
	return SUBKIND_NONE, false
}

func IsInheritSelector(literal string) (subkind Subkind, ok bool) {
	if len(literal) == 0 {
		return SUBKIND_NONE, false
	}
	if literal[0] == '&' {
		return INHERIT_SELECTOR, true
	}
	return SUBKIND_NONE, false
}
