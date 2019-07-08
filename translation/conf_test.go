package translation

import (
	"reflect"
	"testing"
)

func TestDefaultConf(t *testing.T) {
	defaultConf := DefaultConf()
	conf := (&Conf{}).CheckDefault()

	if !reflect.DeepEqual(defaultConf, conf) {
		t.Errorf("failed to check a default config")
	}
}

func TestConf_Validate(t *testing.T) {
	conf := &Conf{}

	errList := conf.Validate()
	if len(errList) == 0 {
		t.Errorf("failed to catch errors")
	} else {
		for _, err := range errList {
			switch err {
			case ErrConfFallbackEmpty, ErrConfSecondEmpty:
			default:
				t.Errorf("unknown err: %v", err)
			}
		}
	}
}
