package shared

type SchemeBearer struct {
	Authorization string `security:"name=Authorization"`
}
