// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"tides-server/pkg/restapi/operations/policy"
	"tides-server/pkg/restapi/operations/project"
	"tides-server/pkg/restapi/operations/resource"
	"tides-server/pkg/restapi/operations/template"
	"tides-server/pkg/restapi/operations/usage"
	"tides-server/pkg/restapi/operations/user"
)

// NewCloudTidesAPI creates a new CloudTides instance
func NewCloudTidesAPI(spec *loads.Document) *CloudTidesAPI {
	return &CloudTidesAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,

		JSONProducer: runtime.JSONProducer(),

		PolicyAddPolicyHandler: policy.AddPolicyHandlerFunc(func(params policy.AddPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.AddPolicy has not yet been implemented")
		}),
		ProjectAddProjectHandler: project.AddProjectHandlerFunc(func(params project.AddProjectParams) middleware.Responder {
			return middleware.NotImplemented("operation project.AddProject has not yet been implemented")
		}),
		UsageAddResourceUsageHandler: usage.AddResourceUsageHandlerFunc(func(params usage.AddResourceUsageParams) middleware.Responder {
			return middleware.NotImplemented("operation usage.AddResourceUsage has not yet been implemented")
		}),
		TemplateAddTemplateHandler: template.AddTemplateHandlerFunc(func(params template.AddTemplateParams) middleware.Responder {
			return middleware.NotImplemented("operation template.AddTemplate has not yet been implemented")
		}),
		UsageAddVMUsageHandler: usage.AddVMUsageHandlerFunc(func(params usage.AddVMUsageParams) middleware.Responder {
			return middleware.NotImplemented("operation usage.AddVMUsage has not yet been implemented")
		}),
		ResourceAddVcdResourceHandler: resource.AddVcdResourceHandlerFunc(func(params resource.AddVcdResourceParams) middleware.Responder {
			return middleware.NotImplemented("operation resource.AddVcdResource has not yet been implemented")
		}),
		ResourceAddVsphereResourceHandler: resource.AddVsphereResourceHandlerFunc(func(params resource.AddVsphereResourceParams) middleware.Responder {
			return middleware.NotImplemented("operation resource.AddVsphereResource has not yet been implemented")
		}),
		ProjectDeleteProjectHandler: project.DeleteProjectHandlerFunc(func(params project.DeleteProjectParams) middleware.Responder {
			return middleware.NotImplemented("operation project.DeleteProject has not yet been implemented")
		}),
		UsageDeleteResourceUsageHandler: usage.DeleteResourceUsageHandlerFunc(func(params usage.DeleteResourceUsageParams) middleware.Responder {
			return middleware.NotImplemented("operation usage.DeleteResourceUsage has not yet been implemented")
		}),
		TemplateDeleteTemplateHandler: template.DeleteTemplateHandlerFunc(func(params template.DeleteTemplateParams) middleware.Responder {
			return middleware.NotImplemented("operation template.DeleteTemplate has not yet been implemented")
		}),
		ResourceDeleteVcdResourceHandler: resource.DeleteVcdResourceHandlerFunc(func(params resource.DeleteVcdResourceParams) middleware.Responder {
			return middleware.NotImplemented("operation resource.DeleteVcdResource has not yet been implemented")
		}),
		ResourceDestroyVMHandler: resource.DestroyVMHandlerFunc(func(params resource.DestroyVMParams) middleware.Responder {
			return middleware.NotImplemented("operation resource.DestroyVM has not yet been implemented")
		}),
		UsageGetPastUsageHandler: usage.GetPastUsageHandlerFunc(func(params usage.GetPastUsageParams) middleware.Responder {
			return middleware.NotImplemented("operation usage.GetPastUsage has not yet been implemented")
		}),
		PolicyGetPolicyHandler: policy.GetPolicyHandlerFunc(func(params policy.GetPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.GetPolicy has not yet been implemented")
		}),
		UsageGetResourceUsageHandler: usage.GetResourceUsageHandlerFunc(func(params usage.GetResourceUsageParams) middleware.Responder {
			return middleware.NotImplemented("operation usage.GetResourceUsage has not yet been implemented")
		}),
		UserGetUserProfileHandler: user.GetUserProfileHandlerFunc(func(params user.GetUserProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUserProfile has not yet been implemented")
		}),
		ResourceGetVcdResourceHandler: resource.GetVcdResourceHandlerFunc(func(params resource.GetVcdResourceParams) middleware.Responder {
			return middleware.NotImplemented("operation resource.GetVcdResource has not yet been implemented")
		}),
		PolicyListPolicyHandler: policy.ListPolicyHandlerFunc(func(params policy.ListPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.ListPolicy has not yet been implemented")
		}),
		ProjectListProjectHandler: project.ListProjectHandlerFunc(func(params project.ListProjectParams) middleware.Responder {
			return middleware.NotImplemented("operation project.ListProject has not yet been implemented")
		}),
		TemplateListTemplateHandler: template.ListTemplateHandlerFunc(func(params template.ListTemplateParams) middleware.Responder {
			return middleware.NotImplemented("operation template.ListTemplate has not yet been implemented")
		}),
		ResourceListVcdResourceHandler: resource.ListVcdResourceHandlerFunc(func(params resource.ListVcdResourceParams) middleware.Responder {
			return middleware.NotImplemented("operation resource.ListVcdResource has not yet been implemented")
		}),
		ResourceListVsphereResourceHandler: resource.ListVsphereResourceHandlerFunc(func(params resource.ListVsphereResourceParams) middleware.Responder {
			return middleware.NotImplemented("operation resource.ListVsphereResource has not yet been implemented")
		}),
		UserRegisterUserHandler: user.RegisterUserHandlerFunc(func(params user.RegisterUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.RegisterUser has not yet been implemented")
		}),
		PolicyRemovePolicyHandler: policy.RemovePolicyHandlerFunc(func(params policy.RemovePolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.RemovePolicy has not yet been implemented")
		}),
		PolicyUpdatePolicyHandler: policy.UpdatePolicyHandlerFunc(func(params policy.UpdatePolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.UpdatePolicy has not yet been implemented")
		}),
		ProjectUpdateProjectHandler: project.UpdateProjectHandlerFunc(func(params project.UpdateProjectParams) middleware.Responder {
			return middleware.NotImplemented("operation project.UpdateProject has not yet been implemented")
		}),
		ResourceUpdateResourceHandler: resource.UpdateResourceHandlerFunc(func(params resource.UpdateResourceParams) middleware.Responder {
			return middleware.NotImplemented("operation resource.UpdateResource has not yet been implemented")
		}),
		UsageUpdateResourceUsageHandler: usage.UpdateResourceUsageHandlerFunc(func(params usage.UpdateResourceUsageParams) middleware.Responder {
			return middleware.NotImplemented("operation usage.UpdateResourceUsage has not yet been implemented")
		}),
		UserUpdateUserProfileHandler: user.UpdateUserProfileHandlerFunc(func(params user.UpdateUserProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UpdateUserProfile has not yet been implemented")
		}),
		UserUserLoginHandler: user.UserLoginHandlerFunc(func(params user.UserLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserLogin has not yet been implemented")
		}),
		ResourceValidateVcdResourceHandler: resource.ValidateVcdResourceHandlerFunc(func(params resource.ValidateVcdResourceParams) middleware.Responder {
			return middleware.NotImplemented("operation resource.ValidateVcdResource has not yet been implemented")
		}),
		ResourceValidateVsphereResourceHandler: resource.ValidateVsphereResourceHandlerFunc(func(params resource.ValidateVsphereResourceParams) middleware.Responder {
			return middleware.NotImplemented("operation resource.ValidateVsphereResource has not yet been implemented")
		}),
	}
}

