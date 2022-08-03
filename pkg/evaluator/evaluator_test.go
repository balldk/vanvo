package evaluator

// import (
// 	"testing"
// 	"vila/pkg/errorhandler"
// 	"vila/pkg/lexer"
// 	"vila/pkg/object"
// 	"vila/pkg/parser"
// )

// func testEval(t *testing.T, input string) object.Object {
// 	errors := errorhandler.NewErrorList(input, "")
// 	env := object.NewEnvironment()

// 	l := lexer.New(input, errors)
// 	p := parser.New(l, errors)
// 	ev := New(env, errors)

// 	program := p.ParseProgram()
// 	value := ev.Eval(program)

// 	if errors.NotEmpty() {
// 		t.Fatalf("input %q has errors: \n%s", input, errors)
// 	}

// 	return value
// }

// func testIntObject(t *testing.T, obj object.Object, expected int64) {
// 	result, ok := obj.(*object.Int)
// 	if !ok {
// 		t.Fatalf("object is not *object.Int. got=%#v, expected=%d", obj, expected)
// 	}

// 	if result.Value != expected {
// 		t.Errorf("object has wrong value. want=%d, got=%d", expected, result.Value)
// 	}
// }

// func TestEvalInteger(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected int64
// 	}{
// 		{"3", 3},
// 		{"-3", -3},
// 		{"--3", 3},
// 		{"-3 + 4", 1},
// 		{"10 - 4", 6},
// 		{"-10 * 4 - 3 + 7", -36},
// 	}

// 	for _, test := range tests {
// 		value := testEval(t, test.input)
// 		testIntObject(t, value, test.expected)
// 	}
// }

// func testRealObject(t *testing.T, obj object.Object, expected float64) {
// 	f, ok := obj.(*object.Real)
// 	if !ok {
// 		t.Errorf("object is not *object.Real. got=%#v", obj)
// 		return
// 	}

// 	if f.Value != expected {
// 		t.Errorf("object has wrong value. want=%f, got=%f", expected, f.Value)
// 	}
// }

// func TestEvalReal(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected float64
// 	}{
// 		{"3.14", 3.14},
// 		{"--3.14", 3.14},
// 		{"-3.14", -3.14},
// 		{"3 + 0.14", 3.14},
// 		{"-3 * 5 -1.59 + 6", -10.59},
// 	}

// 	for _, test := range tests {
// 		value := testEval(t, test.input)
// 		testRealObject(t, value, test.expected)
// 	}
// }
