package repository

type Repos struct {
	Fsio FsioInterface
}

func New() Repos {
	return Repos{
		Fsio: &Fsio{},
	}
}
