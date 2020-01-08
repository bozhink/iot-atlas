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
		RequestDump(request)
		
		response.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(response)

		if request.Header.Get("Content-Type") == "application/json" {

			var event apimodels.EventRequestModel
			decodeError := json.NewDecoder(request.Body).Decode(&event)
			if decodeError != nil {
				log.Println("Request decode error", decodeError)
				response.WriteHeader(400)
				encoder.Encode(nil)
			}

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
					response.WriteHeader(201)
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
