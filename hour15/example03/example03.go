package example03

func translate(s string) string {
	switch s {
	case "en-US":
		return "Hello "
	case "fr-FR":
		return "Bonjour "
	case "it-IT":
		return "Ciao "
	default:
		return "Hello "
	}
}

func Greeting(name, locale string) string {
	salutation := translate(locale)
	return (salutation + name)
}
