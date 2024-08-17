package repository

type Repos struct {
	Log LogInterface
	Fsio FsioInterface
}

func New() Repos {
	return Repos{
		Log: &Log{},
		Fsio: &Fsio{},
	}
}

func NewMock() Repos {
	return Repos{
		Log: &LogMock{
			Out: "",
		},
		Fsio: &FsioMock{},
	}
}
