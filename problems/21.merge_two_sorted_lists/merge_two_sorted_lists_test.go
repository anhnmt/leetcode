package merge_two_sorted_lists

import (
	"reflect"
	"testing"

	"problems/structures"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *structures.ListNode
		list2 *structures.ListNode
	}

	// exam 1
	ex1List1 := &structures.ListNode{
		Val: 1,
		Next: &structures.ListNode{
			Val: 2,
			Next: &structures.ListNode{
				Val: 4,
			},
		},
	}

	ex1List2 := &structures.ListNode{
		Val: 1,
		Next: &structures.ListNode{
			Val: 3,
			Next: &structures.ListNode{
				Val: 4,
			},
		},
	}

	ex1Output := &structures.ListNode{
		Val: 1,
		Next: &structures.ListNode{
			Val: 1,
			Next: &structures.ListNode{
				Val: 2,
				Next: &structures.ListNode{
					Val: 3,
					Next: &structures.ListNode{
						Val: 4,
						Next: &structures.ListNode{
							Val: 4,
						},
					},
				},
			},
		},
	}

	tests := []struct {
		name string
		args args
		want *structures.ListNode
	}{
		{
			name: "exam 1",
			args: args{
				list1: ex1List1,
				list2: ex1List2,
			},
			want: ex1Output,
		},
		{
			name: "exam 2",
			args: args{
				list1: &structures.ListNode{},
				list2: &structures.ListNode{},
			},
			want: &structures.ListNode{
				Next: &structures.ListNode{},
			},
		},
		{
			name: "exam 3",
			args: args{
				list1: &structures.ListNode{},
				list2: &structures.ListNode{
					Val: 0,
				},
			},
			want: &structures.ListNode{
				Val:  0,
				Next: &structures.ListNode{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
