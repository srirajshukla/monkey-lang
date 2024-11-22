package evaluator

import (
	"bytes"
	"fmt"
	"gks/monkey_intp/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			if length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if len(arr.Elements) > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to `push` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"put": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 3 {
				return newError("wrong number of arguments. got=%d, want=3", len(args))
			}

			if args[0].Type() != object.HASH_OBJ {
				return newError("first argument to `put` must be Hash Map, got %s", args[0].Type())
			}

			// Get the original hashmap
			hashmap := args[0].(*object.Hash)

			// Get the key to insert and check if it's correct
			keyObj, ok := args[1].(object.Hashable)

			if !ok {
				return newError("unusable as hash key: %s", args[1].Type())
			}

			// create a new map and copy all the original values
			newMap := make(map[object.HashKey]object.HashPair)

			for k, v := range hashmap.Pairs {
				newMap[k] = v
			}

			// assign the new value to new map
			value := object.HashPair{Key: args[1], Value: args[2]}
			newMap[keyObj.HashKey()] = value

			return &object.Hash{Pairs: newMap}
		},
	},
	"remove": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			switch obj := args[0].(type) {
			case *object.Array:
				length := len(obj.Elements)

				idxToRemove, ok := args[1].(*object.Integer)

				if !ok {
					return newError("index to remove in array must be integer, got %s", args[1].Type())
				}

				if int(idxToRemove.Value) >= length || idxToRemove.Value < 0 {
					return newError("invalid index to remove from array")	
				}

				newArray := make([]object.Object, 0)
				newArray = append(newArray, obj.Elements[:idxToRemove.Value]...)
				newArray = append(newArray, obj.Elements[idxToRemove.Value+1:]...)

				return &object.Array{Elements: newArray}

			case *object.Hash:
				keyToRemove, ok := args[1].(object.Hashable)

				if !ok {
					return newError("key to remove in hashmap must be hashable, got = %s", args[1].Type())
				}
				// create a new map and copy all the original values
				newMap := make(map[object.HashKey]object.HashPair)

				for k, v := range obj.Pairs {
					newMap[k] = v
				}

				// delete the value with key equals keyToRemove
				delete(newMap, keyToRemove.HashKey())

				return &object.Hash{Pairs: newMap}
			default:
				return newError("first argument to `push` must be ARRAY or HASH MAP, got %s", args[0].Type())
			}
		},
	},
	"puts": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {

			for _, arg := range args {
				fmt.Println(arg.Insepct())
			}

			return NULL
		},
	},
	"join": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			var out bytes.Buffer

			for _, arg := range args {
				out.WriteString(arg.Insepct())
			}

			return &object.String{Value: out.String()}
		},
	},
}
