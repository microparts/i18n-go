package translation

import (
	"net/http"
	"reflect"
	"testing"
)

var (
	testTranslationHeaders = map[string]map[string]string{
		"ru": {
			headerDisplay:  "ru",
			headerFallback: "en",
			headerSecond:   "en",
		},
		"ru_fr": {
			headerDisplay:  "ru",
			headerFallback: "fr",
			headerSecond:   "en",
		},
		"ru_ext": {
			headerDisplay:       "ru",
			headerFallback:      "en",
			headerSecond:        "en",
			headerTranslateList: "true",
		},
		"ru_ext_": {
			headerDisplay:        "ru",
			headerFallback:       "en",
			headerSecond:         "en",
			headerTranslateListL: "true",
		},
		"en": {
			headerDisplay:  "en",
			headerFallback: "ru",
			headerSecond:   "ru",
		},
		"en_fr": {
			headerDisplay:  "en",
			headerFallback: "fr",
			headerSecond:   "ru",
		},
		"en_ext": {
			headerDisplay:       "en",
			headerFallback:      "ru",
			headerSecond:        "ru",
			headerTranslateList: "true",
		},
	}
)

func newTestRequest(headerId string) *http.Request {
	req, _ := http.NewRequest("GET", "/", nil)

	if headers, ok := testTranslationHeaders[headerId]; ok {
		for name, value := range headers {
			req.Header.Set(name, value)
		}

		// to prevent making canonical of header name
		if v, ok := headers[headerTranslateList]; ok {
			req.Header[headerTranslateList] = append(req.Header[headerTranslateList], v)
		}
	}

	return req
}

func TestNewContext(t *testing.T) {
	src := map[string]*String{
		"ru_en": {
			Translate: map[string]string{
				"ru": "Привет",
				"en": "Hello",
			},
		},
		"en_fr": {
			Translate: map[string]string{
				"en": "Hello",
				"fr": "Bonjour",
			},
		},
	}

	defaultCtx := &Conf{
		Display:         "ru",
		Fallback:        "en",
		Second:          "en",
		TranslationList: false,
	}

	testSuites := []*struct {
		in       *String
		headerId string
		expected *String
	}{
		{
			in:       src["ru_en"].Clone(),
			headerId: "unknown",
			expected: &String{
				Display:    "Привет",
				Second:     "Hello",
				ctxApplied: true,
			},
		},
		{
			in:       src["ru_en"].Clone(),
			headerId: "ru",
			expected: &String{
				Display:    "Привет",
				Second:     "Hello",
				ctxApplied: true,
			},
		},
		{
			in:       src["en_fr"].Clone(),
			headerId: "ru_fr",
			expected: &String{
				Display:    "Bonjour",
				Second:     "Hello",
				ctxApplied: true,
			},
		},
		{
			in:       src["ru_en"].Clone(),
			headerId: "ru_ext",
			expected: &String{
				Display: "Привет",
				Second:  "Hello",
				Translate: map[string]string{
					"ru": "Привет",
					"en": "Hello",
				},
				ctxApplied: true,
			},
		},
		{
			in:       src["ru_en"].Clone(),
			headerId: "ru_ext_",
			expected: &String{
				Display: "Привет",
				Second:  "Hello",
				Translate: map[string]string{
					"ru": "Привет",
					"en": "Hello",
				},
				ctxApplied: true,
			},
		},
	}

	for i, test := range testSuites {
		test.in.ApplyTranslationCtx(NewContext(defaultCtx, newTestRequest(test.headerId)))

		if !reflect.DeepEqual(test.in, test.expected) {
			t.Errorf("%d: failed to apply context: exist: %v; exp.: %v", i, test.in, test.expected)
		}
	}
}
