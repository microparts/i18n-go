package translation

const (
	defaultLang = "en"
)

type Conf struct {
	Display         string `json:"display" yaml:"display"`
	Fallback        string `json:"fallback" yaml:"fallback"`
	Second          string `json:"second" yaml:"second"`
	TranslationList bool   `json:"translate_list" yaml:"translate_list"`
}

func DefaultConf() *Conf {
	return &Conf{
		Fallback: defaultLang,
		Second:   defaultLang,
	}
}

func (o *Conf) GetDisplay() string {
	return o.Display
}

func (o *Conf) GetFallback() string {
	return o.Fallback
}

func (o *Conf) GetSecond() string {
	return o.Second
}

func (o *Conf) GetTranslationList() bool {
	return o.TranslationList
}

func (o *Conf) CheckDefault() Context {
	if len(o.Fallback) == 0 {
		o.Fallback = defaultLang
	}
	if len(o.Second) == 0 {
		o.Second = defaultLang
	}

	return o
}

func (o *Conf) Validate() (errList []error) {
	if len(o.Fallback) == 0 {
		errList = append(errList, ErrConfFallbackEmpty)
	}

	if len(o.Second) == 0 {
		errList = append(errList, ErrConfSecondEmpty)
	}

	o.CheckDefault()

	return
}
