package server

func Init() {
	r := Router()
	r.Run("localhost:8080")
}