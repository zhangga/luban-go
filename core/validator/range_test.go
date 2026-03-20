package validator

import (
	"testing"

	"github.com/zhangga/luban-go/core/datas"
)

func TestRangeValidator(t *testing.T) {
	v := &RangeValidator{}
	
	// Test int range
	err := v.Validate(nil, datas.NewDInt(5), "[1,10]")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	err = v.Validate(nil, datas.NewDInt(0), "[1,10]")
	if err == nil {
		t.Errorf("Expected error for out of range value")
	}

	err = v.Validate(nil, datas.NewDInt(10), "(1,10)")
	if err == nil {
		t.Errorf("Expected error for exclusive bound")
	}

	// Test float range
	err = v.Validate(nil, datas.NewDFloat(5.5), "[1.0, 10.0]")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	// Test string length
	err = v.Validate(nil, datas.NewDString("hello"), "[1,5]")
	if err != nil {
		t.Errorf("Expected nil for valid string length, got %v", err)
	}

	err = v.Validate(nil, datas.NewDString("hello world"), "[1,5]")
	if err == nil {
		t.Errorf("Expected error for string too long")
	}
}
