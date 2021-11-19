package stack

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	tests := []struct {
		name string
		want Stack
	}{
		{
			name: "Test new stack creation",
			want: Stack{arr: make([]interface{}, 0), lastIndex: -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	type fields struct {
		arr       []interface{}
		lastIndex int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test IsEmpty false",
			fields: fields{
				arr:       []interface{}{1},
				lastIndex: 0,
			},
			want: false,
		},
		{
			name: "Test IsEmpty true",
			fields: fields{
				arr:       []interface{}{},
				lastIndex: -1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				arr:       tt.fields.arr,
				lastIndex: tt.fields.lastIndex,
			}
			if got := s.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type fields struct {
		arr       []interface{}
		lastIndex int
	}
	type args struct {
		a interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test Push(something) and Peek() == something",
			fields: fields{
				arr:       []interface{}{},
				lastIndex: -1,
			},
			args: args{
				a: "something",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				arr:       tt.fields.arr,
				lastIndex: tt.fields.lastIndex,
			}
			s.Push(tt.args.a)
			if s.Peek() != tt.args.a {
				t.Errorf("When Push(%v), Peek() got = %v, want %v", tt.args.a, s.Peek(), tt.args.a)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	type fields struct {
		arr       []interface{}
		lastIndex int
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "Test Peek Value",
			fields: fields{
				arr:       []interface{}{1, 2},
				lastIndex: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				arr:       tt.fields.arr,
				lastIndex: tt.fields.lastIndex,
			}
			if got := s.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peek() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Run("Test Panic on Stack Empty", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Code didn't panic!")
			}
		}()
		s := &Stack{
			arr:       []interface{}{},
			lastIndex: -1,
		}
		s.Peek()
	})
}

func TestStack_Pop(t *testing.T) {
	type fields struct {
		Arr       []interface{}
		LastIndex int
	}
	tests := []struct {
		name          string
		fields        fields
		want          interface{}
		wantArr       []interface{}
		wantLastIndex int
	}{
		{
			name: "Test Pop() with value inside",
			fields: fields{
				Arr:       []interface{}{1, 2},
				LastIndex: 1,
			},
			want:          2,
			wantArr:       []interface{}{1},
			wantLastIndex: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				arr:       tt.fields.Arr,
				lastIndex: tt.fields.LastIndex,
			}
			if got := s.Pop(); !reflect.DeepEqual(got, tt.want) || !reflect.DeepEqual(tt.wantArr, s.arr) || !reflect.DeepEqual(tt.wantLastIndex, s.lastIndex) {
				t.Errorf("Pop() = %v, want %v; arr = %v, want %v; lastIndex = %v, want = %v", got, tt.want, s.arr, tt.wantArr, s.lastIndex, tt.wantLastIndex)
			}
		})
	}
	t.Run("Test Panic on Stack Empty", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Code didn't panic!")
			}
		}()
		s := &Stack{
			arr:       []interface{}{},
			lastIndex: -1,
		}
		s.Pop()
	})
}
