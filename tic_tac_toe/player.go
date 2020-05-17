package tic_tac_toe

type Player interface {
	GetMove(*board) (int, int, error)
	ConvertMove(int) (int, int, error)
	Mark() string
	Name() string
}
