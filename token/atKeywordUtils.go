package token

func IsAtKeyword(value string) (subkind Subkind, ok bool) {
	if subkind, ok = IsCssAtKeyword(value); ok {
		return
	}
	if subkind, ok = IsZcssAtKeyword(value); ok {
		return
	}
	return SUBKIND_NONE, false
}

var CssAtKeywords = map[string]Subkind{
	"charset":             CHARSET_AT_KEYWORD,
	"color-profile":       COLOR_PROFILE_AT_KEYWORD,
	"container":           CONTAINER_AT_KEYWORD,
	"counter-style":       COUNTER_STYLE_AT_KEYWORD,
	"document":            DOCUMENT_AT_KEYWORD,
	"font-face":           FONT_FACE_AT_KEYWORD,
	"font-feature-values": FONT_FEATURE_VALUES_AT_KEYWORD,
	"swash":               SWASH_AT_KEYWORD,
	"annotation":          ANNOTATION_AT_KEYWORD,
	"ornaments":           ORNAMENTS_AT_KEYWORD,
	"stylistic":           STYLISTIC_AT_KEYWORD,
	"styleset":            STYLESET_AT_KEYWORD,
	"character-variant":   CHARACTER_VARIANT_AT_KEYWORD,
	"font-palette-values": FONT_PALETTE_FEATURES_AT_KEYWORD,
	"import":              IMPORT_AT_KEYWORD,
	"keyframes":           KEYFRAMES_AT_KEYWORD,
	"layer":               LAYER_AT_KEYWORD,
	"media":               MEDIA_AT_KEYWORD,
	"namespace":           NAMESPACE_AT_KEYWORD,
	"page":                PAGE_AT_KEYWORD,
	"position-try":        POSITION_TRY_AT_KEYWORD,
	"property":            PROPERTY_AT_KEYWORD,
	"scope":               SCOPE_AT_KEYWORD,
	"starting-style":      STARTING_STYLE_AT_KEYWORD,
	"supports":            SUPPORTS_AT_KEYWORD,
	"view-transition":     VIEW_TRANSITION_AT_KEYWORD,
}

func IsCssAtKeyword(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssAtKeywords[value]
	return
}

var ZcssAtKeywords = map[string]Subkind{
	"function": FUNCTION_AT_KEYWORD,
	"if":       IF_AT_KEYWORD,
	"elseif":   ELSE_IF_AT_KEYWORD,
	"else":     ELSE_AT_KEYWORD,
	"mixin":    MIXIN_AT_KEYWORD,
}

func IsZcssAtKeyword(value string) (subkind Subkind, ok bool) {
	subkind, ok = ZcssAtKeywords[value]
	return
}
