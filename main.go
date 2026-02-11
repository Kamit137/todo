package main
import(
	// "fmt"
	"net/http"
	"github.com/sirupsen/logrus"
	"todo/internal/autorisation"

)

func 




func main(){
	


	http.HandleFunc("/reg", autorisation.Reg)
	http.HandleFunc("/login", autorisation.Login)
	

	http.ListenAndServe(":8080", nill)
}