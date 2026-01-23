package types

import "sigs.k8s.io/kustomize/kyaml/fn/runtime/runtimeutil"

type FunctionType = string

const (
	TransformerType FunctionType = "transformer"
	GeneratorType   FunctionType = "generator"
)

type FunctionTransformer struct {
	// Path is the path to the function configuration
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
	// Spec is the function spec
	Spec runtimeutil.FunctionSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
	// Type is the function type. It can be "transformer" or "generator".
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

func (f *FunctionTransformer) GetType() FunctionType {
	return FunctionType(f.Type)
}

func (f *FunctionTransformer) IsTransformer() bool {
	return f.Type == string(TransformerType)
}

func (f *FunctionTransformer) IsGenerator() bool {
	return f.Type == string(GeneratorType)
}
