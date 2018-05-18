package controller

// 主页
import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to home!"))
	return
}
