package web

import (
	"fmt"
	"github.com/Transfer_HyperledgerFabric/web/controller"
	"net/http"
)

func WebStart(app controller.Application) {
	mux := http.NewServeMux()

	//TODO

	mux.HandleFunc("/api/plain_equipment_api/AddPlainEquipment_API", app.AddPlainEquipment_APP)

	mux.HandleFunc("/api/plain_equipment_api/FindPlainEquipment_API", app.FindPlainEquipment_APP)

	mux.HandleFunc("/api/cipher_equipment_api/AddCipherEquipment_API", app.AddCipherEquipment_APP)

	mux.HandleFunc("/api/plain_equipment_api/FindCipherEquipment_API", app.FindCipherEquipment_APP)





	fmt.Println("启动Web服务, 监听端口号为:(http://localhost:7000/)")
	//fmt.Println("启动Web服务, 监听端口号为:(http://localhost:" + os.Getenv("LISTENING_PORT") + "/)")
	err := http.ListenAndServe(":7000", mux)
	//err := http.ListenAndServe(":"+os.Getenv("LISTENING_PORT"), mux)
	if err != nil {

		fmt.Printf("Web服务启动失败: %v", err)
	}
}
