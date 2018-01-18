package Legacy

import (
	"github.com/autopilothq/lg"

	ct "github.com/autopilothq/banks/contract/types"
	proto "github.com/autopilothq/banks/protocol"
	"github.com/example/taranaki/resources/ApiKeys/protocol"
)

// Enable is something that you should document
func Enable(request *proto.Request, log lg.Log, aux ct.Auxiliary) *proto.Response {
	enableRequest := request.Body.(*protocol.EnableRequest)

	log.Debugf("request %v\n", enableRequest)

	// token := enableRequest.Token
	//TODO lookup ApiKeys using token

	response := request.MakeOkResponse()
	return response
}
