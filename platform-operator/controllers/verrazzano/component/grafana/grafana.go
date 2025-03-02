// Copyright (c) 2022, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package grafana

import (
	"fmt"

	"github.com/verrazzano/verrazzano/pkg/k8s/status"
	"github.com/verrazzano/verrazzano/platform-operator/controllers/verrazzano/component/common"
	"github.com/verrazzano/verrazzano/platform-operator/controllers/verrazzano/component/spi"
	"k8s.io/apimachinery/pkg/types"
)

const grafanaDeployment = "vmi-system-grafana"

// isGrafanaInstalled checks that the Grafana deployment exists
func isGrafanaInstalled(ctx spi.ComponentContext) bool {
	prefix := newPrefix(ctx.GetComponent())
	deployments := newDeployments()
	return status.DoDeploymentsExist(ctx.Log(), ctx.Client(), deployments, 1, prefix)
}

// isGrafanaReady checks that the deployment has the minimum number of replicas available and
// that the admin secret is ready
func isGrafanaReady(ctx spi.ComponentContext) bool {
	prefix := newPrefix(ctx.GetComponent())
	deployments := newDeployments()
	return status.DeploymentsAreReady(ctx.Log(), ctx.Client(), deployments, 1, prefix) && common.IsGrafanaAdminSecretReady(ctx)
}

// newPrefix creates a component prefix string
func newPrefix(component string) string {
	return fmt.Sprintf("Component %s", component)
}

// creates a slice of NamespacedName with the Grafana deployment name
func newDeployments() []types.NamespacedName {
	return []types.NamespacedName{
		{
			Name:      grafanaDeployment,
			Namespace: ComponentNamespace,
		},
	}
}
