package server

func Init() {
	r := Router() // Imported from same package & from file "routes.go"
	r.Run("localhost:8080")
}