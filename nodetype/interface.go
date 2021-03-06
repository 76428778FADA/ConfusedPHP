package nodetype

import (
	"reflect"
)

func HaveField(name string, target interface{}) (ok bool, value interface{}) {
	t := reflect.TypeOf(target)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if f.Name == name {
			v := reflect.ValueOf(target)
			val := v.Field(i).Interface()
			return true, val
		}
	}
	return
}

func NodeIsInterface(fieldName string) func(target interface{}) (bool, interface{}) {
	return func(target interface{}) (bool, interface{}) {
		return HaveField(fieldName, target)
	}
}

var IsHaveValueType = NodeIsInterface("Value")
var IsHaveDimType = NodeIsInterface("Dim")
var IsHaveExprType = NodeIsInterface("Expr")
var IsHaveLeftType = NodeIsInterface("Left")
var IsHaveRightType = NodeIsInterface("Right")
var IsHaveExpressionType = NodeIsInterface("Expression")
var IsHaveVariableType = NodeIsInterface("Variable")
var IsHaveStmtsType = NodeIsInterface("Stmts")
var IsHaveItemsType = NodeIsInterface("Items")
var IsHavePartsType = NodeIsInterface("Parts")