/*CloudTidesAPI the cloud tides API */
type CloudTidesAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for the following mime types:
	//   - multipart/form-data
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// PolicyAddPolicyHandler sets the operation handler for the add policy operation
	PolicyAddPolicyHandler policy.AddPolicyHandler
	// ProjectAddProjectHandler sets the operation handler for the add project operation
	ProjectAddProjectHandler project.AddProjectHandler
	// UsageAddResourceUsageHandler sets the operation handler for the add resource usage operation
	UsageAddResourceUsageHandler usage.AddResourceUsageHandler
	// TemplateAddTemplateHandler sets the operation handler for the add template operation
	TemplateAddTemplateHandler template.AddTemplateHandler
	// UsageAddVMUsageHandler sets the operation handler for the add VM usage operation
	UsageAddVMUsageHandler usage.AddVMUsageHandler
	// ResourceAddVcdResourceHandler sets the operation handler for the add vcd resource operation
	ResourceAddVcdResourceHandler resource.AddVcdResourceHandler
	// ResourceAddVsphereResourceHandler sets the operation handler for the add vsphere resource operation
	ResourceAddVsphereResourceHandler resource.AddVsphereResourceHandler
	// ProjectDeleteProjectHandler sets the operation handler for the delete project operation
	ProjectDeleteProjectHandler project.DeleteProjectHandler
	// UsageDeleteResourceUsageHandler sets the operation handler for the delete resource usage operation
	UsageDeleteResourceUsageHandler usage.DeleteResourceUsageHandler
	// TemplateDeleteTemplateHandler sets the operation handler for the delete template operation
	TemplateDeleteTemplateHandler template.DeleteTemplateHandler
	// ResourceDeleteVcdResourceHandler sets the operation handler for the delete vcd resource operation
	ResourceDeleteVcdResourceHandler resource.DeleteVcdResourceHandler
	// ResourceDestroyVMHandler sets the operation handler for the destroy VM operation
	ResourceDestroyVMHandler resource.DestroyVMHandler
	// UsageGetPastUsageHandler sets the operation handler for the get past usage operation
	UsageGetPastUsageHandler usage.GetPastUsageHandler
	// PolicyGetPolicyHandler sets the operation handler for the get policy operation
	PolicyGetPolicyHandler policy.GetPolicyHandler
	// UsageGetResourceUsageHandler sets the operation handler for the get resource usage operation
	UsageGetResourceUsageHandler usage.GetResourceUsageHandler
	// UserGetUserProfileHandler sets the operation handler for the get user profile operation
	UserGetUserProfileHandler user.GetUserProfileHandler
	// ResourceGetVcdResourceHandler sets the operation handler for the get vcd resource operation
	ResourceGetVcdResourceHandler resource.GetVcdResourceHandler
	// PolicyListPolicyHandler sets the operation handler for the list policy operation
	PolicyListPolicyHandler policy.ListPolicyHandler
	// ProjectListProjectHandler sets the operation handler for the list project operation
	ProjectListProjectHandler project.ListProjectHandler
	// TemplateListTemplateHandler sets the operation handler for the list template operation
	TemplateListTemplateHandler template.ListTemplateHandler
	// ResourceListVcdResourceHandler sets the operation handler for the list vcd resource operation
	ResourceListVcdResourceHandler resource.ListVcdResourceHandler
	// ResourceListVsphereResourceHandler sets the operation handler for the list vsphere resource operation
	ResourceListVsphereResourceHandler resource.ListVsphereResourceHandler
	// UserRegisterUserHandler sets the operation handler for the register user operation
	UserRegisterUserHandler user.RegisterUserHandler
	// PolicyRemovePolicyHandler sets the operation handler for the remove policy operation
	PolicyRemovePolicyHandler policy.RemovePolicyHandler
	// PolicyUpdatePolicyHandler sets the operation handler for the update policy operation
	PolicyUpdatePolicyHandler policy.UpdatePolicyHandler
	// ProjectUpdateProjectHandler sets the operation handler for the update project operation
	ProjectUpdateProjectHandler project.UpdateProjectHandler
	// ResourceUpdateResourceHandler sets the operation handler for the update resource operation
	ResourceUpdateResourceHandler resource.UpdateResourceHandler
	// UsageUpdateResourceUsageHandler sets the operation handler for the update resource usage operation
	UsageUpdateResourceUsageHandler usage.UpdateResourceUsageHandler
	// UserUpdateUserProfileHandler sets the operation handler for the update user profile operation
	UserUpdateUserProfileHandler user.UpdateUserProfileHandler
	// UserUserLoginHandler sets the operation handler for the user login operation
	UserUserLoginHandler user.UserLoginHandler
	// ResourceValidateVcdResourceHandler sets the operation handler for the validate vcd resource operation
	ResourceValidateVcdResourceHandler resource.ValidateVcdResourceHandler
	// ResourceValidateVsphereResourceHandler sets the operation handler for the validate vsphere resource operation
	ResourceValidateVsphereResourceHandler resource.ValidateVsphereResourceHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *CloudTidesAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *CloudTidesAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *CloudTidesAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *CloudTidesAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *CloudTidesAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *CloudTidesAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *CloudTidesAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *CloudTidesAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *CloudTidesAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the CloudTidesAPI
