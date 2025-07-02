package transaction

import "splitwise/internal/user"

type Transaction struct {
	fromUser *user.User
	toUser   *user.User
	// transaction amount should always be positive
	amount int
}

func NewTransaction(fromUser *user.User, toUser *user.User, amount int) *Transaction {

	return &Transaction{
		fromUser: fromUser,
		toUser:   toUser,
		amount:   amount,
	}
}

func (t *Transaction) GetFromUser() *user.User {
	return t.fromUser
}

func (t *Transaction) GetToUser() *user.User {
	return t.toUser
}

func (t *Transaction) GetAmount() int {
	return t.amount
}
