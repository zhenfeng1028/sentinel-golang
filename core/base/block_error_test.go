package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlockError(t *testing.T) {
	type args struct {
		opts []BlockErrorOption
	}
	tests := []struct {
		name string
		args args
		want *BlockError
	}{
		{
			name: "normal",
			args: args{opts: []BlockErrorOption{
				WithBlockType(BlockTypeFlow),
				WithBlockMsg("test"),
				WithRule(new(MockRule)),
				WithSnapshotValue("snapshot"),
			}},
			want: NewBlockError(WithBlockType(BlockTypeFlow),
				WithBlockMsg("test"),
				WithRule(new(MockRule)),
				WithSnapshotValue("snapshot")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBlockError(tt.args.opts...)
			assert.Equal(t, got.blockType, tt.want.blockType)
			assert.Equal(t, got.blockMsg, tt.want.blockMsg)
			assert.Equal(t, got.rule, tt.want.rule)
			assert.Equal(t, got.snapshotValue, tt.want.snapshotValue)
		})
	}
}

type MockRule struct {
	Id string `json:"id"`
}

func (m *MockRule) String() string {
	return "mock rule"
}

func (m *MockRule) ResourceName() string {
	return "mock resource"
}
