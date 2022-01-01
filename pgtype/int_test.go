// Do not edit. Generated from pgtype/int_test.go.erb
package pgtype_test

import (
	"math"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
)

func TestInt2Codec(t *testing.T) {
	testPgxCodec(t, "int2", []PgxTranscodeTestCase{
		{int8(1), new(int16), isExpectedEq(int16(1))},
		{int16(1), new(int16), isExpectedEq(int16(1))},
		{int32(1), new(int16), isExpectedEq(int16(1))},
		{int64(1), new(int16), isExpectedEq(int16(1))},
		{uint8(1), new(int16), isExpectedEq(int16(1))},
		{uint16(1), new(int16), isExpectedEq(int16(1))},
		{uint32(1), new(int16), isExpectedEq(int16(1))},
		{uint64(1), new(int16), isExpectedEq(int16(1))},
		{int(1), new(int16), isExpectedEq(int16(1))},
		{uint(1), new(int16), isExpectedEq(int16(1))},
		{pgtype.Int2{Int: 1, Valid: true}, new(int16), isExpectedEq(int16(1))},
		{1, new(int8), isExpectedEq(int8(1))},
		{1, new(int16), isExpectedEq(int16(1))},
		{1, new(int32), isExpectedEq(int32(1))},
		{1, new(int64), isExpectedEq(int64(1))},
		{1, new(uint8), isExpectedEq(uint8(1))},
		{1, new(uint16), isExpectedEq(uint16(1))},
		{1, new(uint32), isExpectedEq(uint32(1))},
		{1, new(uint64), isExpectedEq(uint64(1))},
		{1, new(int), isExpectedEq(int(1))},
		{1, new(uint), isExpectedEq(uint(1))},
		{math.MinInt16, new(int16), isExpectedEq(int16(math.MinInt16))},
		{-1, new(int16), isExpectedEq(int16(-1))},
		{0, new(int16), isExpectedEq(int16(0))},
		{1, new(int16), isExpectedEq(int16(1))},
		{math.MaxInt16, new(int16), isExpectedEq(int16(math.MaxInt16))},
		{1, new(pgtype.Int2), isExpectedEq(pgtype.Int2{Int: 1, Valid: true})},
		{pgtype.Int2{}, new(pgtype.Int2), isExpectedEq(pgtype.Int2{})},
		{nil, new(*int16), isExpectedEq((*int16)(nil))},
	})
}