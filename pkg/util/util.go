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
const GpuDevice = "nvidia.com/"

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

func IsNvidiaGpuVmi(vmi *v1.VirtualMachineInstance) bool {
	for key := range vmi.Spec.Domain.Resources.Requests {
		if strings.HasPrefix(string(key), GpuDevice) {
			return true
		}
	}

	for key := range vmi.Spec.Domain.Resources.Limits {
		if strings.HasPrefix(string(key), GpuDevice) {
			return true
		}
	}
	return false
}
