package service

type AccountService interface {
	Transfer(
		fromID int,
		toID int,
		amount float64,
	) error
}
