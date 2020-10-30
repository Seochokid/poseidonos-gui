package cmd

import (
	"pnconnector/src/errors"
	iBoFOS "pnconnector/src/routers/m9k/api/ibofos"
	"pnconnector/src/routers/m9k/model"
	"github.com/spf13/cobra"
)

var RebuildCommand = map[string]func(string, interface{}) (model.Request, model.Response, error){
    "perf_impact":      iBoFOS.PerfImpact,
}

var rebuildCmd = &cobra.Command{
	Use:   "rebuild [msg]",
	Short: "Request for rebuild msg to Poseidon OS",
	Long: `Request for rebuild msg to Poseidon OS and get a response fommated by JSON.

Available msg list :

[Category] : [msg]            : [description]                                                    : [example of flag]

rebuild    : perf_impact      : Set rebuild performance impact on USER I/O.                      : not needed


If you want to input multiple flag parameter, you have to seperate with ",". 
For example, "-d dev1,dev2,dev3". seperation by space is not allowed.


You can set ip and port number for connect to Poseidon OS using config.yaml or flags.
Default value is as below.

IP   : 127.0.0.1
Port : 18716


	  `,
	Args: func(cmd *cobra.Command, args []string) error {

		if len(args) != 1 {
			return errors.New("need an one msg !!!")
		}

		_, exist := RebuildCommand[args[0]]

		if !exist {
			return errors.New("not available msg !!!")
		}

		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		Send(cmd, args)
	},
}

func init() {

	rootCmd.AddCommand(rebuildCmd)
	rebuildCmd.PersistentFlags().StringVarP(&level, "level", "l", "", "set level")
}
