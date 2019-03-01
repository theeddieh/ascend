package db

import (
	"testing"
)

func TestReadAndWrite(t *testing.T) {
	tests := []struct {
		writeKey, writeVal string
		readKey, readVal   string
		expected           error
	}{
		{"key-0", "value-0", "key-0", "value-0", nil},
		{"key-0", "", "key-1", "", ErrKeyNonexistant},
		{"key-1", "value-1", "key-1", "value-1", nil},
		{"key-2", "value-2", "key-2", "value-2", nil},
	}

	d := New()
	for i, tt := range tests {

		d.Write(tt.writeKey, tt.writeVal)
		val, err := d.Read(tt.readKey)
		if err != tt.expected {
			t.Errorf("read error '%v', expected '%v' for test case [%d] %v", err, tt.expected, i, tt)
			continue
		}
		if val != tt.readVal {
			t.Errorf("read value '%s', expected '%s' for test case [%d] %v", val, tt.readVal, i, tt)
		}
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		key, val string
		expected error
	}{
		{"key-0", "value-0", ErrKeyDeleted},
	}

	d := New()
	for i, tt := range tests {

		d.Write(tt.key, tt.val)
		_, err := d.Read(tt.key)
		if err != nil {
			t.Errorf("read error '%v'", err)
			continue
		}

		d.Delete(tt.key)
		_, err = d.Read(tt.key)
		if err != tt.expected {
			t.Errorf("read error '%v', expected '%v' for test case [%d] %v", err, tt.expected, i, tt)
		}
	}
}

func TestRollback(t *testing.T) {
	tests := []struct {
		key, val string
		expected error
	}{
		{"key-0", "value-0", ErrKeyMissing},
	}

	d := New()
	for i, tt := range tests {

		d.Write(tt.key, tt.val)
		_, err := d.Read(tt.key)
		if err != nil {
			t.Errorf("read error '%v'", err)
			continue
		}

		d.Rollback()
		_, err = d.Read(tt.key)
		if err != tt.expected {
			t.Errorf("read error '%v', expected '%v' for test case [%d] %v", err, tt.expected, i, tt)
		}
	}
}
