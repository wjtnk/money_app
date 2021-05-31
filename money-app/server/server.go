package server

// Init is initialize server
func Init() {
	r := router()
	r.Run(":8080")
}
