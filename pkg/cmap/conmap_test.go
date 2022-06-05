package cmap

import (
	"reflect"
	"testing"
)

func TestNewConMap(t *testing.T) {
	pairs := []struct {
		k int
		v string
	}{
		{k: 1, v: "a"},
		{k: 2, v: "b"},
		{k: 3, v: "c"},
		{k: 4, v: "d"},
	}

	type args struct {
		ktype reflect.Type
		vtype reflect.Type
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"A", args{reflect.TypeOf(pairs[0].k), reflect.TypeOf(pairs[0].v)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm, err := NewConMap(tt.args.ktype, tt.args.vtype)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// * Test Load
			cm.LoadOrStore(pairs[0].k, pairs[0].v)
			cm.LoadOrStore(pairs[1].k, pairs[1].v)
			cm.LoadOrStore(pairs[2].k, pairs[2].v)
			cm.LoadOrStore(pairs[3].k, pairs[3].v)

			// * Test Range
			cm.Range(func(k, v interface{}) bool {
				t.Logf("%4v - %v \n", k, v)
				return true
			})

			// * Test Delete
			if err := cm.Delete("12"); err != nil {
				t.Logf("ConMap.Delete() error = %v", err)
			}
		})
	}
}
