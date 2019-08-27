package util

import (
	v1 "kubevirt.io/client-go/api/v1"
)

const ExtensionAPIServerAuthenticationConfigMap = "extension-apiserver-authentication"
const RequestHeaderClientCAFileKey = "requestheader-client-ca-file"
const VirtShareDir = "/var/run/kubevirt"
const VirtLibDir = "/var/lib/kubevirt"

// Check if a VMI spec requests GPU
func IsGpuVmi(vmi *v1.VirtualMachineInstance) bool {
	if vmi.Spec.Domain.Devices.Gpus != nil {
		return true
	}
	return false
}
