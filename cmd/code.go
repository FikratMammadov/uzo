/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"uzo/utils"
)

var File string

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code <zip_file_name>",
	Short: "It will open the directory in Visual Studio Code",
	Long: `This command will help to open the unzipped folder to Visual Studio Code.
In order for this command to work, Visual Studio code should be installed in your system`,
	//Args: cobra.ExactArgs(1),
	Args: func(cmd *cobra.Command, args []string) error {
		if File == "" && len(args) < 1 {
			return errors.New("accepts 1 argument")
		}
		return nil
	},
	Example: `uzo code demo.zip
uzo code \Downloads\demo.zip`,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		var fileName string
		var err error
		var argument string

		if File != "" {
			argument = File
		} else {
			argument = args[0]
		}

		fileExists, err := utils.FileExists(argument)

		if err != nil {
			fmt.Println(err.Error())
		}

		if fileExists {
			fileName, err = filepath.Abs(argument)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Printf("File %v does not exist", argument)
			return
		}

		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}

		utils.Unzip(fileName, wd)
		os.Chdir(utils.FilenameWithoutExtension(fileName))

		wd, err = os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}

		commandCode := exec.Command("code", wd)
		err = commandCode.Run()

		if err != nil {
			fmt.Println("VS Code executable file not fount in %PATH%")
		}
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	codeCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "A file name to unzip and open in IDE")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
