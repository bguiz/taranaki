package Legacy

import (
	"github.com/autopilothq/lg"

	ct "github.com/autopilothq/banks/contract/types"
	proto "github.com/autopilothq/banks/protocol"
	"github.com/example/taranaki/resources/ApiKeys/protocol"
)

// Disable is something that you should document
func Disable(request *proto.Request, log lg.Log, aux ct.Auxiliary) *proto.Response {
	disableRequest := request.Body.(*protocol.DisableRequest)

	log.Debugf("request %v\n", disableRequest)

	// token := disableRequest.Token
	//TODO lookup ApiKeys using token

	response := request.MakeOkResponse()
	return response
}
