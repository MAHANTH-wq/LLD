package split

import "splitwise/internal/user"

type SplitType int

const (
	EqualSplit SplitType = iota
	UnequalSplit
	PercentSplit
)

type SplitClass interface {
	validateSplit(*user.User, []*Split) bool
}

func getSplitClass(st SplitType) SplitClass {

	switch st {
	case EqualSplit:
		return &EqualSplitClass{}
	case UnequalSplit:
		return &UnEqualSplitClass{}
	case PercentSplit:
		return &PercentSplitClass{}
	default:
		return nil
	}

}

type EqualSplitClass struct {
}

func (e *EqualSplitClass) validateSplit(paidByUser *user.User, splitList []*Split) bool {

	return true
}

type UnEqualSplitClass struct {
}

func (e *UnEqualSplitClass) validateSplit(paidByUser *user.User, splitList []*Split) bool {

	return true
}

type PercentSplitClass struct {
}

func (e *PercentSplitClass) validateSplit(paidByUser *user.User, splitList []*Split) bool {

	return true
}
