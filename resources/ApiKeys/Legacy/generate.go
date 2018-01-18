package Legacy

import (
	"github.com/autopilothq/lg"

	ct "github.com/autopilothq/banks/contract/types"
	proto "github.com/autopilothq/banks/protocol"
	"github.com/example/taranaki/resources/ApiKeys/protocol"
)

// Generate is something that you should document
func Generate(request *proto.Request, log lg.Log, aux ct.Auxiliary) *proto.Response {
	generateRequest := request.Body.(*protocol.GenerateRequest)

	datascope := request.DatascopeID.String()

	log.Debugf("request %v\n", generateRequest)
	role := generateRequest.Role

	response := request.MakeOkResponse()
	body := response.Body.(*protocol.GenerateResponse)
	body.ApiKey = &protocol.ApiKey{
		Role:    role,
		Token:   "placeholder-token",
		Enabled: true,
		UserID:  datascope,
	}
	return response
}
