package ansible

type Task struct {
	Name              *string                `json:"name" yaml:"name"`
	Action            *string                `json:"action" yaml:"action"`
	AnyErrorsFatal    *bool                  `json:"any_errors_fatal" yaml:"any_errors_fatal"`
	Args              map[string]interface{} `json:"args" yaml:"args"`
	Async             interface{}            `json:"async" yaml:"async"`
	Become            *bool                  `json:"become" yaml:"become"`
	BecomeExe         *string                `json:"become_exe" yaml:"become_exe"`
	BecomeFlags       *string                `json:"become_flags" yaml:"become_flags"`
	BecomeMethod      *string                `json:"become_method" yaml:"become_method"`
	BecomeUser        *string                `json:"become_user" yaml:"become_user"`
	ChangedWhen       *string                `json:"changed_when" yaml:"changed_when"`
	CheckMode         *string                `json:"check_mode" yaml:"check_mode"`
	Collections       *string                `json:"collections" yaml:"collections"`
	Connection        *string                `json:"connection yaml:"connection"`
	Debugger          *bool                  `json:"debugger" yaml:"debugger"`
	Delay             *int                   `json:"delay yaml:"delay"`
	DelegateFacts     *bool                  `json:"delegate_facts" yaml:"delegate_facts"`
	DelegateTo        *string                `json:"delegate_to" yaml:"delegate_to"`
	Diff              *bool                  `json:"diff" yaml:"diff"`
	Environment       map[string]interface{} `json:"environment" yaml:"environment"`
	IgnoreErrors      *bool                  `json:"ignore_errors" yaml:"ignore_errors"`
	IgnoreUnrecheable *bool                  `json:"ignore_unrecheable" yaml:"ignore_unrecheable"`
	LocalAction       *bool                  `json:"local_action" yaml:"local_action"`
	Loop              interface{}            `json:"loop" yaml:"loop"`
	LoopControl       interface{}            `json:"loop_control" yaml:"loop_control"`
	ModuleDefaults    interface{}            `json:"module_defaults" yaml:"module_defaults"`
	NoLog             *bool                  `json:"no_log" yaml:"no_log"`
	Notify            interface{}            `json:"notify" yaml:"notify"`
	Polll             interface{}            `json:"poll" yaml:"poll"`
	Port              interface{}            `json:"port" yaml:"port"`
	Register          interface{}            `json:"register" yaml:"register"`
	RemoteUser        *string                `json:"remote_user" yaml:"remote_user"`
	Retries           *int                   `json:"retries yaml:"retries"`
	Tags              interface{}            `json:"tags" yaml:"tags"`
	Until             interface{}            `json:"until" yaml:"until"`
	When              interface{}            `json:"when" yaml:"when"`
	Throttle          interface{}            `json:"throttle" yaml:"throttle"`
	Timeout           interface{}            `json:"timeout" yaml:"timeout"`
	Vars              map[string]interface{} `json:"vars" yaml:"vars"`
}
