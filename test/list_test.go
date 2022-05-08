package test

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"

	"github.com/aide-cloud/base-types/types"
	"github.com/aide-cloud/base-types/util/date"
	"github.com/stretchr/testify/suite"
)

type listSuite struct {
	suite.Suite
	listCount int
}

func TestList(t *testing.T) {
	suite.Run(t, new(listSuite))
}

func (s *listSuite) SetupSuite() {
	fmt.Println("SetupSuite")
	_, file, _, _ := runtime.Caller(0)
	fmt.Println("执行文件：", file)
	s.listCount = 0
}

// 每个测试前初始化数据
func (s *listSuite) SetupTest() {
	s.listCount += 1
	fmt.Println("[" + date.MicTimeToFormatStr() + "]（第" + strconv.Itoa(s.listCount) + "条）测试用例名：[" + s.T().Name() + "]，开始执行")
}

// 每个测试结束后运行一次。
func (s *listSuite) TearDownTest() {
	fmt.Println("[" + date.MicTimeToFormatStr() + "]（第" + strconv.Itoa(s.listCount) + "条）测试用例执行完成")
}

// 在所有测试或基准测试完成后运行一次。
func (s *listSuite) TearDownSuite() {
	fmt.Println("[" + date.MicTimeToFormatStr() + "]所有测试执行完成，收尾清理数据！！！】，累计测试次数" + strconv.Itoa(s.listCount))
}

func (s *listSuite) TestNewList() {
	l := types.NewList(1, 2, 3)
	s.Equal(3, l.Len())
	s.Equal(1, l.Get(0))
	s.Equal(2, l.Get(1))
	s.Equal(3, l.Get(2))
}

func (s *listSuite) TestListLen() {
	l := types.NewList(1, 2, 3)
	l.Append(4)
	s.Equal(4, l.Len())
}

func (s *listSuite) TestListGet() {
	l := types.NewList(1, 2, 3)
	s.Equal(1, l.Get(0))
	s.Equal(2, l.Get(1))
	s.Equal(3, l.Get(2))
}

func (s *listSuite) TestListSet() {
	l := types.NewList(1, 2, 3)
	l.Set(0, 4)
	s.Equal(4, l.Get(0))
	l.Set(1, 5)
	s.Equal(5, l.Get(1))
	l.Set(2, 6)
	s.Equal(6, l.Get(2))
}

func (s *listSuite) TestListAppend() {
	l := types.NewList(1, 2, 3)
	l.Append(4, 5, 6)
	s.Equal(6, l.Len())
	s.Equal(4, l.Get(3))
	s.Equal(5, l.Get(4))
	s.Equal(6, l.Get(5))
}

func (s *listSuite) TestRemoveByIndex() {
	l := types.NewList(1, 2, 3)
	l.RemoveByIndex(1)
	s.Equal(2, l.Len())
	s.Equal(1, l.Get(0))
	s.Equal(3, l.Get(1))
}

func (s *listSuite) TestRemoveByValue() {
	l := types.NewList(1, 2, 3)
	l.RemoveByValue(2)
	s.Equal(2, l.Len())
	s.Equal(1, l.Get(0))
	s.Equal(3, l.Get(1))
}

func (s *listSuite) TestListClear() {
	l := types.NewList(1, 2, 3)
	l.Clear()
	s.Equal(0, l.Len())
}

func (s *listSuite) TestListContains() {
	l := types.NewList(1, 2, 3)
	s.True(l.Contains(1))
	s.True(l.Contains(2))
	s.True(l.Contains(3))
	s.False(l.Contains(4))
}

func (s *listSuite) TestListContainsAll() {
	l := types.NewList(1, 2, 3)
	s.True(l.ContainsAll(1, 2, 3))
	s.False(l.ContainsAll(1, 2, 4))
	s.False(l.ContainsAll(1, 4, 3))
	s.False(l.ContainsAll(4, 2, 3))
}

func (s *listSuite) TestListContainsAny() {
	l := types.NewList(1, 2, 3)
	s.True(l.ContainsAny(1))
	s.True(l.ContainsAny(1, 2))
	s.True(l.ContainsAny(2, 3))
	s.False(l.ContainsAny(4, 5, 6))
}

func (s *listSuite) TestListContainsNone() {
	l := types.NewList(1, 2, 3)
	s.True(l.ContainsNone(0))
	s.True(l.ContainsNone(4))
	s.False(l.ContainsNone(1))
}

func (s *listSuite) TestListIndex() {
	l := types.NewList(1, 2, 3)
	s.Equal(0, l.Index(1))
	s.Equal(1, l.Index(2))
	s.Equal(2, l.Index(3))
	s.Equal(-1, l.Index(4))
}

func (s *listSuite) TestListUnique() {
	l := types.NewList(1, 2, 3)
	l.Append(1, 2, 3)
	s.Equal(1, l.Get(0))
	s.Equal(2, l.Get(1))
	s.Equal(3, l.Get(2))
	s.Equal(6, l.Len())
	l.Unique()
	s.Equal(3, l.Len())
	s.Equal(nil, l.Get(3))
	s.Equal(nil, l.Get(4))
}

func (s *listSuite) TestLisLast() {
	l := types.NewList(1, 2, 3)
	s.Equal(3, l.Last())
}

func (s *listSuite) TestListFirst() {
	l := types.NewList(1, 2, 3)
	s.Equal(1, l.First())
}

func (s *listSuite) TestListFind() {
	l := types.NewList(1, 2, 3)
	s.Equal(0, l.Find(1))
	s.Equal(1, l.Find(2))
	s.Equal(2, l.Find(3))
	s.Equal(-1, l.Find(4))
}

func (s *listSuite) TestListIndexes() {
	l := types.NewList(1, 2, 3)
	s.Equal([]int{0, 1, 2}, l.Indexes(1, 2, 3))
}

func (s *listSuite) TestListSort() {
	l := types.NewList(3, 2, 1)
	l.Sort(func(i, j int) bool {
		return l.Get(i).(int) < l.Get(j).(int)
	})
	s.Equal(1, l.Get(0))
	s.Equal(2, l.Get(1))
	s.Equal(3, l.Get(2))
}
