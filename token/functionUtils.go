package token

func IsCssBuiltinFunction(value string) (subkind Subkind, ok bool) {
	if subkind, ok = IsCssAnchorPositioningFunction(value); ok {
		return
	}
	if subkind, ok = IsCssAnimationFunction(value); ok {
		return
	}
	if subkind, ok = IsCssColorFunction(value); ok {
		return
	}
	if subkind, ok = IsCssCounterFunction(value); ok {
		return
	}
	if subkind, ok = IsCssEasingFunction(value); ok {
		return
	}
	if subkind, ok = IsCssFilterFunction(value); ok {
		return
	}
	if subkind, ok = IsCssFontFunction(value); ok {
		return
	}
	if subkind, ok = IsCssGridFunction(value); ok {
		return
	}
	if subkind, ok = IsCssImageFunction(value); ok {
		return
	}
	if subkind, ok = IsCssMathFunction(value); ok {
		return
	}
	if subkind, ok = IsCssMatrixFunction(value); ok {
		return
	}
	if subkind, ok = IsCssPerspectiveFunction(value); ok {
		return
	}
	if subkind, ok = IsCssReferenceFunction(value); ok {
		return
	}
	if subkind, ok = IsCssRotateFunction(value); ok {
		return
	}
	if subkind, ok = IsCssScaleFunction(value); ok {
		return
	}
	if subkind, ok = IsCssShapeFunction(value); ok {
		return
	}
	if subkind, ok = IsCssSkewFunction(value); ok {
		return
	}
	if subkind, ok = IsCssTranslateFunction(value); ok {
		return
	}
	return SUBKIND_NONE, false
}

var CssAnchorPositioningFunctions = map[string]Subkind{
	"anchor":      ANCHOR_POSITIONING_FUNCTION,
	"anchor-size": ANCHOR_POSITIONING_FUNCTION,
}

func IsCssAnchorPositioningFunction(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssAnchorPositioningFunctions[value]
	return
}

var CssAnimationFunctions = map[string]Subkind{
	"scroll": ANIMATION_FUNCTION,
	"view":   ANIMATION_FUNCTION,
}

func IsCssAnimationFunction(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssAnimationFunctions[value]
	return
}

var CssColorFunctions = map[string]Subkind{
	"rgb":            COLOR_FUNCTION,
	"hsl":            COLOR_FUNCTION,
	"hwb":            COLOR_FUNCTION,
	"lch":            COLOR_FUNCTION,
	"oklch":          COLOR_FUNCTION,
	"lab":            COLOR_FUNCTION,
	"oklab":          COLOR_FUNCTION,
	"color":          COLOR_FUNCTION,
	"color-mix":      COLOR_FUNCTION,
	"color-contrast": COLOR_FUNCTION,
	"device-cmyk":    COLOR_FUNCTION,
	"light-dark":     COLOR_FUNCTION,
}

func IsCssColorFunction(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssColorFunctions[value]
	return
}

var CssCounterFunctions = map[string]Subkind{
	"counter":  COUNTER_FUNCTION,
	"counters": COUNTER_FUNCTION,
	"symbols":  COUNTER_FUNCTION,
}

func IsCssCounterFunction(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssCounterFunctions[value]
	return
}

var CssEasingFunctions = map[string]Subkind{
	"linear":       EASING_FUNCTION,
	"cubic-bezier": EASING_FUNCTION,
	"steps":        EASING_FUNCTION,
}

func IsCssEasingFunction(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssEasingFunctions[value]
	return
}

var CssFilterFunctions = map[string]Subkind{
	"blur":       FILTER_FUNCTION,
	"brightness": FILTER_FUNCTION,
	"contrast":   FILTER_FUNCTION,
	"grayscale":  FILTER_FUNCTION,
	"hue-rotate": FILTER_FUNCTION,
	"invert":     FILTER_FUNCTION,
	"opacity":    FILTER_FUNCTION,
	"saturate":   FILTER_FUNCTION,
	"sepia":      FILTER_FUNCTION,
}

func IsCssFilterFunction(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssFilterFunctions[value]
	return
}

var CssFontFunctions = map[string]Subkind{
	"stylistic":         FONT_FUNCTION,
	"styleset":          FONT_FUNCTION,
	"character-variant": FONT_FUNCTION,
	"swash":             FONT_FUNCTION,
	"ornaments":         FONT_FUNCTION,
	"annotation":        FONT_FUNCTION,
	// -- not listed on MDN
	"local":  FONT_FUNCTION,
	"format": FONT_FUNCTION,
	"tech":   FONT_FUNCTION,
}

func IsCssFontFunction(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssFontFunctions[value]
	return
}

var CssGridFunctions = map[string]Subkind{
	"fit-content": GRID_FUNCTION,
	"minmax":      GRID_FUNCTION,
	"repeat":      GRID_FUNCTION,
}

