package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"mdalai/mydeviceservice/home/devicestore"
	utils "mdalai/mydeviceservice/utils"
)

type deviceServer struct {
	store *devicestore.DeviceStore
}

func NewDeviceServer() *deviceServer {
	store := devicestore.New()
	return &deviceServer{store: store}
}

// renderJSON renders 'v' as JSON and writes it as a response into w.
func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (dserver *deviceServer) getDevicesHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get all devices at %s\n", req.URL.Path)

	utils.SetupCorsResponse(&w, req)

	devices := dserver.store.GetDevices()
	renderJSON(w, devices)
}

func (dserver *deviceServer) createDeviceHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling add device at %s\n", req.URL.Path)

	utils.SetupCorsResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		log.Printf("handling method:OPTIONS ...\n")
		return
	}

	log.Printf("handling method:POST ... \n")

	// get the header from the request and make sure it is application/json Content-Type.
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	// Decode the body (json) to Device object
	dec := json.NewDecoder(req.Body)
	var rt devicestore.Device
	if err := dec.Decode(&rt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Pass the decoded device object
	// to the add device method
	device := dserver.store.CreateDevice(rt.Name, rt.DeviceType, rt.Owner, rt.MacAddr, rt.IpAddr, rt.StartUseDate, rt.IsCommonlyUsed)

	renderJSON(w, device)
}

func (dserver *deviceServer) deleteDeviceHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling delete device at %s\n", req.URL.Path)

	utils.SetupCorsResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		log.Printf("handling method:OPTIONS ...\n")
		return
	}

	log.Printf("handling method:DELETE ... \n")

	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	err := dserver.store.DeleteDevice(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

}

func (dserver *deviceServer) updateDeviceHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling update device at %s\n", req.URL.Path)

	utils.SetupCorsResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		log.Printf("handling method:OPTIONS ...\n")
		return
	}

	log.Printf("handling method:PUT ... \n")

	// get the header from the request and make sure it is application/json Content-Type.
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	// Decode the body (json) to Device object
	dec := json.NewDecoder(req.Body)
	var rt devicestore.Device
	if err := dec.Decode(&rt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	rt.Id = id
	err2 := dserver.store.UpdateDevice(rt)

	if err2 != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

}

func main() {
	router := mux.NewRouter()
	router.StrictSlash(true)
	server := NewDeviceServer()

	router.HandleFunc("/devices", server.getDevicesHandler).Methods("GET")
	router.HandleFunc("/devices", server.createDeviceHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/devices/{id:[0-9]+}", server.deleteDeviceHandler).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/devices/{id:[0-9]+}", server.updateDeviceHandler).Methods("PUT", "OPTIONS")

	fmt.Println("Listening on port: " + os.Getenv("SERVERPORT"))
	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("SERVERPORT"), router))

}
