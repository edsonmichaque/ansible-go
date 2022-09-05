package cmd

import "github.com/spf13/cobra"

type AnsibleCmd struct {
	module             string
	args               string
	user               string
	privateKeyFile     string
	keyFile            string
	become             bool
	becomeMethod       string
	becomeUser         string
	connection         string
	timeout            int
	inventory          string
	background         int
	askVaultPassword   bool
	becomePasswordFile string
	listHosts          bool
	poll               int
	syntaxCheck        bool
	check              bool
	diff               bool
	askBecomePass      bool
	taskTimeout        int
}

func (a *AnsibleCmd) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "ansible",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.Flags().StringVarP(&a.module, "module", "m", "", "module")
	cmd.Flags().StringVarP(&a.args, "args", "a", "", "args")
	cmd.Flags().StringVarP(&a.user, "user", "u", "", "user")
	cmd.Flags().StringVarP(&a.connection, "connection", "c", "", "connection")
	cmd.Flags().IntVarP(&a.timeout, "timeout", "T", 0, "timeout")
	cmd.Flags().IntVar(&a.taskTimeout, "task-timeout", 0, "task timeout")
	cmd.Flags().IntVarP(&a.background, "background", "B", 0, "background")
	cmd.Flags().IntVarP(&a.poll, "poll", "P", 0, "poll")
	cmd.Flags().StringVar(&a.keyFile, "key-file", "", "key file")
	cmd.Flags().StringVar(&a.becomePasswordFile, "become-password-file", "", "become password file")
	cmd.Flags().StringVar(&a.privateKeyFile, "private-key-file", "", "key file")
	cmd.Flags().StringVar(&a.inventory, "inventory", "", "inventory")
	cmd.Flags().BoolVarP(&a.become, "become", "b", false, "become")
	cmd.Flags().BoolVar(&a.listHosts, "list-hosts", false, "list hosts")
	cmd.Flags().BoolVar(&a.syntaxCheck, "syntax-check", false, "syntax check")
	cmd.Flags().BoolVarP(&a.check, "check", "C", false, "check")
	cmd.Flags().BoolVarP(&a.diff, "diff", "D", false, "diff")
	cmd.Flags().BoolVar(&a.askBecomePass, "ask-become-pass", false, "ask become pass")
	cmd.Flags().BoolVar(&a.become, "ask-vault-password", false, "ask vault password")
	cmd.Flags().StringVar(&a.becomeMethod, "become-method", "", "become method")
	cmd.Flags().StringVar(&a.becomeMethod, "become-user", "", "become user")

	return cmd
}
