package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/RackHD/neighborhood-manager/swagger/restapi/operations"
	"github.com/RackHD/neighborhood-manager/swagger/restapi/operations/api_2_0"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../monorail-2.0.yaml

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

	api.API20AddRoleHandler = api_2_0.AddRoleHandlerFunc(func(params api_2_0.AddRoleParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.AddRole has not yet been implemented")
	})
	api.API20AddUserHandler = api_2_0.AddUserHandlerFunc(func(params api_2_0.AddUserParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.AddUser has not yet been implemented")
	})
	api.API20CatalogsGetHandler = api_2_0.CatalogsGetHandlerFunc(func(params api_2_0.CatalogsGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.CatalogsGet has not yet been implemented")
	})
	api.API20CatalogsIDGetHandler = api_2_0.CatalogsIDGetHandlerFunc(func(params api_2_0.CatalogsIDGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.CatalogsIDGet has not yet been implemented")
	})
	api.API20ConfigGetHandler = api_2_0.ConfigGetHandlerFunc(func(params api_2_0.ConfigGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ConfigGet has not yet been implemented")
	})
	api.API20ConfigPatchHandler = api_2_0.ConfigPatchHandlerFunc(func(params api_2_0.ConfigPatchParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ConfigPatch has not yet been implemented")
	})
	api.API20CreateTagHandler = api_2_0.CreateTagHandlerFunc(func(params api_2_0.CreateTagParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.CreateTag has not yet been implemented")
	})
	api.API20DeleteTagHandler = api_2_0.DeleteTagHandlerFunc(func(params api_2_0.DeleteTagParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.DeleteTag has not yet been implemented")
	})
	api.API20FilesDeleteHandler = api_2_0.FilesDeleteHandlerFunc(func(params api_2_0.FilesDeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.FilesDelete has not yet been implemented")
	})
	api.API20FilesGetHandler = api_2_0.FilesGetHandlerFunc(func(params api_2_0.FilesGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.FilesGet has not yet been implemented")
	})
	api.API20FilesGetAllHandler = api_2_0.FilesGetAllHandlerFunc(func(params api_2_0.FilesGetAllParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.FilesGetAll has not yet been implemented")
	})
	api.API20FilesMd5GetHandler = api_2_0.FilesMd5GetHandlerFunc(func(params api_2_0.FilesMd5GetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.FilesMd5Get has not yet been implemented")
	})
	api.API20FilesMetadataGetHandler = api_2_0.FilesMetadataGetHandlerFunc(func(params api_2_0.FilesMetadataGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.FilesMetadataGet has not yet been implemented")
	})
	api.API20FilesPutHandler = api_2_0.FilesPutHandlerFunc(func(params api_2_0.FilesPutParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.FilesPut has not yet been implemented")
	})
	api.API20GetAllTagsHandler = api_2_0.GetAllTagsHandlerFunc(func(params api_2_0.GetAllTagsParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.GetAllTags has not yet been implemented")
	})
	api.API20GetBootstrapHandler = api_2_0.GetBootstrapHandlerFunc(func(params api_2_0.GetBootstrapParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.GetBootstrap has not yet been implemented")
	})
	api.API20GetNodesByTagHandler = api_2_0.GetNodesByTagHandlerFunc(func(params api_2_0.GetNodesByTagParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.GetNodesByTag has not yet been implemented")
	})
	api.API20GetRoleHandler = api_2_0.GetRoleHandlerFunc(func(params api_2_0.GetRoleParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.GetRole has not yet been implemented")
	})
	api.API20GetTagHandler = api_2_0.GetTagHandlerFunc(func(params api_2_0.GetTagParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.GetTag has not yet been implemented")
	})
	api.API20GetTasksByIDHandler = api_2_0.GetTasksByIDHandlerFunc(func(params api_2_0.GetTasksByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.GetTasksByID has not yet been implemented")
	})
	api.API20GetUserHandler = api_2_0.GetUserHandlerFunc(func(params api_2_0.GetUserParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.GetUser has not yet been implemented")
	})
	api.API20ListRolesHandler = api_2_0.ListRolesHandlerFunc(func(params api_2_0.ListRolesParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ListRoles has not yet been implemented")
	})
	api.API20ListUsersHandler = api_2_0.ListUsersHandlerFunc(func(params api_2_0.ListUsersParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ListUsers has not yet been implemented")
	})
	api.API20LookupsDelByIDHandler = api_2_0.LookupsDelByIDHandlerFunc(func(params api_2_0.LookupsDelByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.LookupsDelByID has not yet been implemented")
	})
	api.API20LookupsGetHandler = api_2_0.LookupsGetHandlerFunc(func(params api_2_0.LookupsGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.LookupsGet has not yet been implemented")
	})
	api.API20LookupsGetByIDHandler = api_2_0.LookupsGetByIDHandlerFunc(func(params api_2_0.LookupsGetByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.LookupsGetByID has not yet been implemented")
	})
	api.API20LookupsPatchByIDHandler = api_2_0.LookupsPatchByIDHandlerFunc(func(params api_2_0.LookupsPatchByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.LookupsPatchByID has not yet been implemented")
	})
	api.API20LookupsPostHandler = api_2_0.LookupsPostHandlerFunc(func(params api_2_0.LookupsPostParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.LookupsPost has not yet been implemented")
	})
	api.API20ModifyRoleHandler = api_2_0.ModifyRoleHandlerFunc(func(params api_2_0.ModifyRoleParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ModifyRole has not yet been implemented")
	})
	api.API20ModifyUserHandler = api_2_0.ModifyUserHandlerFunc(func(params api_2_0.ModifyUserParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ModifyUser has not yet been implemented")
	})
	api.API20NodesAddRelationsHandler = api_2_0.NodesAddRelationsHandlerFunc(func(params api_2_0.NodesAddRelationsParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesAddRelations has not yet been implemented")
	})
	api.API20NodesDelByIDHandler = api_2_0.NodesDelByIDHandlerFunc(func(params api_2_0.NodesDelByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesDelByID has not yet been implemented")
	})
	api.API20NodesDelRelationsHandler = api_2_0.NodesDelRelationsHandlerFunc(func(params api_2_0.NodesDelRelationsParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesDelRelations has not yet been implemented")
	})
	api.API20NodesDelTagByIDHandler = api_2_0.NodesDelTagByIDHandlerFunc(func(params api_2_0.NodesDelTagByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesDelTagByID has not yet been implemented")
	})
	api.API20NodesGetAllHandler = api_2_0.NodesGetAllHandlerFunc(func(params api_2_0.NodesGetAllParams) middleware.Responder {
		// pass params.HTTPRequest

		return middleware.Responder
		return middleware.NotImplemented("operation api_2_0.NodesGetAll has not yet been implemented")
	})
	api.API20NodesGetByIDHandler = api_2_0.NodesGetByIDHandlerFunc(func(params api_2_0.NodesGetByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesGetByID has not yet been implemented")
	})
	api.API20NodesGetCatalogByIDHandler = api_2_0.NodesGetCatalogByIDHandlerFunc(func(params api_2_0.NodesGetCatalogByIDParams) middleware.Responder {
		params.
		return middleware.NotImplemented("operation api_2_0.NodesGetCatalogByID has not yet been implemented")
	})
	api.API20NodesGetCatalogSourceByIDHandler = api_2_0.NodesGetCatalogSourceByIDHandlerFunc(func(params api_2_0.NodesGetCatalogSourceByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesGetCatalogSourceByID has not yet been implemented")
	})
	api.API20NodesGetObmsByNodeIDHandler = api_2_0.NodesGetObmsByNodeIDHandlerFunc(func(params api_2_0.NodesGetObmsByNodeIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesGetObmsByNodeID has not yet been implemented")
	})
	api.API20NodesGetPollersByIDHandler = api_2_0.NodesGetPollersByIDHandlerFunc(func(params api_2_0.NodesGetPollersByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesGetPollersByID has not yet been implemented")
	})
	api.API20NodesGetRelationsHandler = api_2_0.NodesGetRelationsHandlerFunc(func(params api_2_0.NodesGetRelationsParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesGetRelations has not yet been implemented")
	})
	api.API20NodesGetSSHByIDHandler = api_2_0.NodesGetSSHByIDHandlerFunc(func(params api_2_0.NodesGetSSHByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesGetSSHByID has not yet been implemented")
	})
	api.API20NodesGetTagsByIDHandler = api_2_0.NodesGetTagsByIDHandlerFunc(func(params api_2_0.NodesGetTagsByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesGetTagsByID has not yet been implemented")
	})
	api.API20NodesGetWorkflowByIDHandler = api_2_0.NodesGetWorkflowByIDHandlerFunc(func(params api_2_0.NodesGetWorkflowByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesGetWorkflowByID has not yet been implemented")
	})
	api.API20NodesMasterDelTagByIDHandler = api_2_0.NodesMasterDelTagByIDHandlerFunc(func(params api_2_0.NodesMasterDelTagByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesMasterDelTagByID has not yet been implemented")
	})
	api.API20NodesPatchByIDHandler = api_2_0.NodesPatchByIDHandlerFunc(func(params api_2_0.NodesPatchByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesPatchByID has not yet been implemented")
	})
	api.API20NodesPatchTagByIDHandler = api_2_0.NodesPatchTagByIDHandlerFunc(func(params api_2_0.NodesPatchTagByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesPatchTagByID has not yet been implemented")
	})
	api.API20NodesPostHandler = api_2_0.NodesPostHandlerFunc(func(params api_2_0.NodesPostParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesPost has not yet been implemented")
	})
	api.API20NodesPostSSHByIDHandler = api_2_0.NodesPostSSHByIDHandlerFunc(func(params api_2_0.NodesPostSSHByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesPostSSHByID has not yet been implemented")
	})
	api.API20NodesPostWorkflowByIDHandler = api_2_0.NodesPostWorkflowByIDHandlerFunc(func(params api_2_0.NodesPostWorkflowByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesPostWorkflowByID has not yet been implemented")
	})
	api.API20NodesPutObmsByNodeIDHandler = api_2_0.NodesPutObmsByNodeIDHandlerFunc(func(params api_2_0.NodesPutObmsByNodeIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesPutObmsByNodeID has not yet been implemented")
	})
	api.API20NodesWorkflowActionByIDHandler = api_2_0.NodesWorkflowActionByIDHandlerFunc(func(params api_2_0.NodesWorkflowActionByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.NodesWorkflowActionByID has not yet been implemented")
	})
	api.API20ObmsDefinitionsGetAllHandler = api_2_0.ObmsDefinitionsGetAllHandlerFunc(func(params api_2_0.ObmsDefinitionsGetAllParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ObmsDefinitionsGetAll has not yet been implemented")
	})
	api.API20ObmsDefinitionsGetByNameHandler = api_2_0.ObmsDefinitionsGetByNameHandlerFunc(func(params api_2_0.ObmsDefinitionsGetByNameParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ObmsDefinitionsGetByName has not yet been implemented")
	})
	api.API20ObmsDeleteByIDHandler = api_2_0.ObmsDeleteByIDHandlerFunc(func(params api_2_0.ObmsDeleteByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ObmsDeleteByID has not yet been implemented")
	})
	api.API20ObmsGetHandler = api_2_0.ObmsGetHandlerFunc(func(params api_2_0.ObmsGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ObmsGet has not yet been implemented")
	})
	api.API20ObmsGetByIDHandler = api_2_0.ObmsGetByIDHandlerFunc(func(params api_2_0.ObmsGetByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ObmsGetByID has not yet been implemented")
	})
	api.API20ObmsPatchByIDHandler = api_2_0.ObmsPatchByIDHandlerFunc(func(params api_2_0.ObmsPatchByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ObmsPatchByID has not yet been implemented")
	})
	api.API20ObmsPostLedHandler = api_2_0.ObmsPostLedHandlerFunc(func(params api_2_0.ObmsPostLedParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ObmsPostLed has not yet been implemented")
	})
	api.API20ObmsPutHandler = api_2_0.ObmsPutHandlerFunc(func(params api_2_0.ObmsPutParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ObmsPut has not yet been implemented")
	})
	api.API20PollersCurrentDataGetHandler = api_2_0.PollersCurrentDataGetHandlerFunc(func(params api_2_0.PollersCurrentDataGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.PollersCurrentDataGet has not yet been implemented")
	})
	api.API20PollersDataGetHandler = api_2_0.PollersDataGetHandlerFunc(func(params api_2_0.PollersDataGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.PollersDataGet has not yet been implemented")
	})
	api.API20PollersDeleteHandler = api_2_0.PollersDeleteHandlerFunc(func(params api_2_0.PollersDeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.PollersDelete has not yet been implemented")
	})
	api.API20PollersGetHandler = api_2_0.PollersGetHandlerFunc(func(params api_2_0.PollersGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.PollersGet has not yet been implemented")
	})
	api.API20PollersIDGetHandler = api_2_0.PollersIDGetHandlerFunc(func(params api_2_0.PollersIDGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.PollersIDGet has not yet been implemented")
	})
	api.API20PollersLibByIDGetHandler = api_2_0.PollersLibByIDGetHandlerFunc(func(params api_2_0.PollersLibByIDGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.PollersLibByIDGet has not yet been implemented")
	})
	api.API20PollersLibGetHandler = api_2_0.PollersLibGetHandlerFunc(func(params api_2_0.PollersLibGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.PollersLibGet has not yet been implemented")
	})
	api.API20PollersPatchHandler = api_2_0.PollersPatchHandlerFunc(func(params api_2_0.PollersPatchParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.PollersPatch has not yet been implemented")
	})
	api.API20PollersPostHandler = api_2_0.PollersPostHandlerFunc(func(params api_2_0.PollersPostParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.PollersPost has not yet been implemented")
	})
	api.API20PostTaskByIDHandler = api_2_0.PostTaskByIDHandlerFunc(func(params api_2_0.PostTaskByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.PostTaskByID has not yet been implemented")
	})
	api.API20PostWorkflowByIDHandler = api_2_0.PostWorkflowByIDHandlerFunc(func(params api_2_0.PostWorkflowByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.PostWorkflowByID has not yet been implemented")
	})
	api.API20ProfilesGetHandler = api_2_0.ProfilesGetHandlerFunc(func(params api_2_0.ProfilesGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ProfilesGet has not yet been implemented")
	})
	api.API20ProfilesGetLibByNameHandler = api_2_0.ProfilesGetLibByNameHandlerFunc(func(params api_2_0.ProfilesGetLibByNameParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ProfilesGetLibByName has not yet been implemented")
	})
	api.API20ProfilesGetMetadataHandler = api_2_0.ProfilesGetMetadataHandlerFunc(func(params api_2_0.ProfilesGetMetadataParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ProfilesGetMetadata has not yet been implemented")
	})
	api.API20ProfilesGetMetadataByNameHandler = api_2_0.ProfilesGetMetadataByNameHandlerFunc(func(params api_2_0.ProfilesGetMetadataByNameParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ProfilesGetMetadataByName has not yet been implemented")
	})
	api.API20ProfilesGetSwitchVendorHandler = api_2_0.ProfilesGetSwitchVendorHandlerFunc(func(params api_2_0.ProfilesGetSwitchVendorParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ProfilesGetSwitchVendor has not yet been implemented")
	})
	api.API20ProfilesPostSwitchErrorHandler = api_2_0.ProfilesPostSwitchErrorHandlerFunc(func(params api_2_0.ProfilesPostSwitchErrorParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ProfilesPostSwitchError has not yet been implemented")
	})
	api.API20ProfilesPutLibByNameHandler = api_2_0.ProfilesPutLibByNameHandlerFunc(func(params api_2_0.ProfilesPutLibByNameParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ProfilesPutLibByName has not yet been implemented")
	})
	api.API20RemoveRoleHandler = api_2_0.RemoveRoleHandlerFunc(func(params api_2_0.RemoveRoleParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.RemoveRole has not yet been implemented")
	})
	api.API20RemoveUserHandler = api_2_0.RemoveUserHandlerFunc(func(params api_2_0.RemoveUserParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.RemoveUser has not yet been implemented")
	})
	api.API20SchemasGetHandler = api_2_0.SchemasGetHandlerFunc(func(params api_2_0.SchemasGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.SchemasGet has not yet been implemented")
	})
	api.API20SchemasIDGetHandler = api_2_0.SchemasIDGetHandlerFunc(func(params api_2_0.SchemasIDGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.SchemasIDGet has not yet been implemented")
	})
	api.API20SkuPackPostHandler = api_2_0.SkuPackPostHandlerFunc(func(params api_2_0.SkuPackPostParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.SkuPackPost has not yet been implemented")
	})
	api.API20SkusGetHandler = api_2_0.SkusGetHandlerFunc(func(params api_2_0.SkusGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.SkusGet has not yet been implemented")
	})
	api.API20SkusIDDeleteHandler = api_2_0.SkusIDDeleteHandlerFunc(func(params api_2_0.SkusIDDeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.SkusIDDelete has not yet been implemented")
	})
	api.API20SkusIDDeletePackHandler = api_2_0.SkusIDDeletePackHandlerFunc(func(params api_2_0.SkusIDDeletePackParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.SkusIDDeletePack has not yet been implemented")
	})
	api.API20SkusIDGetHandler = api_2_0.SkusIDGetHandlerFunc(func(params api_2_0.SkusIDGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.SkusIDGet has not yet been implemented")
	})
	api.API20SkusIDGetNodesHandler = api_2_0.SkusIDGetNodesHandlerFunc(func(params api_2_0.SkusIDGetNodesParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.SkusIDGetNodes has not yet been implemented")
	})
	api.API20SkusIDPutPackHandler = api_2_0.SkusIDPutPackHandlerFunc(func(params api_2_0.SkusIDPutPackParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.SkusIDPutPack has not yet been implemented")
	})
	api.API20SkusPatchHandler = api_2_0.SkusPatchHandlerFunc(func(params api_2_0.SkusPatchParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.SkusPatch has not yet been implemented")
	})
	api.API20SkusPostHandler = api_2_0.SkusPostHandlerFunc(func(params api_2_0.SkusPostParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.SkusPost has not yet been implemented")
	})
	api.API20SkusPutHandler = api_2_0.SkusPutHandlerFunc(func(params api_2_0.SkusPutParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.SkusPut has not yet been implemented")
	})
	api.API20TaskSchemasGetHandler = api_2_0.TaskSchemasGetHandlerFunc(func(params api_2_0.TaskSchemasGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.TaskSchemasGet has not yet been implemented")
	})
	api.API20TaskSchemasIDGetHandler = api_2_0.TaskSchemasIDGetHandlerFunc(func(params api_2_0.TaskSchemasIDGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.TaskSchemasIDGet has not yet been implemented")
	})
	api.API20TemplatesGetByNameHandler = api_2_0.TemplatesGetByNameHandlerFunc(func(params api_2_0.TemplatesGetByNameParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.TemplatesGetByName has not yet been implemented")
	})
	api.API20TemplatesHeadByNameHandler = api_2_0.TemplatesHeadByNameHandlerFunc(func(params api_2_0.TemplatesHeadByNameParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.TemplatesHeadByName has not yet been implemented")
	})
	api.API20TemplatesLibDeleteHandler = api_2_0.TemplatesLibDeleteHandlerFunc(func(params api_2_0.TemplatesLibDeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.TemplatesLibDelete has not yet been implemented")
	})
	api.API20TemplatesLibGetHandler = api_2_0.TemplatesLibGetHandlerFunc(func(params api_2_0.TemplatesLibGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.TemplatesLibGet has not yet been implemented")
	})
	api.API20TemplatesLibPutHandler = api_2_0.TemplatesLibPutHandlerFunc(func(params api_2_0.TemplatesLibPutParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.TemplatesLibPut has not yet been implemented")
	})
	api.API20TemplatesMetaGetHandler = api_2_0.TemplatesMetaGetHandlerFunc(func(params api_2_0.TemplatesMetaGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.TemplatesMetaGet has not yet been implemented")
	})
	api.API20TemplatesMetaGetByNameHandler = api_2_0.TemplatesMetaGetByNameHandlerFunc(func(params api_2_0.TemplatesMetaGetByNameParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.TemplatesMetaGetByName has not yet been implemented")
	})
	api.API20ViewsDeleteHandler = api_2_0.ViewsDeleteHandlerFunc(func(params api_2_0.ViewsDeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ViewsDelete has not yet been implemented")
	})
	api.API20ViewsGetHandler = api_2_0.ViewsGetHandlerFunc(func(params api_2_0.ViewsGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ViewsGet has not yet been implemented")
	})
	api.API20ViewsGetByIDHandler = api_2_0.ViewsGetByIDHandlerFunc(func(params api_2_0.ViewsGetByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ViewsGetByID has not yet been implemented")
	})
	api.API20ViewsPutHandler = api_2_0.ViewsPutHandlerFunc(func(params api_2_0.ViewsPutParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.ViewsPut has not yet been implemented")
	})
	api.API20WorkflowsActionHandler = api_2_0.WorkflowsActionHandlerFunc(func(params api_2_0.WorkflowsActionParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.WorkflowsAction has not yet been implemented")
	})
	api.API20WorkflowsDeleteByInstanceIDHandler = api_2_0.WorkflowsDeleteByInstanceIDHandlerFunc(func(params api_2_0.WorkflowsDeleteByInstanceIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.WorkflowsDeleteByInstanceID has not yet been implemented")
	})
	api.API20WorkflowsDeleteGraphsByNameHandler = api_2_0.WorkflowsDeleteGraphsByNameHandlerFunc(func(params api_2_0.WorkflowsDeleteGraphsByNameParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.WorkflowsDeleteGraphsByName has not yet been implemented")
	})
	api.API20WorkflowsDeleteTasksByNameHandler = api_2_0.WorkflowsDeleteTasksByNameHandlerFunc(func(params api_2_0.WorkflowsDeleteTasksByNameParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.WorkflowsDeleteTasksByName has not yet been implemented")
	})
	api.API20WorkflowsGetHandler = api_2_0.WorkflowsGetHandlerFunc(func(params api_2_0.WorkflowsGetParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.WorkflowsGet has not yet been implemented")
	})
	api.API20WorkflowsGetAllTasksHandler = api_2_0.WorkflowsGetAllTasksHandlerFunc(func(params api_2_0.WorkflowsGetAllTasksParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.WorkflowsGetAllTasks has not yet been implemented")
	})
	api.API20WorkflowsGetByInstanceIDHandler = api_2_0.WorkflowsGetByInstanceIDHandlerFunc(func(params api_2_0.WorkflowsGetByInstanceIDParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.WorkflowsGetByInstanceID has not yet been implemented")
	})
	api.API20WorkflowsGetGraphsHandler = api_2_0.WorkflowsGetGraphsHandlerFunc(func(params api_2_0.WorkflowsGetGraphsParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.WorkflowsGetGraphs has not yet been implemented")
	})
	api.API20WorkflowsGetGraphsByNameHandler = api_2_0.WorkflowsGetGraphsByNameHandlerFunc(func(params api_2_0.WorkflowsGetGraphsByNameParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.WorkflowsGetGraphsByName has not yet been implemented")
	})
	api.API20WorkflowsGetTasksByNameHandler = api_2_0.WorkflowsGetTasksByNameHandlerFunc(func(params api_2_0.WorkflowsGetTasksByNameParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.WorkflowsGetTasksByName has not yet been implemented")
	})
	api.API20WorkflowsPostHandler = api_2_0.WorkflowsPostHandlerFunc(func(params api_2_0.WorkflowsPostParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.WorkflowsPost has not yet been implemented")
	})
	api.API20WorkflowsPutGraphsHandler = api_2_0.WorkflowsPutGraphsHandlerFunc(func(params api_2_0.WorkflowsPutGraphsParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.WorkflowsPutGraphs has not yet been implemented")
	})
	api.API20WorkflowsPutTaskHandler = api_2_0.WorkflowsPutTaskHandlerFunc(func(params api_2_0.WorkflowsPutTaskParams) middleware.Responder {
		return middleware.NotImplemented("operation api_2_0.WorkflowsPutTask has not yet been implemented")
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
