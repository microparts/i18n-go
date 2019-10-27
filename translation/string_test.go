package translation

import (
	"reflect"
	"testing"
)

func TestNewString(t *testing.T) {
	exist := NewString("cn", "Test")
	expected := String{
		Translate: map[string]string{
			"cn": "Test",
		},
	}

	if !reflect.DeepEqual(expected, exist) {
		t.Errorf("failed to create a new string: exist: %v; exp.: %v", exist, expected)
	}
}

func TestString_Init(t *testing.T) {
	exist := String{}
	exist.Init()

	expected := String{Translate: map[string]string{}}
	if !reflect.DeepEqual(expected, exist) {
		t.Errorf("failed to create a new string: exist: %v; exp.: %v", exist, expected)
	}
}

func TestString_ApplyTranslationCtx(t *testing.T) {
	str := (&String{}).
		AddTranslate("ru", "Привет").
		AddTranslate("en", "Hello")

	conf := &Conf{
		Display:  "ru",
		Fallback: "en",
		Second:   "en",
	}

	expected := &String{
		Display:    "Привет",
		Second:     "Hello",
		ctxApplied: true,
	}

	str.ApplyTranslationCtx(conf)
	if !reflect.DeepEqual(expected, str) {
		t.Errorf("failed to apply a translation context: exist: %v; exp.: %v", str, expected)
	}

	str.ApplyTranslationCtx(conf)
	if !reflect.DeepEqual(expected, str) {
		t.Errorf("failed to repetedly apply a translation context: exist: %v; exp.: %v", str, expected)
	}
}

func TestString_ResetCtxApplied(t *testing.T) {
	str := (&String{}).
		AddTranslate("ru", "Привет").
		AddTranslate("en", "Hello").
		ApplyTranslationCtx(&Conf{
			Display:         "ru",
			Fallback:        "en",
			Second:          "en",
			TranslationList: true,
		})

	if !str.ctxApplied {
		t.Errorf("failed to set a flag of already applied translaction context: exist: %v; exp.: %v", str.ctxApplied, true)
	}

	str.ResetCtxApplied()

	if str.ctxApplied {
		t.Errorf("failed to reset a flag of already applied translaction context: exist: %v; exp.: %v", str.ctxApplied, false)
	}
}

func TestString_ClearContext(t *testing.T) {
	str := (&String{}).
		AddTranslate("ru", "Привет").
		AddTranslate("en", "Hello").
		ApplyTranslationCtx(&Conf{
			Display:         "ru",
			Fallback:        "en",
			Second:          "en",
			TranslationList: true,
		})

	expected := &String{
		Display: "Привет",
		Second:  "Hello",
		Translate: map[string]string{
			"ru": "Привет",
			"en": "Hello",
		},
		ctxApplied: true,
	}

	if !reflect.DeepEqual(expected, str) {
		t.Errorf("failed to apply a translation context: exist: %v; exp.: %v", str, expected)
	}

	expected = &String{
		Translate: map[string]string{
			"ru": "Привет",
			"en": "Hello",
		},
		ctxApplied: true,
	}

	str.ClearContext()
	if !reflect.DeepEqual(expected, str) {
		t.Errorf("failed to apply a translation context: exist: %v; exp.: %v", str, expected)
	}
}

func TestString_Reset(t *testing.T) {
	str := (&String{}).
		AddTranslate("ru", "Привет").
		AddTranslate("en", "Hello").
		ApplyTranslationCtx(&Conf{
			Display:         "ru",
			Fallback:        "en",
			Second:          "en",
			TranslationList: true,
		})

	str.Reset()
	if !reflect.DeepEqual(&String{}, str) {
		t.Errorf("failed to reset a string")
	}
}

func TestString_Clone(t *testing.T) {
	var str *String

	if str.Clone() != nil {
		t.Errorf("failed to clone a nil string")
	}

	expected := (&String{}).
		AddTranslate("ru", "hello").
		ApplyTranslationCtx(&Conf{
			Display:         "ru",
			Fallback:        "en",
			Second:          "en",
			TranslationList: true,
		})

	exist := expected.Clone()
	if !reflect.DeepEqual(expected, exist) {
		t.Errorf("failed to clone a translation string: exist: %v; exp.: %v", exist, expected)
	}

	exist.ClearContext()
	if reflect.DeepEqual(expected, exist) {
		t.Errorf("failed to clone a translation string, because it isn't a clone: exist: %v; exp.: %v", exist, expected)
	}
}

