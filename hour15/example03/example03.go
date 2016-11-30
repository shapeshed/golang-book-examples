package example03

func translate(s string) string {
	switch s {
	case "en-US":
		return "Hello "
	case "fr":
		return "Bonjour "
	default:
		return "Hello "
	}
}

func Greeting(name, locale string) string {
	salutation := translate(locale)
	//return ("Hello" + s)
	return (salutation + name)
}
