package vector

import (
	logging "github.com/openshift/cluster-logging-operator/apis/logging/v1"
	"github.com/openshift/cluster-logging-operator/internal/generator"
	"github.com/openshift/cluster-logging-operator/internal/generator/vector/output/kafka"
	corev1 "k8s.io/api/core/v1"
)

func Outputs(clspec *logging.ClusterLoggingSpec, secrets map[string]*corev1.Secret, clfspec *logging.ClusterLogForwarderSpec, op generator.Options) []generator.Element {
	outputs := []generator.Element{}
	route := PipelineToOutputs(clfspec, op)
	for _, o := range clfspec.Outputs {
		secret := secrets[o.Name]
		inputs := route[o.Name].List()
		if o.Type == logging.OutputTypeKafka {
			outputs = generator.MergeElements(outputs, kafka.Conf(o, inputs, secret, op))
		}
	}
	return outputs
}
