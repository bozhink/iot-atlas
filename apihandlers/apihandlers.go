package apihandlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"

	"go.mongodb.org/mongo-driver/mongo"

	"../apimodels"
	"../datamappings"
	"../dataservice"
)

// RequestDump dumps HTTP request data
func RequestDump(request *http.Request) {
	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(requestDump))
}

// GetInsertRecordEndpointHandler returns api handler for inserting requested data to MongoDB.
func GetInsertRecordEndpointHandler(client *mongo.Client) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(response)

		if request.Header.Get("Content-Type") == "application/json" {

			decoder := json.NewDecoder(request.Body)

			RequestDump(request)

			var event apimodels.EventRequestModel
			_ = decoder.Decode(&event)
			log.Println(event)

			eventDTO := datamappings.GetDTO(&event)

			if eventDTO == nil {
				log.Println("Invalid readings")
				response.WriteHeader(400)
				encoder.Encode(nil)
			} else {
				result, err := dataservice.RegisterEvent(client, eventDTO)

				if err != nil {
					response.WriteHeader(500)
					encoder.Encode(err)
				} else {
					response.WriteHeader(200)
					encoder.Encode(result)
				}
			}
		} else {
			log.Println("Unsupported Content-Type")
			response.WriteHeader(400)
			encoder.Encode(nil)
		}
	}
}
