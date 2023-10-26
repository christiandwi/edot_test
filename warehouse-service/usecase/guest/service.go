package guest

type Service interface {
	HelloService(name string) (text map[string]interface{})
}
