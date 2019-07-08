i18n-go
------------

[![CircleCI](https://circleci.com/gh/microparts/i18n-go.svg?style=shield)](https://circleci.com/gh/microparts/i18n-go)
[![codecov](https://codecov.io/gh/microparts/i18n-go/graph/badge.svg)](https://codecov.io/gh/microparts/i18n-go)

A small set of helpers for using as i18n part of companies services and applications.

# Translation

## Usage

Basic usage a translation string and a context in a record and handling HTTP request 

```go
package pack

import (
	"net/http"
	"encoding/json"

	"github.com/microparts/i18n-go/translation"
)

var (
	// a configuration of default properties of translation context
	// it could be loaded from json or yaml configuration file
	Conf = translation.Conf{
		Display: "ru",
		Fallback: "en",
		Second:  "en",
		TranslationList: false,
	}
)

// A model definition

type Record struct {
	Id   int                `json:"id"`                  
	Name translation.String `json:"name"`
}

func (o *Record) ApplyTranslationCtx(ctx translation.Context) {
	o.Name.ApplyTranslationCtx(ctx)
}

// A handler of an http request

func Handler(w http.ResponseWriter, req *http.Request) {
	translateCtx := translation.NewContext(&Conf, req)
	
	rec := &Record{
	    Id: 10,
	    Name: translation.String {
	        Translate: map[string]string {
    		    "en": "Hello world!",
    		    "ru": "Здравствуй мир!",
	    	},
	    },
	}
	
	rec.ApplyTranslationCtx(translateCtx)
	
	json.NewEncoder(w).Encode(rec)
}
```

### Response

An expected output without headers:

```json
{
  "id": 10,
  "name": {
    "display": "Здравствуй мир!",
    "second": "Hello world!",
    "translate": null
  }   
}
```

An expected output for a request with headers:

```http request
GET /record
X-Lang-Display: en
X-Lang-Second: ru
X-Lang-TranslateList: true
```

a response:

```json
{
  "id": 10,
  "name": {
    "display": "Здравствуй world!",
    "second": "Привет мир!",
    "translate": {
      "ru": "Hello world!",
      "en": "Здравствуй мир!"
    }
  }
}
```

### Request

An example of a request to update of translation string

```http request
PUT /record
Content-Type: application/json

{
  "id": 12,
  "name": {
    "translate": {
      "ru": "Hi world!",
      "en": "Привет мир!"
    }
  }
}
```
