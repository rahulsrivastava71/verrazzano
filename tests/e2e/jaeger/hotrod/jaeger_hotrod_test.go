// Copyright (c) 2022, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package hotrod

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/verrazzano/verrazzano/pkg/test/framework"
	"github.com/verrazzano/verrazzano/pkg/test/framework/metrics"
	"github.com/verrazzano/verrazzano/tests/e2e/jaeger"
	"github.com/verrazzano/verrazzano/tests/e2e/pkg"
	"time"
)

const (
	shortPollingInterval = 10 * time.Second
	shortWaitTimeout     = 5 * time.Minute
)

const (
	testAppComponentFilePath     = "testdata/jaeger/hotrod/hotrod-tracing-comp.yaml"
	testAppConfigurationFilePath = "testdata/jaeger/hotrod/hotrod-tracing-app.yaml"
)

var (
	t                  = framework.NewTestFramework("jaeger")
	generatedNamespace = pkg.GenerateNamespace("hotrod-tracing")
	expectedPodsHotrod = []string{"hotrod-workload"}
	beforeSuitePassed  = false
	start              = time.Now()
	hotrodServiceName  = fmt.Sprintf("hotrod.%s", generatedNamespace)
)

var _ = t.BeforeSuite(func() {
	start = time.Now()
	jaeger.DeployApplication(namespace, testAppComponentFilePath, testAppConfigurationFilePath, expectedPodsHotrod)
	beforeSuitePassed = true
	metrics.Emit(t.Metrics.With("deployment_elapsed_time", time.Since(start).Milliseconds()))
})

var _ = t.AfterSuite(func() {
	if !beforeSuitePassed {
		pkg.ExecuteBugReport(namespace)
	}
	// undeploy the application here
	start := time.Now()

	jaeger.UndeployApplication(namespace, testAppComponentFilePath, testAppConfigurationFilePath, expectedPodsHotrod)
	metrics.Emit(t.Metrics.With("undeployment_elapsed_time", time.Since(start).Milliseconds()))
})

var _ = t.Describe("Hotrod App with Jaeger Traces", Label("f:jaeger.hotrod-workload"), func() {
	t.Context("after successful installation", func() {
		// GIVEN the Jaeger Operator is enabled and a sample application is installed,
		// WHEN we check for traces for that service,
		// THEN we are able to get the traces
		jaeger.WhenJaegerOperatorEnabledIt(t, "traces for the hotrod app should be available when queried from Jaeger", func() {
			validatorFn := pkg.ValidateApplicationTraces(start, hotrodServiceName)
			Eventually(validatorFn).WithPolling(shortPollingInterval).WithTimeout(shortWaitTimeout).Should(BeTrue())
		})

		// GIVEN the Jaeger Operator component is enabled,
		// WHEN a sample application is installed,
		// THEN the traces are found in OpenSearch Backend
		jaeger.WhenJaegerOperatorEnabledIt(t, "traces for the hotrod app should be available in the OS backend storage.", func() {
			validatorFn := pkg.ValidateApplicationTracesInOS(start, hotrodServiceName)
			Eventually(validatorFn).WithPolling(shortPollingInterval).WithTimeout(shortWaitTimeout).Should(BeTrue())
		})
	})

})
