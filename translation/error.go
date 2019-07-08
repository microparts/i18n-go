package translation

const (
	ErrConfFallbackEmpty = ErrTranslation("lang/fallback - empty, using '" + defaultLang + "'")
	ErrConfSecondEmpty   = ErrTranslation("lang/second - empty, using '" + defaultLang + "'")
)

type ErrTranslation string

func (o ErrTranslation) Error() string {
	return string(o)
}
