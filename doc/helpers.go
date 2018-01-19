package doc

// Helpers provides additional information for the different entities of documentation.
type helpers struct {
	parameters map[string]Parameter // keys of the map represents different types of URLs
}

var h = &helpers{
	parameters: make(map[string]Parameter),
}

func SetHelperParameter(url string, p Parameter) {
	h.parameters[url] = p
}
