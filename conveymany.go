package conveymany

import (
	"fmt"
	"reflect"

	"github.com/smartystreets/goconvey/convey"
)

// TestTableData hold the data for a test table
type TestTableData struct {
	entries []interface{}
}

// TestTable creates a test table object
func TestTable(entries ...interface{}) *TestTableData {

	// if raw values are passed instead of a slice of TestValues
	// then wrap them in an []interface{}
	parsedEntries := []interface{}{}
	for _, entry := range entries {

		value := reflect.ValueOf(entry)
		if value.Kind() == reflect.Slice {
			parsedEntries = append(parsedEntries, entry)
		} else {
			parsedEntries = append(parsedEntries, []interface{}{entry})
		}
	}

	return &TestTableData{
		entries: parsedEntries,
	}
}

// TestValues create a test value object
func TestValues(values ...interface{}) []interface{} {
	return values
}

// ConveyMany defines a Convey test case with multiple values defined in a TestTable
func ConveyMany(situation string, items ...interface{}) {
	// fmt.Println("1111 >>>>>>>>>>>\n>")
	// fmt.Printf("ITEMS: %+v\n", items)
	// fmt.Printf("ITEMS[0]: %+v\n", items[0])
	test, items := parseGoTest(items)
	table, items := parseTestTable(items)
	action, items := parseAction(items)

	// fmt.Printf("TEST: %+v >>> \n", test)
	// fmt.Printf("ACTION: %+v >>> \n", &action)
	// fmt.Println("2222222 >>>>>>>>>>>\n>")

	if test != nil {
		// fmt.Printf("TEST: %+v >>> \n", test)
		items = append([]interface{}{test}, items...)
	}
	if table == nil {
		//items = append([]interface{}{situation}, items...)
		//convey.Convey(items...)
		panic("ConveyMany called without TestTable")
	}
	// fmt.Println("33333333 >>>>>>>>>>>\n>")

	f := func() {
		for _, entry := range table.entries {
			//				situationN := fmt.Sprintf("%s [%d]", situation, i)
			callTestFunction(entry, action)
		}
	}

	values := append([]interface{}{situation}, items...)
	values = append(values, f)
	// fmt.Printf("VALUES: %#v", values)
	convey.Convey(values...)
}

func callTestFunction(entry interface{}, action reflect.Value) {
	// fmt.Printf("CALL TEST FUNCTION: %#v, %#v", entry, action)

	actionValue := reflect.ValueOf(action)
	actionType := actionValue.Type()

	testValues, ok := entry.([]interface{})
	if !ok {
		panic("Parse values error")
	}

	params := make([]reflect.Value, len(testValues))
	for i, param := range testValues {
		if param == nil {
			inType := actionType.In(i)
			params[i] = reflect.Zero(inType)
		} else {
			params[i] = reflect.ValueOf(param)
		}
	}
	action.Call(params)

	//	convey.So(1, convey.ShouldEqual, 1)
}

func item(items []interface{}) interface{} {
	if len(items) == 0 {
		panic("parseError")
	}
	return items[0]
}

func parseTestTable(items []interface{}) (*TestTableData, []interface{}) {
	if table, parsed := item(items).(*TestTableData); parsed {
		return table, items[1:]
	}
	return nil, items
}

func parseGoTest(items []interface{}) (t, []interface{}) {
	if test, parsed := item(items).(t); parsed {
		return test, items[1:]
	}
	return nil, items
}

func parseAction(items []interface{}) (reflect.Value, []interface{}) {

	action := item(items)
	actionValue := reflect.ValueOf(action)
	if actionValue.Kind() != reflect.Func {
		panic(fmt.Sprintf("Expected a function, got %#v", action))
	}
	return actionValue, items[1:]
}

// This interface allows us to pass the *testing.T struct
// throughout the internals of this package without ever
// having to import the "testing" package.
type t interface {
	Fail()
}
