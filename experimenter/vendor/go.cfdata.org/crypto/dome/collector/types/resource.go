package types

type Schema[C any] interface {
	Unit() C
}

type Resource[A any, T any] struct {
	Name   string
	Config *Config
	Params Conf
	Handle Handle[T]
	Schema Schema[A]
}

type Handle[T any] interface {
	Client[T]
	Socket
}

type Client[T any] interface {
	Init(*Config) (bool, error)
}

type Socket interface {
	Read(string, *Context) func(Query) ([]byte, error)
	Write(string, *Context) (bool, error)
}

func (r *Resource[A, B]) Initialize() (bool, error) {
	r.Handle.Init(r.Config)
	return true, nil
}

// changelog := core.Resource[types.ChangelogSchema, types.Changelog]{Name: "changelog", Schema: types.ChangelogSchema{}, Config: ctx.Config, Handle: &types.Changelog{}}
func MakeResource[A any, B any](n string, c *Config, p Conf, a Schema[A], b Handle[B]) Resource[A, B] {
	return Resource[A, B]{
		Name:   n,
		Config: c,
		Params: p,
		Schema: a,
		Handle: b,
	}
}
