package types

type Metadata struct {
	Name string `yaml:"name"`
}

type Spec struct {
	Components []map[string]any `yaml:"components"`
}

type OAM struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

type BackstageMetadata struct {
	Name string `yaml:"name"`
}

type BackstageSpec struct {
	Type      string `yaml:"type"`
	Lifecycle string `yaml:"lifecycle"`
	Owner     string `yaml:"owner"`
	System    string `yaml:"system"`
}

type Backstage struct {
	ApiVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   BackstageMetadata `yaml:"metadata"`
	Spec       BackstageSpec     `yaml:"spec"`
}

func NewBackstage() *Backstage {
	return &Backstage{
		ApiVersion: "backstage.io/v1alpha1",
		Kind:       "Component",
		Metadata:   BackstageMetadata{},
		Spec: BackstageSpec{
			Lifecycle: "production",
			Owner:     "Fake-Team",
			System:    "Fake-System",
		},
	}
}
