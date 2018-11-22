package hello

//ProvideHello provides hello
func ProvideHello() *Hello {
	return &Hello{"Hello World"}
}