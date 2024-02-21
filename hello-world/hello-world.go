package helloworld

const (
	es = "Spanish"
	fr = "French"

	defaultHelloPrefix = "Hello"
	esHelloPrefix      = "Hola"
	frHelloPrefix      = "Bonjour"
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World!"
	}

	return greetingPrefix(lang) + ", " + name
}

func greetingPrefix(lang string) (prefix string) {
	// var prefix string
	switch lang {
	case es:
		prefix = esHelloPrefix
	case fr:
		prefix = frHelloPrefix
	default:
		prefix = defaultHelloPrefix
	}

	// inferred 'prefix' return as declared in func
	return
}
