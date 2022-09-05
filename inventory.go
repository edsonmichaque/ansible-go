package ansible

type Inventory map[string]Host

type Host struct {
	Hosts    map[string]interface{} `json:"hosts" yaml:"hosts" ini:"hosts"`
	Children map[string]Host        `json:"children" yaml:"children" ini:"children"`
	Vars     map[string]string      `json:"vars" yaml:"vars" ini:"vars"`
}
