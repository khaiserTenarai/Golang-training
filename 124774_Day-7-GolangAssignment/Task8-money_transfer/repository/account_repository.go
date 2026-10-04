package repository

type AccountRepository interface {
	Transfer(
		fromID int,
		toID int,
		amount float64,
	) error
}
