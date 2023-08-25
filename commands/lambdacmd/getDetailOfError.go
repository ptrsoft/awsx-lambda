/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package lambdacmd

import (
	"fmt"

	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/controllers"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetDetailOfErrorCmd = &cobra.Command{
	Use:   "errorDetail",
	Short: "to get details of error in lambda",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth := authenticater.ChildCommandAuth(cmd)
		function, _ := cmd.Flags().GetString("function")

		if authFlag {
			controllers.LambdaDetailsErrorController(function, *clientAuth)

		}

	},
}

func init() {

	GetDetailOfErrorCmd.Flags().StringP("function", "f", "", "lambda function name")

	if err := GetConfigDataCmd.MarkFlagRequired("function"); err != nil {
		fmt.Println("--function is required", err)
	}

}
