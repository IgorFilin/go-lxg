/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// timezoneCmd represents the timezone command
var timezoneCmd = &cobra.Command{
	Use:   "timezone",
	Short: "Возвращает время определённой географической зоны",
	Long: "Возвращает время определённой географической зоны принимая один аргумент зону",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		timezone := args[0]
		flag, _ := cmd.Flags().GetString("format")
		fmt.Println(flag)
		timeNow, err := GetTimeZone(timezone, flag)
		if err != nil {
			log.Fatalln("Неверная временная зона")
		}
		fmt.Println(timeNow)
	},
}

func init() {
	rootCmd.AddCommand(timezoneCmd)
	timezoneCmd.Flags().String("format", "", "Выводит время в формате yyyy-mm-dd")
}
