package encoding

import (
	"testing"
)

func TestU1Byte(t *testing.T) {
	buf := []byte{0x01}
	res, step, err := Upvarint(buf, 0)
	var expected uint64 = 1
	if err != nil {
		t.Errorf("err should nil, actual = %v", err)
	}
	if res != expected {
		t.Errorf("value should %v, actual = %v", expected, res)
	}
	if step != 1 {
		t.Errorf("step should %v, actual = %v", 1, step)
	}
}

func TestUWrongByte(t *testing.T) {
	buf := []byte{0x81, 0x82}
	_, _, err := Upvarint(buf, 0)
	if err == nil {
		t.Errorf("err should not nil")
	}
	if err.Error() != "malformed buffer" {
		t.Errorf("err should not nil, actual err = %v", err)
	}
}

func TestU2Bytes(t *testing.T) {
	buf := []byte{0x81, 0x02}
	res, step, err := Upvarint(buf, 0)
	var expected uint64 = 130
	if err != nil {
		t.Errorf("err should nil, actual = %v", err)
	}

	if res != expected {
		t.Errorf("value should %v, actual = %v", expected, res)
	}
	if step != 2 {
		t.Errorf("step should %v, actual = %v", 2, step)
	}
}

func TestU2BytesSkip(t *testing.T) {
	buf := []byte{0x81, 0x82, 0x81, 0x02}
	res, step, err := Upvarint(buf, 2)
	var expected uint64 = 130
	if err != nil {
		t.Errorf("err should nil, actual = %v", err)
	}

	if res != expected {
		t.Errorf("value should %v, actual = %v", expected, res)
	}
	if step != 2 {
		t.Errorf("step should %v, actual = %v", 2, step)
	}
}

func Test1ByteNagitive(t *testing.T) {
	buf := []byte{0x7B}
	res, step, err := Pvarint(buf, 0)
	var expected int64 = -5
	if err != nil {
		t.Errorf("err should nil, actual = %v", err)
	}
	if res != expected {
		t.Errorf("value should %v, actual = %v", expected, res)
	}
	if step != 1 {
		t.Errorf("step should %v, actual = %v", 1, step)
	}
}

func Test1BytePositive(t *testing.T) {
	buf := []byte{0x3F}
	res, step, err := Pvarint(buf, 0)
	var expected int64 = 63
	if err != nil {
		t.Errorf("err should nil, actual = %v", err)
	}
	if res != expected {
		t.Errorf("value should %v, actual = %v", expected, res)
	}
	if step != 1 {
		t.Errorf("step should %v, actual = %v", 1, step)
	}
}

func Test2BytesNagitive(t *testing.T) {
	buf := []byte{0xFF, 0x3F}
	res, step, err := Pvarint(buf, 0)
	var expected int64 = -65
	if err != nil {
		t.Errorf("err should nil, actual = %v", err)
	}
	if res != expected {
		t.Errorf("value should %v, actual = %v", expected, res)
	}
	if step != 2 {
		t.Errorf("step should %v, actual = %v", 2, step)
	}
}

func Test2BytesPositive2(t *testing.T) {
	buf := []byte{0x80, 0x7F}
	res, step, err := Pvarint(buf, 0)
	var expected int64 = 127
	if err != nil {
		t.Errorf("err should nil, actual = %v", err)
	}
	if res != expected {
		t.Errorf("value should %v, actual = %v", expected, res)
	}
	if step != 2 {
		t.Errorf("step should %v, actual = %v", 2, step)
	}
}

func Test2BytesPositive(t *testing.T) {
	buf := []byte{0xFF, 0x81, 0x7F}
	res, step, err := Pvarint(buf, 1)
	var expected int64 = 255
	if err != nil {
		t.Errorf("err should nil, actual = %v", err)
	}
	if res != expected {
		t.Errorf("value should %v, actual = %v", expected, res)
	}
	if step != 2 {
		t.Errorf("step should %v, actual = %v", 2, step)
	}
}