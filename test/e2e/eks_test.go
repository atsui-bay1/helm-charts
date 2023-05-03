package e2e

import (
	"context"
	"fmt"
	"github.com/DataDog/datadog-agent/test/new-e2e/runner"
	"github.com/DataDog/datadog-agent/test/new-e2e/utils/infra"
	"github.com/DataDog/helm-charts/test/utils"
	"github.com/DataDog/test-infra-definitions/aws/scenarios/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAgentOnEKS(t *testing.T) {
	// Creating the stack
	config := utils.SetupConfig()
	stackConfig := runner.ConfigMap{
		"ddinfra:aws/eks/linuxNodeGroup":             auto.ConfigValue{Value: "false"},
		"ddinfra:aws/eks/linuxARMNodeGroup":          auto.ConfigValue{Value: "false"},
		"ddinfra:aws/eks/linuxBottlerocketNodeGroup": auto.ConfigValue{Value: "false"},
		"ddinfra:aws/eks/windowsLTSCNodeGroup":       auto.ConfigValue{Value: "false"},
		"pulumi:disable-default-providers":           auto.ConfigValue{Value: "[]"},
		"ddagent:deploy":                             auto.ConfigValue{Value: "false"},
	}
	stackConfig.Merge(config)

	_, stackOutput, err := infra.GetStackManager().GetStack(context.Background(), "helm-charts-eks-cluster", stackConfig, eks.Run, false)
	fmt.Println(stackOutput)
	require.NoError(t, err)
}
