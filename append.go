package jsonvalue

// Append type is for InTheEnd() or InTheBeginning() function. Please refer to related functions.
//
// Shoud ONLY be generated by V.Append() function
//
// Append 类型是用于 InTheEnd() 和 InTheBeginning() 函数的。使用者可以不用关注这个类型。并且这个类型只应当由 V.Append() 产生。
type Append struct {
	v *V
	c *V // child
}

// Append starts appending a child JSON value to a JSON array.
//
// Append 开始将一个 JSON 值添加到一个数组中。需结合 InTheEnd() 和 InTheBeginning() 函数使用。
func (v *V) Append(child *V) *Append {
	if nil == child {
		child = NewNull()
	}
	return &Append{
		v: v,
		c: child,
	}
}

// AppendString is equivalent to Append(jsonvalue.NewString(s))
//
// AppendString 等价于 Append(jsonvalue.NewString(s))
func (v *V) AppendString(s string) *Append {
	return v.Append(NewString(s))
}

// AppendBytes is equivalent to Append(jsonvalue.NewBytes(b))
//
// AppendBytes 等价于 Append(jsonvalue.NewBytes(b))
func (v *V) AppendBytes(b []byte) *Append {
	return v.Append(NewBytes(b))
}

// AppendBool is equivalent to Append(jsonvalue.NewBool(b))
//
// AppendBool 等价于 Append(jsonvalue.NewBool(b))
func (v *V) AppendBool(b bool) *Append {
	return v.Append(NewBool(b))
}

// AppendInt is equivalent to Append(jsonvalue.NewInt(b))
//
// AppendInt 等价于 Append(jsonvalue.NewInt(b))
func (v *V) AppendInt(i int) *Append {
	return v.Append(NewInt(i))
}

// AppendInt64 is equivalent to Append(jsonvalue.NewInt64(b))
//
// AppendInt64 等价于 Append(jsonvalue.NewInt64(b))
func (v *V) AppendInt64(i int64) *Append {
	return v.Append(NewInt64(i))
}

// AppendInt32 is equivalent to Append(jsonvalue.NewInt32(b))
//
// AppendInt32 等价于 Append(jsonvalue.NewInt32(b))
func (v *V) AppendInt32(i int32) *Append {
	return v.Append(NewInt32(i))
}

// AppendUint is equivalent to Append(jsonvalue.NewUint(b))
//
// AppendUint 等价于 Append(jsonvalue.NewUint(b))
func (v *V) AppendUint(u uint) *Append {
	return v.Append(NewUint(u))
}

// AppendUint64 is equivalent to Append(jsonvalue.NewUint64(b))
//
// AppendUint64 等价于 Append(jsonvalue.NewUint64(b))
func (v *V) AppendUint64(u uint64) *Append {
	return v.Append(NewUint64(u))
}

// AppendUint32 is equivalent to Append(jsonvalue.NewUint32(b))
//
// AppendUint32 等价于 Append(jsonvalue.NewUint32(b))
func (v *V) AppendUint32(u uint32) *Append {
	return v.Append(NewUint32(u))
}

// AppendFloat64 is equivalent to Append(jsonvalue.NewFloat64(b))
//
// AppendUint32 等价于 Append(jsonvalue.NewUint32(b))
func (v *V) AppendFloat64(f float64) *Append {
	return v.Append(NewFloat64(f))
}

// AppendFloat32 is equivalent to Append(jsonvalue.NewFloat32(b))
//
// AppendFloat32 等价于 Append(jsonvalue.NewFloat32(b))
func (v *V) AppendFloat32(f float32) *Append {
	return v.Append(NewFloat32(f))
}

// AppendNull is equivalent to Append(jsonvalue.NewNull())
//
// AppendNull 等价于 Append(jsonvalue.NewNull())
func (v *V) AppendNull() *Append {
	return v.Append(NewNull())
}

// AppendObject is equivalent to Append(jsonvalue.NewObject())
//
// AppendObject 等价于 Append(jsonvalue.NewObject())
func (v *V) AppendObject() *Append {
	return v.Append(NewObject())
}

// AppendArray is equivalent to Append(jsonvalue.NewArray())
//
// AppendArray 等价于 Append(jsonvalue.NewArray())
func (v *V) AppendArray() *Append {
	return v.Append(NewArray())
}

// InTheBeginning completes the following operation of Append().
//
// InTheBeginning 函数将 Append 函数指定的 JSON 值，添加到参数指定的数组的最前端
func (apd *Append) InTheBeginning(params ...interface{}) (*V, error) {
	v := apd.v
	c := apd.c
	if nil == v || v.valueType == NotExist {
		return &V{}, ErrValueUninitialized
	}

	// this is the last iteration
	paramCount := len(params)
	if paramCount == 0 {
		if v.valueType != Array {
			return &V{}, ErrNotArrayValue
		}

		v.appendToArr(c)
		return c, nil
	}

	// this is not the last iterarion
	child, err := v.GetArray(params[0], params[1:paramCount]...)
	if err != nil {
		return &V{}, err
	}

	if child.Len() == 0 {
		child.appendToArr(c)
	} else {
		child.insertToArr(0, c)
	}
	return c, nil
}

// InTheEnd completes the following operation of Append().
//
// InTheEnd 函数将 Append 函数指定的 JSON 值，添加到参数指定的数组的最后面
func (apd *Append) InTheEnd(params ...interface{}) (*V, error) {
	v := apd.v
	c := apd.c
	if v.valueType == NotExist {
		return &V{}, ErrValueUninitialized
	}

	// this is the last iteration
	paramCount := len(params)
	if paramCount == 0 {
		if v.valueType != Array {
			return &V{}, ErrNotArrayValue
		}

		v.appendToArr(c)
		return c, nil
	}

	// this is not the last iterarion
	child, err := v.GetArray(params[0], params[1:paramCount]...)
	if err != nil {
		return &V{}, err
	}

	child.appendToArr(c)
	return c, nil
}
