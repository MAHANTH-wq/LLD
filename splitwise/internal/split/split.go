package split

import "splitwise/internal/user"

type Split struct {
	//here user is the person who should give back the amount
	user *user.User

	//Here amountToGive is the total amount that the user has to return to the paid user.
	amountToGive int
}

func CreateNewSplit(u *user.User, amountToGive int) *Split {

	return &Split{
		user:         u,
		amountToGive: amountToGive,
	}

}

func (s *Split) GetSplitUserId() int {
	return s.user.GetUserId()
}

func (s *Split) GetSplitUser() *user.User {
	return s.user
}

func (s *Split) GetSplitAmount() int {
	return s.amountToGive
}
