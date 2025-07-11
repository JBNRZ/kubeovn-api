package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/JBNRZ/kubeovn-api/pkg/apis/kubeovn"
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: kubeovn.GroupName, Version: "v1"}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	// SchemeBuilder initializes a scheme builder
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	// AddToScheme is a global function that registers this API group & version to a scheme
	AddToScheme = SchemeBuilder.AddToScheme
)

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&IP{},
		&IPList{},
		&IPPool{},
		&IPPoolList{},
		&IptablesDnatRule{},
		&IptablesDnatRuleList{},
		&IptablesEIP{},
		&IptablesEIPList{},
		&IptablesFIPRule{},
		&IptablesFIPRuleList{},
		&IptablesSnatRule{},
		&IptablesSnatRuleList{},
		&OvnDnatRule{},
		&OvnDnatRuleList{},
		&OvnEip{},
		&OvnEipList{},
		&OvnFip{},
		&OvnFipList{},
		&OvnSnatRule{},
		&OvnSnatRuleList{},
		&ProviderNetwork{},
		&ProviderNetworkList{},
		&QoSPolicy{},
		&QoSPolicyList{},
		&SecurityGroup{},
		&SecurityGroupList{},
		&Subnet{},
		&SubnetList{},
		&SwitchLBRule{},
		&SwitchLBRuleList{},
		&Vip{},
		&VipList{},
		&Vlan{},
		&VlanList{},
		&Vpc{},
		&VpcList{},
		&VpcDns{},
		&VpcDnsList{},
		&VpcEgressGateway{},
		&VpcEgressGatewayList{},
		&VpcNatGateway{},
		&VpcNatGatewayList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
