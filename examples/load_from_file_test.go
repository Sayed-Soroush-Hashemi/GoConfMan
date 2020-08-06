package examples

import (
	"github.com/Sayed-Soroush-Hashemi/GoConfMan/pkg/goconfman"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestLoadFromFileOnGlobalConfig(t *testing.T) {
	g := GlobalConfig{}
	err := goconfman.LoadFromFile(&g, "global_config.json")
	if err != nil {
		t.Fatalf("Error in LoadFromMap: %s", err.Error())
	}

	expectedG := GlobalConfig{
		IntegerValue:       1234,
		FloatValue:         987.3,
		StringValue:        "globalConfig.StringValue in global_config.json",
		LocalConfig:        LocalConfig{
			IntegerValue:      42,
			FloatValue:        0,
			StringValue:       "",
			SliceValue:        nil,
			SliceOfSliceValue: nil,
			MapValue:          nil,
			ComplicatedValue:  nil,
		},
		NonGoConfManConfig: NonGoConfManConfig{
			IntegerValue: 0,
			FloatValue:   0,
			StringValue:  "",
			LocalConfig:  LocalConfig{
				IntegerValue:      71,
				FloatValue:        0,
				StringValue:       "",
				SliceValue:        nil,
				SliceOfSliceValue: nil,
				MapValue:          nil,
				ComplicatedValue:  nil,
			},
		},
	}

	if reflect.DeepEqual(g, expectedG) == false {
		t.Errorf("g is different from the expectedG. here's the diff: \n%s", cmp.Diff(g, expectedG))
	}
}

func TestLoadFromFileOnNonGoconfmanConfig(t *testing.T) {
	ng := NonGoConfManConfig{}
	err := goconfman.LoadFromFile(&ng, "non_goconfman_config.json")
	if err != nil {
		t.Fatalf("Error in LoadFromMap: %s", err.Error())
	}

	expectedNG := NonGoConfManConfig{
		IntegerValue: 0,
		FloatValue:   0,
		StringValue:  "",
		LocalConfig:  LocalConfig{
			IntegerValue:      71,
			FloatValue:        13.31,
			StringValue:       "",
			SliceValue:        nil,
			SliceOfSliceValue: nil,
			MapValue:          nil,
			ComplicatedValue:  nil,
		},
	}

	if reflect.DeepEqual(ng, expectedNG) == false {
		t.Errorf("ng is different from the expectedNG. here's the diff: \n%s", cmp.Diff(ng, expectedNG))
	}
}

func TestLoadFromFileAfterLoadFromDefaultsOnGlobalConfig(t *testing.T) {
	g := GlobalConfig{}
	goconfman.LoadFromDefaults(&g)
	err := goconfman.LoadFromFile(&g, "global_config.json")
	if err != nil {
		t.Fatalf("Error in LoadFromMap: %s", err.Error())
	}

	expectedG := GlobalConfig{
		IntegerValue:       1234,
		FloatValue:         987.3,
		StringValue:        "globalConfig.StringValue in global_config.json",
		LocalConfig:        LocalConfig{
			IntegerValue:      42,
			FloatValue:        31.4,
			StringValue:       "in local config",
			SliceValue:        nil,
			SliceOfSliceValue: nil,
			MapValue:          nil,
			ComplicatedValue:  nil,
		},
		NonGoConfManConfig: NonGoConfManConfig{
			IntegerValue: 0,
			FloatValue:   0,
			StringValue:  "",
			LocalConfig:  LocalConfig{
				IntegerValue:      71,
				FloatValue:        31.4,
				StringValue:       "in local config",
				SliceValue:        nil,
				SliceOfSliceValue: nil,
				MapValue:          nil,
				ComplicatedValue:  nil,
			},
		},
	}

	if reflect.DeepEqual(g, expectedG) == false {
		t.Errorf("g is different from the expectedG. here's the diff: \n%s", cmp.Diff(g, expectedG))
	}
}