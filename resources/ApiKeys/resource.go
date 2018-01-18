// Code generated by 'banks apply'; DO NOT EDIT.

package ApiKeys

import (
	"context"
	"fmt"

	"github.com/autopilothq/lg"

	"github.com/autopilothq/banks/contract"
	ct "github.com/autopilothq/banks/contract/types"
	bp "github.com/autopilothq/banks/protocol"
	"github.com/autopilothq/banks/util"

	dep1 "github.com/example/taranaki/resources/ApiKeys/Legacy"
	dep0 "github.com/example/taranaki/resources/ApiKeys/protocol"
)

const (
	resourceID = 1

	msgIDGenerate = 65537
	msgIDEnable   = 65538
	msgIDDisable  = 65539
	msgIDGet      = 65540

	spIDLegacy = 1
)

func init() {
	var _ util.Fuse

	contract.StoragePlan(resourceID, spIDLegacy, dep1.SetupDatascope, dep1.TeardownDatascope)

	contract.Request(msgIDGenerate, &dep0.GenerateRequest{}, &dep0.GenerateResponse{},
		func(ctx context.Context, request *bp.Request, log lg.Log, aux ct.Auxiliary) (res *bp.Response) {
			var dispatch func(uint32) *bp.Response
			dispatch = func(step uint32) (res *bp.Response) { // nolint: shadow
				switch step {
				case 0:

					var (
						migrationID uint32
						storageID   uint32
						exists      bool
						err         error
						ready       bool
					)

					ready, err = contract.WaitUntilNullMigrationComplete(ctx, request.DatascopeID, resourceID)
					if err != nil {
						return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
							fmt.Sprintf("ApiKeys.Generate error waiting: %s", err.Error()))
					}
					if !ready {
						return dispatch(0)
					}

					migrationID, exists, err = contract.GetActiveMigrationID(request.DatascopeID, resourceID)
					if err != nil {
						return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
							fmt.Sprintf("ApiKeys.Generate error getting active migration: %s", err.Error()))
					}

					if exists {
						return dispatch(migrationID)
					}

					storageID, exists, err = contract.GetStoragePlanID(request.DatascopeID, resourceID)
					if err != nil {
						return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
							fmt.Sprintf("ApiKeys.Generate error getting current storage plan: %s", err.Error()))
					}
					if exists {
						return dispatch(storageID)
					}

					return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
						"ApiKeys.Generate failed: no storage plan available")

				case spIDLegacy:

					contract.TraceStoragePlanImpl(log, spIDLegacy, "Legacy", msgIDGenerate, request, res)
					return dep1.Generate(request, log, aux)

				default:
					return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
						fmt.Sprintf("Invalid ApiKeys.Generate dispatch step %d", step))
				}
			}

			return dispatch(0)
		})
	contract.Request(msgIDEnable, &dep0.EnableRequest{}, &dep0.EnableResponse{},
		func(ctx context.Context, request *bp.Request, log lg.Log, aux ct.Auxiliary) (res *bp.Response) {
			var dispatch func(uint32) *bp.Response
			dispatch = func(step uint32) (res *bp.Response) { // nolint: shadow
				switch step {
				case 0:

					var (
						migrationID uint32
						storageID   uint32
						exists      bool
						err         error
						ready       bool
					)

					ready, err = contract.WaitUntilNullMigrationComplete(ctx, request.DatascopeID, resourceID)
					if err != nil {
						return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
							fmt.Sprintf("ApiKeys.Enable error waiting: %s", err.Error()))
					}
					if !ready {
						return dispatch(0)
					}

					migrationID, exists, err = contract.GetActiveMigrationID(request.DatascopeID, resourceID)
					if err != nil {
						return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
							fmt.Sprintf("ApiKeys.Enable error getting active migration: %s", err.Error()))
					}

					if exists {
						return dispatch(migrationID)
					}

					storageID, exists, err = contract.GetStoragePlanID(request.DatascopeID, resourceID)
					if err != nil {
						return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
							fmt.Sprintf("ApiKeys.Enable error getting current storage plan: %s", err.Error()))
					}
					if exists {
						return dispatch(storageID)
					}

					return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
						"ApiKeys.Enable failed: no storage plan available")

				case spIDLegacy:

					contract.TraceStoragePlanImpl(log, spIDLegacy, "Legacy", msgIDEnable, request, res)
					return dep1.Enable(request, log, aux)

				default:
					return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
						fmt.Sprintf("Invalid ApiKeys.Enable dispatch step %d", step))
				}
			}

			return dispatch(0)
		})
	contract.Request(msgIDDisable, &dep0.DisableRequest{}, &dep0.DisableResponse{},
		func(ctx context.Context, request *bp.Request, log lg.Log, aux ct.Auxiliary) (res *bp.Response) {
			var dispatch func(uint32) *bp.Response
			dispatch = func(step uint32) (res *bp.Response) { // nolint: shadow
				switch step {
				case 0:

					var (
						migrationID uint32
						storageID   uint32
						exists      bool
						err         error
						ready       bool
					)

					ready, err = contract.WaitUntilNullMigrationComplete(ctx, request.DatascopeID, resourceID)
					if err != nil {
						return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
							fmt.Sprintf("ApiKeys.Disable error waiting: %s", err.Error()))
					}
					if !ready {
						return dispatch(0)
					}

					migrationID, exists, err = contract.GetActiveMigrationID(request.DatascopeID, resourceID)
					if err != nil {
						return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
							fmt.Sprintf("ApiKeys.Disable error getting active migration: %s", err.Error()))
					}

					if exists {
						return dispatch(migrationID)
					}

					storageID, exists, err = contract.GetStoragePlanID(request.DatascopeID, resourceID)
					if err != nil {
						return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
							fmt.Sprintf("ApiKeys.Disable error getting current storage plan: %s", err.Error()))
					}
					if exists {
						return dispatch(storageID)
					}

					return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
						"ApiKeys.Disable failed: no storage plan available")

				case spIDLegacy:

					contract.TraceStoragePlanImpl(log, spIDLegacy, "Legacy", msgIDDisable, request, res)
					return dep1.Disable(request, log, aux)

				default:
					return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
						fmt.Sprintf("Invalid ApiKeys.Disable dispatch step %d", step))
				}
			}

			return dispatch(0)
		})
	contract.Request(msgIDGet, &dep0.GetRequest{}, &dep0.GetResponse{},
		func(ctx context.Context, request *bp.Request, log lg.Log, aux ct.Auxiliary) (res *bp.Response) {
			var dispatch func(uint32) *bp.Response
			dispatch = func(step uint32) (res *bp.Response) { // nolint: shadow
				switch step {
				case 0:

					var (
						migrationID uint32
						storageID   uint32
						exists      bool
						err         error
						ready       bool
					)

					ready, err = contract.WaitUntilNullMigrationComplete(ctx, request.DatascopeID, resourceID)
					if err != nil {
						return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
							fmt.Sprintf("ApiKeys.Get error waiting: %s", err.Error()))
					}
					if !ready {
						return dispatch(0)
					}

					migrationID, exists, err = contract.GetActiveMigrationID(request.DatascopeID, resourceID)
					if err != nil {
						return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
							fmt.Sprintf("ApiKeys.Get error getting active migration: %s", err.Error()))
					}

					if exists {
						return dispatch(migrationID)
					}

					storageID, exists, err = contract.GetStoragePlanID(request.DatascopeID, resourceID)
					if err != nil {
						return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
							fmt.Sprintf("ApiKeys.Get error getting current storage plan: %s", err.Error()))
					}
					if exists {
						return dispatch(storageID)
					}

					return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
						"ApiKeys.Get failed: no storage plan available")

				case spIDLegacy:

					contract.TraceStoragePlanImpl(log, spIDLegacy, "Legacy", msgIDGet, request, res)
					return dep1.Get(request, log, aux)

				default:
					return request.MakeFatalResponse(bp.ResponseCode_ERR_UNKNOWN,
						fmt.Sprintf("Invalid ApiKeys.Get dispatch step %d", step))
				}
			}

			return dispatch(0)
		})
}