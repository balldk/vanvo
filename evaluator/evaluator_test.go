package evaluator

import (
	"testing"
	"vila/errorhandler"
	"vila/lexer"
	"vila/object"
	"vila/parser"
)

func testEval(t *testing.T, input string) object.Object {
	lexerErr := errorhandler.NewErrorList(input, "")
	parserErr := errorhandler.NewErrorList(input, "")
	evaluatorErr := errorhandler.NewErrorList(input, "")

	l := lexer.New(input, lexerErr)
	p := parser.New(l, parserErr)
	ev := New(evaluatorErr)

	program := p.ParseProgram()
	value := ev.Eval(program)

	if lexerErr.NotEmpty() {
		t.Fatalf("input %q has errors: \n%s", input, lexerErr)
	} else if parserErr.NotEmpty() {
		t.Fatalf("input %q has errors: \n%s", input, parserErr)
	} else if evaluatorErr.NotEmpty() {
		t.Fatalf("input %q has errors: \n%s", input, evaluatorErr)
	}

	return value
}

func testIntObject(t *testing.T, obj object.Object, expected int64) {
	result, ok := obj.(*object.Int)
	if !ok {
		t.Fatalf("object is not *object.Int. got=%#v, expected=%d", obj, expected)
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. want=%d, got=%d", expected, result.Value)
	}
}

func TestEvalInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"3", 3},
		{"-3", -3},
		{"--3", 3},
		{"-3 + 4", 1},
	}

	for _, test := range tests {
		value := testEval(t, test.input)
		testIntObject(t, value, test.expected)
	}
}

func testRealObject(t *testing.T, obj object.Object, expected float64) {
	f, ok := obj.(*object.Real)
	if !ok {
		t.Errorf("object is not *object.Real. got=%#v", obj)
		return
	}

	if f.Value != expected {
		t.Errorf("object has wrong value. want=%f, got=%f", expected, f.Value)
	}
}

func TestEvalReal(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"3.14", 3.14},
		{"--3.14", 3.14},
		{"-3.14", -3.14},
		{"3 + 0.14", 3.14},
	}

	for _, test := range tests {
		value := testEval(t, test.input)
		testRealObject(t, value, test.expected)
	}
}
