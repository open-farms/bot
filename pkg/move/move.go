package move

const (
	Stop Direction = iota + 1
	Backward
	Forward
	Left
	Right
)

type Direction int

func (d Direction) String() string {
	return [...]string{"stop", "backward", "forward", "left", "right"}[d-1]
}

type Mover interface {
	Move(d Direction)
}

type Controller struct {
	Mover
}

func New(mover Mover) *Controller {
	return &Controller{
		mover,
	}
}
