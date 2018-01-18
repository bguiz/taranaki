package Legacy

import (
	"github.com/autopilothq/lg"

	ct "github.com/autopilothq/banks/contract/types"
	proto "github.com/autopilothq/banks/protocol"
	"github.com/example/taranaki/resources/ApiKeys/protocol"
)

// Get is something that you should document
func Get(request *proto.Request, log lg.Log, aux ct.Auxiliary) *proto.Response {
	getRequest := request.Body.(*protocol.GetRequest)

	datascope := request.DatascopeID.String()

	log.Debugf("request %v\n", getRequest)

	role := getRequest.Role

	response := request.MakeOkResponse()
	body := response.Body.(*protocol.GetResponse)
	body.ApiKey = &protocol.ApiKey{
		Role:    role,
		Token:   "placeholder-token",
		Enabled: true,
		UserID:  datascope,
	}
	return response
}
