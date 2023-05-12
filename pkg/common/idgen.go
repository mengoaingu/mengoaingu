package common

import (
	"context"
	"strconv"
	"time"

	"github.com/sony/sonyflake"
)

var _ IDGenerator = &LocalSonyFlakeIdGenerator{}

type IDGenerator interface {
	NextUint64(ctx context.Context) (uint64, error)
	NextID(ctx context.Context) (string, error)
}

type LocalSonyFlakeIdGenerator struct {
	s *sonyflake.Sonyflake
}

func (l *LocalSonyFlakeIdGenerator) NextUint64(ctx context.Context) (uint64, error) {
	return l.s.NextID()
}

func NewLocalSonyFlakeIdGenerator() *LocalSonyFlakeIdGenerator {
	t, _ := time.Parse(time.RFC3339, "2020-12-31T23:59:59.999Z7:00")
	return &LocalSonyFlakeIdGenerator{
		s: sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime: t,
		}),
	}
}

func (l *LocalSonyFlakeIdGenerator) NextID(ctx context.Context) (string, error) {
	id, err := l.s.NextID()
	return strconv.FormatUint(id, 36), err
}
