// Credits to following links
// https://flaviocopes.com/golang-enable-cors/ 
// https://perennialsky.medium.com/handle-cors-in-golang-7c5c3902dc08 

package utils

import "net/http"


func SetupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
 }