func IsCssGridFunction(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssGridFunctions[value]
	return
}

var CssImageFunctions = map[string]Subkind{
	"linear-gradient":           IMAGE_FUNCTION,
	"radial-graidnet":           IMAGE_FUNCTION,
	"conic-gradient":            IMAGE_FUNCTION,
	"repeating-linear-gradient": IMAGE_FUNCTION,
	"repeating-radial-graidnet": IMAGE_FUNCTION,
	"repeating-conic-gradient":  IMAGE_FUNCTION,
	"image":                     IMAGE_FUNCTION,
	"image-set":                 IMAGE_FUNCTION,
	"cross-fade":                IMAGE_FUNCTION,
	"element":                   IMAGE_FUNCTION,
	"paint":                     IMAGE_FUNCTION,
}

func IsCssImageFunction(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssImageFunctions[value]
	return
}

var CssMathFunctions = map[string]Subkind{
	"calc":  MATH_FUNCTION,
	"max":   MATH_FUNCTION,
	"min":   MATH_FUNCTION,
	"clamp": MATH_FUNCTION,
	"round": MATH_FUNCTION,
	"mod":   MATH_FUNCTION,
	"sin":   MATH_FUNCTION,
	"cos":   MATH_FUNCTION,
	"tan":   MATH_FUNCTION,
	"asin":  MATH_FUNCTION,
	"acos":  MATH_FUNCTION,
	"atan":  MATH_FUNCTION,
	"atan2": MATH_FUNCTION,
	"pow":   MATH_FUNCTION,
	"sqrt":  MATH_FUNCTION,
	"log":   MATH_FUNCTION,
	"exp":   MATH_FUNCTION,
	"abs":   MATH_FUNCTION,
	"sign":  MATH_FUNCTION,
}

func IsCssMathFunction(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssMathFunctions[value]
	return
}

var CssMatrixFunctions = map[string]Subkind{
	"matrix":   MATRIX_FUNCTION,
	"matrix3d": MATRIX_FUNCTION,
}

func IsCssMatrixFunction(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssMatrixFunctions[value]
	return
}

var CssPerspectiveFunctions = map[string]Subkind{
	"perspective": PERSPECTIVE_FUNCTION,
}

func IsCssPerspectiveFunction(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssPerspectiveFunctions[value]
	return
}

var CssReferenceFunctions = map[string]Subkind{
	"attr": REFERENCE_FUNCTION,
	"env":  REFERENCE_FUNCTION,
	"url":  REFERENCE_FUNCTION,
	"var":  REFERENCE_FUNCTION,
}

func IsCssReferenceFunction(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssReferenceFunctions[value]
	return
}

var CssRotateFunctions = map[string]Subkind{
	"rotateX":  ROTATE_FUNCTION,
	"rotateY":  ROTATE_FUNCTION,
	"rotateZ":  ROTATE_FUNCTION,
	"rotate":   ROTATE_FUNCTION,
	"rotate3d": ROTATE_FUNCTION,
}

func IsCssRotateFunction(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssRotateFunctions[value]
	return
}

var CssScaleFunctions = map[string]Subkind{
	"scaleX":  SCALE_FUNCTION,
	"scaleY":  SCALE_FUNCTION,
	"scaleZ":  SCALE_FUNCTION,
	"scale":   SCALE_FUNCTION,
	"scale3d": SCALE_FUNCTION,
}

func IsCssScaleFunction(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssScaleFunctions[value]
	return
}

var CssShapeFunctions = map[string]Subkind{
	"circle":  SHAPE_FUNCTION,
	"ellipse": SHAPE_FUNCTION,
	"inset":   SHAPE_FUNCTION,
	"rect":    SHAPE_FUNCTION,
	"xywh":    SHAPE_FUNCTION,
	"polygon": SHAPE_FUNCTION,
	"path":    SHAPE_FUNCTION,
	"shape":   SHAPE_FUNCTION,
}

func IsCssShapeFunction(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssShapeFunctions[value]
	return
}

var CssSkewFunctions = map[string]Subkind{
	"skewX": SKEW_FUNCTION,
	"skewY": SKEW_FUNCTION,
	"skew":  SKEW_FUNCTION,
}

func IsCssSkewFunction(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssSkewFunctions[value]
	return
}

var CssTranslateFunctions = map[string]Subkind{
	"translateX":  TRANSLATE_FUNCTION,
	"translateY":  TRANSLATE_FUNCTION,
	"translateZ":  TRANSLATE_FUNCTION,
	"translate":   TRANSLATE_FUNCTION,
	"translate3d": TRANSLATE_FUNCTION,
}

func IsCssTranslateFunction(value string) (subkind Subkind, ok bool) {
	subkind, ok = CssTranslateFunctions[value]
	return
}