func TestString_Empty(t *testing.T) {
	if !(&String{}).Empty() {
		t.Error("failed to check an empty for empty string")
	}

	str := (&String{}).
		AddTranslate("ru", "Привет").
		AddTranslate("en", "Hello!")
	if str.Empty() {
		t.Error("failed to check an empty for a string")
	}

	if ln, expected := str.Len(), len("Привет"); ln != expected {
		t.Errorf("failed to check a length for a nonempty string: exist: %d; exp.: %d", ln, expected)
	}
}

func TestString_HasTranslation(t *testing.T) {
	if (&String{}).HasTranslation() {
		t.Error("failed to check an empty string")
	}

	str := (&String{}).
		AddTranslate("ru", "Привет").
		AddTranslate("en", "Hello!")
	if !str.HasTranslation() {
		t.Error("failed to check a translation for a string containing a source of translation")
	}

	str.ApplyTranslationCtx(DefaultConf())
	if !str.HasTranslation() {
		t.Error("failed to check a translation for a string containing a result of translation")
	}
}

func TestString_Update(t *testing.T) {
	exist := (&String{}).
		AddTranslate("ru", "Привет").
		AddTranslate("en", "Hello")

	str := (&String{}).
		AddTranslate("ru", "Привет!").
		AddTranslate("fr", "Bonjour")

	exist.Update(*str)

	expected := &String{
		Translate: map[string]string{
			"ru": "Привет!",
			"en": "Hello",
			"fr": "Bonjour",
		},
	}
	if !reflect.DeepEqual(expected, exist) {
		t.Errorf("failed to update a translation string: exist: %v; exp.: %v", exist, expected)
	}
}

func TestString_Join(t *testing.T) {
	str := (&String{}).
		AddTranslate("ru", "Привет").
		AddTranslate("en", "Hello")

	exist := str.Join(
		*(&String{}).
			AddTranslate("ru", "Привет!").
			AddTranslate("fr", "Bonjour"),
		", ",
	)

	expected := String{
		Translate: map[string]string{
			"ru": "Привет, Привет!",
			"en": "Hello",
			"fr": "Bonjour",
		},
	}
	if !reflect.DeepEqual(expected, exist) {
		t.Errorf("failed to join a translation string: exist: %v; exp.: %v", exist, expected)
	}
}

func TestString_Add(t *testing.T) {
	str := (&String{}).
		AddTranslate("ru", "Привет").
		AddTranslate("en", "Hello")

	str.Add(
		*(&String{}).
			AddTranslate("ru", ", Привет!").
			AddTranslate("fr", "Bonjour"),
	)

	expected := &String{
		Translate: map[string]string{
			"ru": "Привет, Привет!",
			"en": "Hello",
			"fr": "Bonjour",
		},
	}
	if !reflect.DeepEqual(expected, str) {
		t.Errorf("failed to join a translation string: exist: %v; exp.: %v", str, expected)
	}
}

func TestString_Trim(t *testing.T) {
	str := (&String{}).
		AddTranslate("ru", "		Привет    ").
		AddTranslate("en", "    Hello ")

	str.Trim()

	expected := &String{
		Translate: map[string]string{
			"ru": "Привет",
			"en": "Hello",
		},
	}
	if !reflect.DeepEqual(expected, str) {
		t.Errorf("failed to trim spaces for a translation string: exist: %v; exp.: %v", str, expected)
	}
}

func TestString_Map(t *testing.T) {
	str := (&String{}).Map(func(string) string { return "" })
	if !reflect.DeepEqual(String{}, str) {
		t.Errorf("failed to apply map for an empty string")
	}

}

func TestString_GetTranslate(t *testing.T) {
	translation := map[string]string{
		"ru": "Привет",
		"en": "Hello",
	}

	str := &String{}

	for lang, v := range translation {
		str.AddTranslate(lang, v)
	}

	for lang, expected := range translation {
		if exist := str.GetTranslate(lang); exist != expected {
			t.Errorf("failed to get a translation for a lang '%s': exist: '%s'; exp.: '%s'", lang, exist, expected)
		}
	}

	if str.GetTranslate("unknown") != "" {
		t.Error("failed to get a translation for an unknown language")
	}
}

func TestString_String(t *testing.T) {
	str := (&String{}).AddTranslate("ru", "Привет").String()
	expected := `{"ru":"Привет"}`

	if str != expected {
		t.Errorf("failed to stringify a translation string. Exist: %v, expected: %v", str, expected)
	}
}
