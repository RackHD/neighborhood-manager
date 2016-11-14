package monorail

import (
	"crypto/tls"
	"net/http"

	"github.com/RackHD/neighborhood-manager/rackhd/api/proxy"
	"github.com/RackHD/neighborhood-manager/rackhd/api/swagger/monorail/operations"
	"github.com/RackHD/neighborhood-manager/rackhd/api/swagger/monorail/operations/api_2_0"
	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"golang.org/x/net/context"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../monorail-2.0.yaml --server-package monorail --exclude-main --with-context

func configureFlags(api *operations.RackHD20API) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.RackHD20API) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.BinConsumer = runtime.ByteStreamConsumer()

	api.UrlformConsumer = runtime.DiscardConsumer

	api.MultipartformConsumer = runtime.DiscardConsumer

	api.TxtConsumer = runtime.TextConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.API20AddRoleHandler = api_2_0.AddRoleHandlerFunc(func(ctx context.Context, params api_2_0.AddRoleParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20AddUserHandler = api_2_0.AddUserHandlerFunc(func(ctx context.Context, params api_2_0.AddUserParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20CatalogsGetHandler = api_2_0.CatalogsGetHandlerFunc(func(ctx context.Context, params api_2_0.CatalogsGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20CatalogsIDGetHandler = api_2_0.CatalogsIDGetHandlerFunc(func(ctx context.Context, params api_2_0.CatalogsIDGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ConfigGetHandler = api_2_0.ConfigGetHandlerFunc(func(ctx context.Context, params api_2_0.ConfigGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ConfigPatchHandler = api_2_0.ConfigPatchHandlerFunc(func(ctx context.Context, params api_2_0.ConfigPatchParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20CreateTagHandler = api_2_0.CreateTagHandlerFunc(func(ctx context.Context, params api_2_0.CreateTagParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20DeleteTagHandler = api_2_0.DeleteTagHandlerFunc(func(ctx context.Context, params api_2_0.DeleteTagParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20FilesDeleteHandler = api_2_0.FilesDeleteHandlerFunc(func(ctx context.Context, params api_2_0.FilesDeleteParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20FilesGetHandler = api_2_0.FilesGetHandlerFunc(func(ctx context.Context, params api_2_0.FilesGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20FilesGetAllHandler = api_2_0.FilesGetAllHandlerFunc(func(ctx context.Context, params api_2_0.FilesGetAllParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20FilesMd5GetHandler = api_2_0.FilesMd5GetHandlerFunc(func(ctx context.Context, params api_2_0.FilesMd5GetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20FilesMetadataGetHandler = api_2_0.FilesMetadataGetHandlerFunc(func(ctx context.Context, params api_2_0.FilesMetadataGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20FilesPutHandler = api_2_0.FilesPutHandlerFunc(func(ctx context.Context, params api_2_0.FilesPutParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20GetAllTagsHandler = api_2_0.GetAllTagsHandlerFunc(func(ctx context.Context, params api_2_0.GetAllTagsParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20GetBootstrapHandler = api_2_0.GetBootstrapHandlerFunc(func(ctx context.Context, params api_2_0.GetBootstrapParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20GetNodesByTagHandler = api_2_0.GetNodesByTagHandlerFunc(func(ctx context.Context, params api_2_0.GetNodesByTagParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20GetRoleHandler = api_2_0.GetRoleHandlerFunc(func(ctx context.Context, params api_2_0.GetRoleParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20GetTagHandler = api_2_0.GetTagHandlerFunc(func(ctx context.Context, params api_2_0.GetTagParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20GetTasksByIDHandler = api_2_0.GetTasksByIDHandlerFunc(func(ctx context.Context, params api_2_0.GetTasksByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20GetUserHandler = api_2_0.GetUserHandlerFunc(func(ctx context.Context, params api_2_0.GetUserParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ListRolesHandler = api_2_0.ListRolesHandlerFunc(func(ctx context.Context, params api_2_0.ListRolesParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ListUsersHandler = api_2_0.ListUsersHandlerFunc(func(ctx context.Context, params api_2_0.ListUsersParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20LookupsDelByIDHandler = api_2_0.LookupsDelByIDHandlerFunc(func(ctx context.Context, params api_2_0.LookupsDelByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20LookupsGetHandler = api_2_0.LookupsGetHandlerFunc(func(ctx context.Context, params api_2_0.LookupsGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20LookupsGetByIDHandler = api_2_0.LookupsGetByIDHandlerFunc(func(ctx context.Context, params api_2_0.LookupsGetByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20LookupsPatchByIDHandler = api_2_0.LookupsPatchByIDHandlerFunc(func(ctx context.Context, params api_2_0.LookupsPatchByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20LookupsPostHandler = api_2_0.LookupsPostHandlerFunc(func(ctx context.Context, params api_2_0.LookupsPostParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ModifyRoleHandler = api_2_0.ModifyRoleHandlerFunc(func(ctx context.Context, params api_2_0.ModifyRoleParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ModifyUserHandler = api_2_0.ModifyUserHandlerFunc(func(ctx context.Context, params api_2_0.ModifyUserParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesAddRelationsHandler = api_2_0.NodesAddRelationsHandlerFunc(func(ctx context.Context, params api_2_0.NodesAddRelationsParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesDelByIDHandler = api_2_0.NodesDelByIDHandlerFunc(func(ctx context.Context, params api_2_0.NodesDelByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesDelRelationsHandler = api_2_0.NodesDelRelationsHandlerFunc(func(ctx context.Context, params api_2_0.NodesDelRelationsParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesDelTagByIDHandler = api_2_0.NodesDelTagByIDHandlerFunc(func(ctx context.Context, params api_2_0.NodesDelTagByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesGetAllHandler = api_2_0.NodesGetAllHandlerFunc(func(ctx context.Context, params api_2_0.NodesGetAllParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesGetByIDHandler = api_2_0.NodesGetByIDHandlerFunc(func(ctx context.Context, params api_2_0.NodesGetByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesGetCatalogByIDHandler = api_2_0.NodesGetCatalogByIDHandlerFunc(func(ctx context.Context, params api_2_0.NodesGetCatalogByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesGetCatalogSourceByIDHandler = api_2_0.NodesGetCatalogSourceByIDHandlerFunc(func(ctx context.Context, params api_2_0.NodesGetCatalogSourceByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesGetObmsByNodeIDHandler = api_2_0.NodesGetObmsByNodeIDHandlerFunc(func(ctx context.Context, params api_2_0.NodesGetObmsByNodeIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesGetPollersByIDHandler = api_2_0.NodesGetPollersByIDHandlerFunc(func(ctx context.Context, params api_2_0.NodesGetPollersByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesGetRelationsHandler = api_2_0.NodesGetRelationsHandlerFunc(func(ctx context.Context, params api_2_0.NodesGetRelationsParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesGetSSHByIDHandler = api_2_0.NodesGetSSHByIDHandlerFunc(func(ctx context.Context, params api_2_0.NodesGetSSHByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesGetTagsByIDHandler = api_2_0.NodesGetTagsByIDHandlerFunc(func(ctx context.Context, params api_2_0.NodesGetTagsByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesGetWorkflowByIDHandler = api_2_0.NodesGetWorkflowByIDHandlerFunc(func(ctx context.Context, params api_2_0.NodesGetWorkflowByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesMasterDelTagByIDHandler = api_2_0.NodesMasterDelTagByIDHandlerFunc(func(ctx context.Context, params api_2_0.NodesMasterDelTagByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesPatchByIDHandler = api_2_0.NodesPatchByIDHandlerFunc(func(ctx context.Context, params api_2_0.NodesPatchByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesPatchTagByIDHandler = api_2_0.NodesPatchTagByIDHandlerFunc(func(ctx context.Context, params api_2_0.NodesPatchTagByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesPostHandler = api_2_0.NodesPostHandlerFunc(func(ctx context.Context, params api_2_0.NodesPostParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesPostSSHByIDHandler = api_2_0.NodesPostSSHByIDHandlerFunc(func(ctx context.Context, params api_2_0.NodesPostSSHByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesPostWorkflowByIDHandler = api_2_0.NodesPostWorkflowByIDHandlerFunc(func(ctx context.Context, params api_2_0.NodesPostWorkflowByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesPutObmsByNodeIDHandler = api_2_0.NodesPutObmsByNodeIDHandlerFunc(func(ctx context.Context, params api_2_0.NodesPutObmsByNodeIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20NodesWorkflowActionByIDHandler = api_2_0.NodesWorkflowActionByIDHandlerFunc(func(ctx context.Context, params api_2_0.NodesWorkflowActionByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ObmsDefinitionsGetAllHandler = api_2_0.ObmsDefinitionsGetAllHandlerFunc(func(ctx context.Context, params api_2_0.ObmsDefinitionsGetAllParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ObmsDefinitionsGetByNameHandler = api_2_0.ObmsDefinitionsGetByNameHandlerFunc(func(ctx context.Context, params api_2_0.ObmsDefinitionsGetByNameParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ObmsDeleteByIDHandler = api_2_0.ObmsDeleteByIDHandlerFunc(func(ctx context.Context, params api_2_0.ObmsDeleteByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ObmsGetHandler = api_2_0.ObmsGetHandlerFunc(func(ctx context.Context, params api_2_0.ObmsGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ObmsGetByIDHandler = api_2_0.ObmsGetByIDHandlerFunc(func(ctx context.Context, params api_2_0.ObmsGetByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ObmsPatchByIDHandler = api_2_0.ObmsPatchByIDHandlerFunc(func(ctx context.Context, params api_2_0.ObmsPatchByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ObmsPostLedHandler = api_2_0.ObmsPostLedHandlerFunc(func(ctx context.Context, params api_2_0.ObmsPostLedParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ObmsPutHandler = api_2_0.ObmsPutHandlerFunc(func(ctx context.Context, params api_2_0.ObmsPutParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20PollersCurrentDataGetHandler = api_2_0.PollersCurrentDataGetHandlerFunc(func(ctx context.Context, params api_2_0.PollersCurrentDataGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20PollersDataGetHandler = api_2_0.PollersDataGetHandlerFunc(func(ctx context.Context, params api_2_0.PollersDataGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20PollersDeleteHandler = api_2_0.PollersDeleteHandlerFunc(func(ctx context.Context, params api_2_0.PollersDeleteParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20PollersGetHandler = api_2_0.PollersGetHandlerFunc(func(ctx context.Context, params api_2_0.PollersGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20PollersIDGetHandler = api_2_0.PollersIDGetHandlerFunc(func(ctx context.Context, params api_2_0.PollersIDGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20PollersLibByIDGetHandler = api_2_0.PollersLibByIDGetHandlerFunc(func(ctx context.Context, params api_2_0.PollersLibByIDGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20PollersLibGetHandler = api_2_0.PollersLibGetHandlerFunc(func(ctx context.Context, params api_2_0.PollersLibGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20PollersPatchHandler = api_2_0.PollersPatchHandlerFunc(func(ctx context.Context, params api_2_0.PollersPatchParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20PollersPostHandler = api_2_0.PollersPostHandlerFunc(func(ctx context.Context, params api_2_0.PollersPostParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20PostTaskByIDHandler = api_2_0.PostTaskByIDHandlerFunc(func(ctx context.Context, params api_2_0.PostTaskByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20PostWorkflowByIDHandler = api_2_0.PostWorkflowByIDHandlerFunc(func(ctx context.Context, params api_2_0.PostWorkflowByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ProfilesGetHandler = api_2_0.ProfilesGetHandlerFunc(func(ctx context.Context, params api_2_0.ProfilesGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ProfilesGetLibByNameHandler = api_2_0.ProfilesGetLibByNameHandlerFunc(func(ctx context.Context, params api_2_0.ProfilesGetLibByNameParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ProfilesGetMetadataHandler = api_2_0.ProfilesGetMetadataHandlerFunc(func(ctx context.Context, params api_2_0.ProfilesGetMetadataParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ProfilesGetMetadataByNameHandler = api_2_0.ProfilesGetMetadataByNameHandlerFunc(func(ctx context.Context, params api_2_0.ProfilesGetMetadataByNameParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ProfilesGetSwitchVendorHandler = api_2_0.ProfilesGetSwitchVendorHandlerFunc(func(ctx context.Context, params api_2_0.ProfilesGetSwitchVendorParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ProfilesPostSwitchErrorHandler = api_2_0.ProfilesPostSwitchErrorHandlerFunc(func(ctx context.Context, params api_2_0.ProfilesPostSwitchErrorParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ProfilesPutLibByNameHandler = api_2_0.ProfilesPutLibByNameHandlerFunc(func(ctx context.Context, params api_2_0.ProfilesPutLibByNameParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20RemoveRoleHandler = api_2_0.RemoveRoleHandlerFunc(func(ctx context.Context, params api_2_0.RemoveRoleParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20RemoveUserHandler = api_2_0.RemoveUserHandlerFunc(func(ctx context.Context, params api_2_0.RemoveUserParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20SchemasGetHandler = api_2_0.SchemasGetHandlerFunc(func(ctx context.Context, params api_2_0.SchemasGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20SchemasIDGetHandler = api_2_0.SchemasIDGetHandlerFunc(func(ctx context.Context, params api_2_0.SchemasIDGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20SkuPackPostHandler = api_2_0.SkuPackPostHandlerFunc(func(ctx context.Context, params api_2_0.SkuPackPostParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20SkusGetHandler = api_2_0.SkusGetHandlerFunc(func(ctx context.Context, params api_2_0.SkusGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20SkusIDDeleteHandler = api_2_0.SkusIDDeleteHandlerFunc(func(ctx context.Context, params api_2_0.SkusIDDeleteParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20SkusIDDeletePackHandler = api_2_0.SkusIDDeletePackHandlerFunc(func(ctx context.Context, params api_2_0.SkusIDDeletePackParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20SkusIDGetHandler = api_2_0.SkusIDGetHandlerFunc(func(ctx context.Context, params api_2_0.SkusIDGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20SkusIDGetNodesHandler = api_2_0.SkusIDGetNodesHandlerFunc(func(ctx context.Context, params api_2_0.SkusIDGetNodesParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20SkusIDPutPackHandler = api_2_0.SkusIDPutPackHandlerFunc(func(ctx context.Context, params api_2_0.SkusIDPutPackParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20SkusPatchHandler = api_2_0.SkusPatchHandlerFunc(func(ctx context.Context, params api_2_0.SkusPatchParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20SkusPostHandler = api_2_0.SkusPostHandlerFunc(func(ctx context.Context, params api_2_0.SkusPostParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20SkusPutHandler = api_2_0.SkusPutHandlerFunc(func(ctx context.Context, params api_2_0.SkusPutParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20TaskSchemasGetHandler = api_2_0.TaskSchemasGetHandlerFunc(func(ctx context.Context, params api_2_0.TaskSchemasGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20TaskSchemasIDGetHandler = api_2_0.TaskSchemasIDGetHandlerFunc(func(ctx context.Context, params api_2_0.TaskSchemasIDGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20TemplatesHeadByNameHandler = api_2_0.TemplatesHeadByNameHandlerFunc(func(ctx context.Context, params api_2_0.TemplatesHeadByNameParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20TemplatesLibDeleteHandler = api_2_0.TemplatesLibDeleteHandlerFunc(func(ctx context.Context, params api_2_0.TemplatesLibDeleteParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20TemplatesLibGetHandler = api_2_0.TemplatesLibGetHandlerFunc(func(ctx context.Context, params api_2_0.TemplatesLibGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20TemplatesLibPutHandler = api_2_0.TemplatesLibPutHandlerFunc(func(ctx context.Context, params api_2_0.TemplatesLibPutParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20TemplatesMetaGetHandler = api_2_0.TemplatesMetaGetHandlerFunc(func(ctx context.Context, params api_2_0.TemplatesMetaGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20TemplatesMetaGetByNameHandler = api_2_0.TemplatesMetaGetByNameHandlerFunc(func(ctx context.Context, params api_2_0.TemplatesMetaGetByNameParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ViewsDeleteHandler = api_2_0.ViewsDeleteHandlerFunc(func(ctx context.Context, params api_2_0.ViewsDeleteParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ViewsGetHandler = api_2_0.ViewsGetHandlerFunc(func(ctx context.Context, params api_2_0.ViewsGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ViewsGetByIDHandler = api_2_0.ViewsGetByIDHandlerFunc(func(ctx context.Context, params api_2_0.ViewsGetByIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20ViewsPutHandler = api_2_0.ViewsPutHandlerFunc(func(ctx context.Context, params api_2_0.ViewsPutParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20WorkflowsActionHandler = api_2_0.WorkflowsActionHandlerFunc(func(ctx context.Context, params api_2_0.WorkflowsActionParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20WorkflowsDeleteByInstanceIDHandler = api_2_0.WorkflowsDeleteByInstanceIDHandlerFunc(func(ctx context.Context, params api_2_0.WorkflowsDeleteByInstanceIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20WorkflowsDeleteGraphsByNameHandler = api_2_0.WorkflowsDeleteGraphsByNameHandlerFunc(func(ctx context.Context, params api_2_0.WorkflowsDeleteGraphsByNameParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20WorkflowsDeleteTasksByNameHandler = api_2_0.WorkflowsDeleteTasksByNameHandlerFunc(func(ctx context.Context, params api_2_0.WorkflowsDeleteTasksByNameParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20WorkflowsGetHandler = api_2_0.WorkflowsGetHandlerFunc(func(ctx context.Context, params api_2_0.WorkflowsGetParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20WorkflowsGetAllTasksHandler = api_2_0.WorkflowsGetAllTasksHandlerFunc(func(ctx context.Context, params api_2_0.WorkflowsGetAllTasksParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20WorkflowsGetByInstanceIDHandler = api_2_0.WorkflowsGetByInstanceIDHandlerFunc(func(ctx context.Context, params api_2_0.WorkflowsGetByInstanceIDParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20WorkflowsGetGraphsHandler = api_2_0.WorkflowsGetGraphsHandlerFunc(func(ctx context.Context, params api_2_0.WorkflowsGetGraphsParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20WorkflowsGetGraphsByNameHandler = api_2_0.WorkflowsGetGraphsByNameHandlerFunc(func(ctx context.Context, params api_2_0.WorkflowsGetGraphsByNameParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20WorkflowsGetTasksByNameHandler = api_2_0.WorkflowsGetTasksByNameHandlerFunc(func(ctx context.Context, params api_2_0.WorkflowsGetTasksByNameParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20WorkflowsPostHandler = api_2_0.WorkflowsPostHandlerFunc(func(ctx context.Context, params api_2_0.WorkflowsPostParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20WorkflowsPutGraphsHandler = api_2_0.WorkflowsPutGraphsHandlerFunc(func(ctx context.Context, params api_2_0.WorkflowsPutGraphsParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})
	api.API20WorkflowsPutTaskHandler = api_2_0.WorkflowsPutTaskHandlerFunc(func(ctx context.Context, params api_2_0.WorkflowsPutTaskParams) middleware.Responder {
		return proxy.HandleAllMiddleware(params.HTTPRequest)
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
