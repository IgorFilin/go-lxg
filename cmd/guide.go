/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// linuxguideCmd represents the linuxguide command
var linuxguideCmd = &cobra.Command{
	Use:   "guide",
	Short: "Мини гайд по командам линукса",
	Long: `Команды линукса разбиты по темам, выберете нужную и вы получите список`,
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "Выберите тему: ",
			Items: []string{
				"Управление доступами/правами",
				"Команды пользователей",
				"Сетевые команды", 
				"Работа с файлами", 
				"Тренажёр",
			},
		}
	
		index, _, err := prompt.Run()
	
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
	    GetThemeInfo(index)
	},
}

func init() {
	rootCmd.AddCommand(linuxguideCmd)
	// linuxguideCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
