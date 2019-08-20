package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	v1 "kubevirt.io/client-go/api/v1"
)

const ServiceAccountNamespaceFile = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"
const namespaceKubevirt = "kubevirt"
const ExtensionAPIServerAuthenticationConfigMap = "extension-apiserver-authentication"
const RequestHeaderClientCAFileKey = "requestheader-client-ca-file"
const VirtShareDir = "/var/run/kubevirt"
const VirtLibDir = "/var/lib/kubevirt"

func GetNamespace() (string, error) {
	if data, err := ioutil.ReadFile(ServiceAccountNamespaceFile); err == nil {
		if ns := strings.TrimSpace(string(data)); len(ns) > 0 {
			return ns, nil
		}
	} else if err != nil && !os.IsNotExist(err) {
		return "", fmt.Errorf("failed to determine namespace from %s: %v", ServiceAccountNamespaceFile, err)
	}
	return namespaceKubevirt, nil
}

// Check if a VMI spec requests GPU
func IsGpuVmi(vmi *v1.VirtualMachineInstance) bool {
	if vmi.Spec.Domain.Devices.Gpus != nil {
		return true
	}
	return false
}
