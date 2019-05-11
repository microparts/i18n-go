package translation

import (
	"net/http"
	"strings"
)

const (
	headerDisplay        = "X-Lang-Display"
	headerFallback       = "X-Lang-Fallback"
	headerSecond         = "X-Lang-Second"
	headerTranslateList  = "X-Lang-TranslateList"
	headerTranslateListL = "X-Lang-Translatelist" // header's name with lower case version of TranslateList
)

type ApplyTranslationCtx interface {
	ApplyTranslationCtx(Context)
}

type Context interface {
	GetDisplay() string
	GetFallback() string
	GetSecond() string
	GetTranslationList() bool
	CheckDefault() Context
	Validate() []error
}

func NewContext(base Context, req *http.Request) Context {
	ctx := &Conf{}

	if _, ok := req.Header[headerDisplay]; ok {
		ctx.Display = req.Header.Get(headerDisplay)
	} else {
		ctx.Display = base.GetDisplay()
	}

	if _, ok := req.Header[headerFallback]; ok {
		ctx.Fallback = req.Header.Get(headerFallback)
	} else {
		ctx.Fallback = base.GetFallback()
	}

	if _, ok := req.Header[headerSecond]; ok {
		ctx.Second = req.Header.Get(headerSecond)
	} else {
		ctx.Second = base.GetSecond()
	}

	if _, ok := req.Header[headerTranslateList]; ok {
		ctx.TranslationList = strings.ToLower(req.Header.Get(headerTranslateList)) == "true"
	} else if _, ok := req.Header[headerTranslateListL]; ok {
		ctx.TranslationList = strings.ToLower(req.Header.Get(headerTranslateListL)) == "true"
	} else {
		ctx.TranslationList = base.GetTranslationList()
	}

	return ctx
}
