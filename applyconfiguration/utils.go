/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
	v1 "sigs.k8s.io/gateway-api/apis/v1"
	v1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	v1alpha3 "sigs.k8s.io/gateway-api/apis/v1alpha3"
	v1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
	v1alpha1 "sigs.k8s.io/gateway-api/apisx/v1alpha1"
	apisv1 "sigs.k8s.io/gateway-api/applyconfiguration/apis/v1"
	apisv1alpha2 "sigs.k8s.io/gateway-api/applyconfiguration/apis/v1alpha2"
	apisv1alpha3 "sigs.k8s.io/gateway-api/applyconfiguration/apis/v1alpha3"
	apisv1beta1 "sigs.k8s.io/gateway-api/applyconfiguration/apis/v1beta1"
	apisxv1alpha1 "sigs.k8s.io/gateway-api/applyconfiguration/apisx/v1alpha1"
	internal "sigs.k8s.io/gateway-api/applyconfiguration/internal"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=gateway.networking.k8s.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("AllowedListeners"):
		return &apisv1.AllowedListenersApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AllowedRoutes"):
		return &apisv1.AllowedRoutesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BackendObjectReference"):
		return &apisv1.BackendObjectReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BackendRef"):
		return &apisv1.BackendRefApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CommonRouteSpec"):
		return &apisv1.CommonRouteSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CookieConfig"):
		return &apisv1.CookieConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Fraction"):
		return &apisv1.FractionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FrontendTLSValidation"):
		return &apisv1.FrontendTLSValidationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Gateway"):
		return &apisv1.GatewayApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GatewayBackendTLS"):
		return &apisv1.GatewayBackendTLSApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GatewayClass"):
		return &apisv1.GatewayClassApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GatewayClassSpec"):
		return &apisv1.GatewayClassSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GatewayClassStatus"):
		return &apisv1.GatewayClassStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GatewayInfrastructure"):
		return &apisv1.GatewayInfrastructureApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GatewaySpec"):
		return &apisv1.GatewaySpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GatewaySpecAddress"):
		return &apisv1.GatewaySpecAddressApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GatewayStatus"):
		return &apisv1.GatewayStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GatewayStatusAddress"):
		return &apisv1.GatewayStatusAddressApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GatewayTLSConfig"):
		return &apisv1.GatewayTLSConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GRPCBackendRef"):
		return &apisv1.GRPCBackendRefApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GRPCHeaderMatch"):
		return &apisv1.GRPCHeaderMatchApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GRPCMethodMatch"):
		return &apisv1.GRPCMethodMatchApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GRPCRoute"):
		return &apisv1.GRPCRouteApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GRPCRouteFilter"):
		return &apisv1.GRPCRouteFilterApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GRPCRouteMatch"):
		return &apisv1.GRPCRouteMatchApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GRPCRouteRule"):
		return &apisv1.GRPCRouteRuleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GRPCRouteSpec"):
		return &apisv1.GRPCRouteSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GRPCRouteStatus"):
		return &apisv1.GRPCRouteStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPBackendRef"):
		return &apisv1.HTTPBackendRefApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPCORSFilter"):
		return &apisv1.HTTPCORSFilterApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPHeader"):
		return &apisv1.HTTPHeaderApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPHeaderFilter"):
		return &apisv1.HTTPHeaderFilterApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPHeaderMatch"):
		return &apisv1.HTTPHeaderMatchApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPPathMatch"):
		return &apisv1.HTTPPathMatchApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPPathModifier"):
		return &apisv1.HTTPPathModifierApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPQueryParamMatch"):
		return &apisv1.HTTPQueryParamMatchApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPRequestMirrorFilter"):
		return &apisv1.HTTPRequestMirrorFilterApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPRequestRedirectFilter"):
		return &apisv1.HTTPRequestRedirectFilterApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPRoute"):
		return &apisv1.HTTPRouteApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPRouteFilter"):
		return &apisv1.HTTPRouteFilterApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPRouteMatch"):
		return &apisv1.HTTPRouteMatchApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPRouteRetry"):
		return &apisv1.HTTPRouteRetryApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPRouteRule"):
		return &apisv1.HTTPRouteRuleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPRouteSpec"):
		return &apisv1.HTTPRouteSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPRouteStatus"):
		return &apisv1.HTTPRouteStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPRouteTimeouts"):
		return &apisv1.HTTPRouteTimeoutsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPURLRewriteFilter"):
		return &apisv1.HTTPURLRewriteFilterApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Listener"):
		return &apisv1.ListenerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ListenerNamespaces"):
		return &apisv1.ListenerNamespacesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ListenerStatus"):
		return &apisv1.ListenerStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("LocalObjectReference"):
		return &apisv1.LocalObjectReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("LocalParametersReference"):
		return &apisv1.LocalParametersReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ObjectReference"):
		return &apisv1.ObjectReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ParametersReference"):
		return &apisv1.ParametersReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ParentReference"):
		return &apisv1.ParentReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RouteGroupKind"):
		return &apisv1.RouteGroupKindApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RouteNamespaces"):
		return &apisv1.RouteNamespacesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RouteParentStatus"):
		return &apisv1.RouteParentStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RouteStatus"):
		return &apisv1.RouteStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SecretObjectReference"):
		return &apisv1.SecretObjectReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SessionPersistence"):
		return &apisv1.SessionPersistenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SupportedFeature"):
		return &apisv1.SupportedFeatureApplyConfiguration{}

		// Group=gateway.networking.k8s.io, Version=v1alpha2
	case v1alpha2.SchemeGroupVersion.WithKind("GRPCRoute"):
		return &apisv1alpha2.GRPCRouteApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("LocalPolicyTargetReference"):
		return &apisv1alpha2.LocalPolicyTargetReferenceApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("LocalPolicyTargetReferenceWithSectionName"):
		return &apisv1alpha2.LocalPolicyTargetReferenceWithSectionNameApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("PolicyAncestorStatus"):
		return &apisv1alpha2.PolicyAncestorStatusApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("PolicyStatus"):
		return &apisv1alpha2.PolicyStatusApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("ReferenceGrant"):
		return &apisv1alpha2.ReferenceGrantApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("TCPRoute"):
		return &apisv1alpha2.TCPRouteApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("TCPRouteRule"):
		return &apisv1alpha2.TCPRouteRuleApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("TCPRouteSpec"):
		return &apisv1alpha2.TCPRouteSpecApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("TCPRouteStatus"):
		return &apisv1alpha2.TCPRouteStatusApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("TLSRoute"):
		return &apisv1alpha2.TLSRouteApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("TLSRouteRule"):
		return &apisv1alpha2.TLSRouteRuleApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("TLSRouteSpec"):
		return &apisv1alpha2.TLSRouteSpecApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("TLSRouteStatus"):
		return &apisv1alpha2.TLSRouteStatusApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("UDPRoute"):
		return &apisv1alpha2.UDPRouteApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("UDPRouteRule"):
		return &apisv1alpha2.UDPRouteRuleApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("UDPRouteSpec"):
		return &apisv1alpha2.UDPRouteSpecApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("UDPRouteStatus"):
		return &apisv1alpha2.UDPRouteStatusApplyConfiguration{}

		// Group=gateway.networking.k8s.io, Version=v1alpha3
	case v1alpha3.SchemeGroupVersion.WithKind("BackendTLSPolicy"):
		return &apisv1alpha3.BackendTLSPolicyApplyConfiguration{}
	case v1alpha3.SchemeGroupVersion.WithKind("BackendTLSPolicySpec"):
		return &apisv1alpha3.BackendTLSPolicySpecApplyConfiguration{}
	case v1alpha3.SchemeGroupVersion.WithKind("BackendTLSPolicyValidation"):
		return &apisv1alpha3.BackendTLSPolicyValidationApplyConfiguration{}
	case v1alpha3.SchemeGroupVersion.WithKind("SubjectAltName"):
		return &apisv1alpha3.SubjectAltNameApplyConfiguration{}
	case v1alpha3.SchemeGroupVersion.WithKind("TLSRoute"):
		return &apisv1alpha3.TLSRouteApplyConfiguration{}
	case v1alpha3.SchemeGroupVersion.WithKind("TLSRouteSpec"):
		return &apisv1alpha3.TLSRouteSpecApplyConfiguration{}

		// Group=gateway.networking.k8s.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithKind("Gateway"):
		return &apisv1beta1.GatewayApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("GatewayClass"):
		return &apisv1beta1.GatewayClassApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("HTTPRoute"):
		return &apisv1beta1.HTTPRouteApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ReferenceGrant"):
		return &apisv1beta1.ReferenceGrantApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ReferenceGrantFrom"):
		return &apisv1beta1.ReferenceGrantFromApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ReferenceGrantSpec"):
		return &apisv1beta1.ReferenceGrantSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ReferenceGrantTo"):
		return &apisv1beta1.ReferenceGrantToApplyConfiguration{}

		// Group=gateway.networking.x-k8s.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("BackendTrafficPolicySpec"):
		return &apisxv1alpha1.BackendTrafficPolicySpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("BudgetDetails"):
		return &apisxv1alpha1.BudgetDetailsApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ListenerEntry"):
		return &apisxv1alpha1.ListenerEntryApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ListenerEntryStatus"):
		return &apisxv1alpha1.ListenerEntryStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ListenerSetSpec"):
		return &apisxv1alpha1.ListenerSetSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ListenerSetStatus"):
		return &apisxv1alpha1.ListenerSetStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ParentGatewayReference"):
		return &apisxv1alpha1.ParentGatewayReferenceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RequestRate"):
		return &apisxv1alpha1.RequestRateApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RetryConstraint"):
		return &apisxv1alpha1.RetryConstraintApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("XBackendTrafficPolicy"):
		return &apisxv1alpha1.XBackendTrafficPolicyApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("XListenerSet"):
		return &apisxv1alpha1.XListenerSetApplyConfiguration{}

	}
	return nil
}

func NewTypeConverter(scheme *runtime.Scheme) *testing.TypeConverter {
	return &testing.TypeConverter{Scheme: scheme, TypeResolver: internal.Parser()}
}
