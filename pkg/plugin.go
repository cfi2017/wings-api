package pkg

type Plugin interface {
	Load(api Api) error
}