func (o *CloudTidesAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.PolicyAddPolicyHandler == nil {
		unregistered = append(unregistered, "policy.AddPolicyHandler")
	}
	if o.ProjectAddProjectHandler == nil {
		unregistered = append(unregistered, "project.AddProjectHandler")
	}
	if o.UsageAddResourceUsageHandler == nil {
		unregistered = append(unregistered, "usage.AddResourceUsageHandler")
	}
	if o.TemplateAddTemplateHandler == nil {
		unregistered = append(unregistered, "template.AddTemplateHandler")
	}
	if o.UsageAddVMUsageHandler == nil {
		unregistered = append(unregistered, "usage.AddVMUsageHandler")
	}
	if o.ResourceAddVcdResourceHandler == nil {
		unregistered = append(unregistered, "resource.AddVcdResourceHandler")
	}
	if o.ResourceAddVsphereResourceHandler == nil {
		unregistered = append(unregistered, "resource.AddVsphereResourceHandler")
	}
	if o.ProjectDeleteProjectHandler == nil {
		unregistered = append(unregistered, "project.DeleteProjectHandler")
	}
	if o.UsageDeleteResourceUsageHandler == nil {
		unregistered = append(unregistered, "usage.DeleteResourceUsageHandler")
	}
	if o.TemplateDeleteTemplateHandler == nil {
		unregistered = append(unregistered, "template.DeleteTemplateHandler")
	}
	if o.ResourceDeleteVcdResourceHandler == nil {
		unregistered = append(unregistered, "resource.DeleteVcdResourceHandler")
	}
	if o.ResourceDestroyVMHandler == nil {
		unregistered = append(unregistered, "resource.DestroyVMHandler")
	}
	if o.UsageGetPastUsageHandler == nil {
		unregistered = append(unregistered, "usage.GetPastUsageHandler")
	}
	if o.PolicyGetPolicyHandler == nil {
		unregistered = append(unregistered, "policy.GetPolicyHandler")
	}
	if o.UsageGetResourceUsageHandler == nil {
		unregistered = append(unregistered, "usage.GetResourceUsageHandler")
	}
	if o.UserGetUserProfileHandler == nil {
		unregistered = append(unregistered, "user.GetUserProfileHandler")
	}
	if o.ResourceGetVcdResourceHandler == nil {
		unregistered = append(unregistered, "resource.GetVcdResourceHandler")
	}
	if o.PolicyListPolicyHandler == nil {
		unregistered = append(unregistered, "policy.ListPolicyHandler")
	}
	if o.ProjectListProjectHandler == nil {
		unregistered = append(unregistered, "project.ListProjectHandler")
	}
	if o.TemplateListTemplateHandler == nil {
		unregistered = append(unregistered, "template.ListTemplateHandler")
	}
	if o.ResourceListVcdResourceHandler == nil {
		unregistered = append(unregistered, "resource.ListVcdResourceHandler")
	}
	if o.ResourceListVsphereResourceHandler == nil {
		unregistered = append(unregistered, "resource.ListVsphereResourceHandler")
	}
	if o.UserRegisterUserHandler == nil {
		unregistered = append(unregistered, "user.RegisterUserHandler")
	}
	if o.PolicyRemovePolicyHandler == nil {
		unregistered = append(unregistered, "policy.RemovePolicyHandler")
	}
	if o.PolicyUpdatePolicyHandler == nil {
		unregistered = append(unregistered, "policy.UpdatePolicyHandler")
	}
	if o.ProjectUpdateProjectHandler == nil {
		unregistered = append(unregistered, "project.UpdateProjectHandler")
	}
	if o.ResourceUpdateResourceHandler == nil {
		unregistered = append(unregistered, "resource.UpdateResourceHandler")
	}
	if o.UsageUpdateResourceUsageHandler == nil {
		unregistered = append(unregistered, "usage.UpdateResourceUsageHandler")
	}
	if o.UserUpdateUserProfileHandler == nil {
		unregistered = append(unregistered, "user.UpdateUserProfileHandler")
	}
	if o.UserUserLoginHandler == nil {
		unregistered = append(unregistered, "user.UserLoginHandler")
	}
	if o.ResourceValidateVcdResourceHandler == nil {
		unregistered = append(unregistered, "resource.ValidateVcdResourceHandler")
	}
	if o.ResourceValidateVsphereResourceHandler == nil {
		unregistered = append(unregistered, "resource.ValidateVsphereResourceHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *CloudTidesAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *CloudTidesAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *CloudTidesAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *CloudTidesAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *CloudTidesAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *CloudTidesAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the cloud tides API
func (o *CloudTidesAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *CloudTidesAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/policy"] = policy.NewAddPolicy(o.context, o.PolicyAddPolicyHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/project"] = project.NewAddProject(o.context, o.ProjectAddProjectHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/usage"] = usage.NewAddResourceUsage(o.context, o.UsageAddResourceUsageHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/template"] = template.NewAddTemplate(o.context, o.TemplateAddTemplateHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/usage/vm"] = usage.NewAddVMUsage(o.context, o.UsageAddVMUsageHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/resource/vcd"] = resource.NewAddVcdResource(o.context, o.ResourceAddVcdResourceHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/resource/vsphere"] = resource.NewAddVsphereResource(o.context, o.ResourceAddVsphereResourceHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/project/{id}"] = project.NewDeleteProject(o.context, o.ProjectDeleteProjectHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/usage/{id}"] = usage.NewDeleteResourceUsage(o.context, o.UsageDeleteResourceUsageHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/template/{id}"] = template.NewDeleteTemplate(o.context, o.TemplateDeleteTemplateHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/resource/vcd/{id}"] = resource.NewDeleteVcdResource(o.context, o.ResourceDeleteVcdResourceHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/resource/destroy_vm"] = resource.NewDestroyVM(o.context, o.ResourceDestroyVMHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/usage/past/{id}"] = usage.NewGetPastUsage(o.context, o.UsageGetPastUsageHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/policy/{id}"] = policy.NewGetPolicy(o.context, o.PolicyGetPolicyHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/usage/{id}"] = usage.NewGetResourceUsage(o.context, o.UsageGetResourceUsageHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/profile"] = user.NewGetUserProfile(o.context, o.UserGetUserProfileHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/resource/vcd/{id}"] = resource.NewGetVcdResource(o.context, o.ResourceGetVcdResourceHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/policy"] = policy.NewListPolicy(o.context, o.PolicyListPolicyHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/project"] = project.NewListProject(o.context, o.ProjectListProjectHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/template"] = template.NewListTemplate(o.context, o.TemplateListTemplateHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/resource/vcd"] = resource.NewListVcdResource(o.context, o.ResourceListVcdResourceHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/resource/vsphere"] = resource.NewListVsphereResource(o.context, o.ResourceListVsphereResourceHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/register"] = user.NewRegisterUser(o.context, o.UserRegisterUserHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/policy/{id}"] = policy.NewRemovePolicy(o.context, o.PolicyRemovePolicyHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/policy/{id}"] = policy.NewUpdatePolicy(o.context, o.PolicyUpdatePolicyHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/project/{id}"] = project.NewUpdateProject(o.context, o.ProjectUpdateProjectHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/resource/{id}"] = resource.NewUpdateResource(o.context, o.ResourceUpdateResourceHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/usage/{id}"] = usage.NewUpdateResourceUsage(o.context, o.UsageUpdateResourceUsageHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/users/profile"] = user.NewUpdateUserProfile(o.context, o.UserUpdateUserProfileHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/login"] = user.NewUserLogin(o.context, o.UserUserLoginHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/resource/vcd/validate"] = resource.NewValidateVcdResource(o.context, o.ResourceValidateVcdResourceHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/resource/vsphere/validate"] = resource.NewValidateVsphereResource(o.context, o.ResourceValidateVsphereResourceHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *CloudTidesAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *CloudTidesAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *CloudTidesAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *CloudTidesAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *CloudTidesAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
