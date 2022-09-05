package ansible

type Playbook []Play

type Play struct {
	Name              *string                `json:"name" yaml:"name"`
	Hosts             interface{}            `json:"hosts" yaml:"hosts"`
	AnyErrorsFatal    *bool                  `json:"any_errors_fatal" yaml:"any_errors_fatal"`
	Become            *bool                  `json:"become" yaml:"become"`
	RemoteUser        string                 `json:"remote_user" yaml:"remote_user"`
	BecomeExe         *string                `json:"become_exe" yaml:"become_exe"`
	BecomeMethod      *string                `json:"become_method" yaml:"become_method"`
	BecomeFlags       *string                `json:"become_flags" yaml:"become_flags"`
	CheckMode         *bool                  `json:"check_mode" yaml:"check_mode"`
	Collections       interface{}            `json:"collections" yaml:"collections"`
	Connection        interface{}            `json:"connection" yaml:"connection"`
	Debugger          *bool                  `json:"debug" yaml:"debug"`
	Diff              *bool                  `json:"diff" yaml:"diff"`
	Environment       map[string]interface{} `json:"environment" yaml:"environment"`
	FactPath          *string                `json:"fact_path" yaml:"fact_path"`
	ForceHandlers     *bool                  `json:"force_handlers" yaml:"force_handlers"`
	GatherFacts       *bool                  `json:"gather_facts" yaml:"gather_facts"`
	GatherSubset      *bool                  `json:"gather_subset" yaml:"gather_subset"`
	GatherTimeout     *bool                  `json:"gather_timeout" yaml:"gather_timeout"`
	Handlers          []interface{}          `json:"handlers" yaml:"handlers"`
	IgnoreErrors      *bool                  `json:"ignore_errors" yaml:"ignore_errors"`
	IgnoreUnrecheable *bool                  `json:"ignore_unrecheable" yaml:"ignore_unrecheable"`
	MaxfailPercentage *bool                  `json:"max_fail_percentage" yaml:"max_fail_percentage"`
	ModuleDefaults    interface{}            `json:"module_defaults" yaml:"module_defaults"`
	NoLog             *bool                  `json:"no_log" yaml:"no_log"`
	PostTasks         []interface{}          `json:"post_tasks" yaml:"post_tasks"`
	PreTasks          []interface{}          `json:"pre_tasks" yaml:"pre_tasks"`
	RunOnce           *bool                  `json:"run_once" yaml:"run_once"`
	Roles             []interface{}          `json:"roles" yaml:"roles"`
	Serial            interface{}            `json:"serial" yaml:"serial"`
	Strategy          interface{}            `json:"strategy" yaml:"strategy"`
	Tags              interface{}            `json:"tags" yaml:"tags"`
	Throttle          interface{}            `json:"throttle" yaml:"throttle"`
	Timeout           interface{}            `json:"timeout" yaml:"timeout"`
	Vars              map[string]interface{} `json:"vars" yaml:"vars"`
	VarsFiles         []interface{}          `json:"vars_files" yaml:"vars_files"`
	VarsPrompt        interface{}            `json:"vars_prompt" yaml:"vars_prompt"`

	Tasks []interface{} `json:"tasks" yaml:"tasks"`
}
