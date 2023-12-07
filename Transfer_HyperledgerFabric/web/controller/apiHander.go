package controller

import (
	"github.com/Transfer_HyperledgerFabric/api/plain_equipment_api"
	"github.com/Transfer_HyperledgerFabric/api/cipher_equipment_api"
	"github.com/Transfer_HyperledgerFabric/sdk_service"
	"io/ioutil"
	"net/http"
)

type Application struct {
	SdkSetup *sdk_service.ServiceSetup
}


func (app *Application) AddPlainEquipment_APP(w http.ResponseWriter, r *http.Request) {
	htmlData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		BuildErr(w, err.Error())
		return
	}

	//fmt.Println(string(htmlData))
	resp, err := plain_equipment_api.AddPlainEquipment_API(*app.SdkSetup, string(htmlData))
	if err != nil {
		BuildErr(w, "502")
		return
	}
	BuildResp(w, resp)
}

func (app *Application) FindPlainEquipment_APP(w http.ResponseWriter, r *http.Request) {
	htmlData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		BuildErr(w, err.Error())
		return
	}

	//fmt.Println(string(htmlData))
	resp, err := plain_equipment_api.FindPlainEquipment_API(*app.SdkSetup, string(htmlData))
	if err != nil {
		BuildErr(w, "501")
		return
	}
	BuildResp(w, resp)
}

func (app *Application) AddCipherEquipment_APP(w http.ResponseWriter, r *http.Request) {
	htmlData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		BuildErr(w, err.Error())
		return
	}

	//fmt.Println(string(htmlData))
	resp, err := cipher_equipment_api.AddCipherEquipment_API(*app.SdkSetup, string(htmlData))
	if err != nil {
		BuildErr(w, "502")
		return
	}
	BuildResp(w, resp)
}

func (app *Application) FindCipherEquipment_APP(w http.ResponseWriter, r *http.Request) {
	htmlData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		BuildErr(w, err.Error())
		return
	}

	//fmt.Println(string(htmlData))
	resp, err := cipher_equipment_api.FindCipherEquipment_API(*app.SdkSetup, string(htmlData))
	if err != nil {
		BuildErr(w, "501")
		return
	}
	BuildResp(w, resp)
}
