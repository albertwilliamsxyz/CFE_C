package server

const (
	CFEProtocol = "https"
	CFEHostName = "cfecontigo.cfe.mx"
	CFEPort     = "8765"

	CFEOperationalUserId = "2941713"
	CFEOperationalUserPassword = "Nosenada002$"

	CFEHTTPGrantType = "password"
	CFEHTTPScope     = "trust"
)

const (
	PathNull        = iota
	PathGetToken    = iota
	PathGetInvoices = iota
)

var Path = map[int]string{
	PathNull:        "",
	PathGetToken:    "uaa/oauth/token",
	PathGetInvoices: "recibos/obtener",
}