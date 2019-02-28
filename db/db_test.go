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
		{"key-0", "", "key-0", "", nil},
		{"key-0", "", "key-1", "", ErrMissingKey},
		{"key-1", "", "key-1", "", nil},
		{"key-1", "value-1", "key-1", "value-1", nil},
		{"", "", "", "", nil},
		{"key-2", "value-2", "key-2", "value-2", nil},
	}

	d := New()
	for _, tt := range tests {

		d.Write(tt.writeKey, tt.writeVal)
		val, err := d.Read(tt.readKey)
		if err != tt.expected {
			t.Errorf("read error '%v', expected '%v'", err, tt.expected)
			continue
		}
		if val != tt.readVal {
			t.Errorf("read value '%s', expected '%s'", val, tt.readVal)
		}
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		key, val string
		expected error
	}{
		{"key-0", "value-0", ErrMissingKey},
	}

	d := New()
	for _, tt := range tests {

		d.Write(tt.key, tt.val)
		_, err := d.Read(tt.key)
		if err != nil {
			t.Errorf("read error '%v'", err)
			continue
		}

		d.Delete(tt.key)
		_, err = d.Read(tt.key)
		if err != tt.expected {
			t.Errorf("read error '%v', expected '%v'", err, tt.expected)
		}
	}
}
