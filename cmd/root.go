package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "dev-cli",
	Short: "客厅开发脚手架",
	Long: `客厅开发脚手架，方便在本地仓库下，通过命令行关联本地代码开发以及tapd单扭转流程, 规避需要手动扭转导致不及时的问题:
	覆盖启动需求开发，启动bugfix，提交测试，提交mr等操作`,
}

func Execute() {
	fmt.Println("falg is 4")
	isDone := Prepare()
	if isDone {
		os.Exit(0)
	}
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func Prepare() bool {
	return UpdateLatestVersion()
}
