package entity

type Achievement int

const (
	LoginAchieve Achievement = iota
	FollowerAchieve
)

func (a Achievement) Number() int {
	return int(a)
}

func (a Achievement) String() string {
	switch a {
	case LoginAchieve:
		return "LoginAchieved"
	case FollowerAchieve:
		return "FollowerAchieved"
	default:
		return "Undefined"
	}
}
