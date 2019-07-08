package translation

import "testing"

func TestRandomStr(t *testing.T) {
	if len(RandomStr()) == 0 {
		t.Error("failed to generate random string")
	}
}

func TestGenerateString(t *testing.T) {
	languages := []string{
		"ru",
		"fr",
		"en",
		"cn",
	}

	str := GenerateString(languages)
	translate := str.Translate
	if len(languages) != len(translate) {
		t.Errorf("failed to generate all languages string: generate %d, but expected %d", len(languages), len(translate))
	} else {
		for _, lang := range languages {
			if _, ok := translate[lang]; !ok {
				t.Errorf("failed to generate a string for lang '%s'", lang)
			}
		}
	}
}
