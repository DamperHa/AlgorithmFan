package link

import (
	"reflect"
	"testing"
)

func TestCreateListBySlice(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateListBySlice(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateListBySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
