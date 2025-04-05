package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pterm/pterm"
)

type PermissionStructure struct {
	Metadata struct {
		Title string `json:"title"`
		Version string  `json:"version"`
		Generated string  `json:"generated"`
		Author string  `json:"author"`
	} `json:"metadata"`
	Data []struct {
	  Category string `json:"category"`
	  Title string `json:"title"`
	  Commands []struct {
		Name string `json:"name"`
		Description string `json:"description"`
		Usage *string `json:"usage"`
		Examples *[]string `json:"examples"`
		Subtypes *[]struct {
			Subtype string `json:"subtype"`
			Examples []string `json:"examples"`
		} `json:"subtypes"`
	  } `json:"commands"`
	} `json:"data"`
 }

func GetThemeInfo(chosedId int) {
	switch (chosedId) {
	case 0:
	  createdStructureForReading()
	}
}

func createdStructureForReading() {
	 var commandsStructure PermissionStructure

	 resp, err := http.Get("https://filin.tech/static/permission.json")
	 if err != nil {
		fmt.Println(err)
	 }
     defer resp.Body.Close()

	 if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка получения статических данных от сервера")
	 }
	 
	 err = json.NewDecoder(resp.Body).Decode(&commandsStructure)

	 if err != nil {
		fmt.Println(err)
	 } else {
		fmt.Println()
		fmt.Println(commandsStructure.Metadata.Title)
		for _, dataCommand := range commandsStructure.Data {
		   fmt.Println()
           fmt.Println(dataCommand.Title)
		   fmt.Println()
		   for _, command := range dataCommand.Commands {
			bulletListItems := []pterm.BulletListItem{
				{
					Level:       0,
					Text:        command.Name + " - " + command.Description,  
					Bullet:      ">",
					TextStyle:   pterm.NewStyle(pterm.BgGreen),
					BulletStyle: pterm.NewStyle(pterm.FgYellow),
				},
			}
			if command.Usage != nil {
				appendedItem := pterm.BulletListItem{
					Level:       2,
					Text:        *command.Usage,     
					Bullet: "$",
					TextStyle:   pterm.NewStyle(pterm.FgGreen),
					BulletStyle: pterm.NewStyle(pterm.FgGreen),
				}
				bulletListItems = append(bulletListItems, appendedItem)
			}
			if command.Examples != nil {
				appendedItem := []pterm.BulletListItem{
				  {
					Level:       2,
					Text:        "Примеры использования",
					Bullet:      "🛠️",       
					TextStyle:   pterm.NewStyle(pterm.FgGray),
				  },
				}
				for _, example := range *command.Examples {
					appendedItemExample := pterm.BulletListItem{
						Level:       2,
						Text:        example,              
						TextStyle:   pterm.NewStyle(pterm.FgLightYellow),
						BulletStyle: pterm.NewStyle(pterm.FgRed),
					}
					appendedItem = append(appendedItem, appendedItemExample)
				}
				bulletListItems = append(bulletListItems, appendedItem...)
			}
			pterm.DefaultBulletList.WithItems(bulletListItems).Render()
			if command.Subtypes != nil {
				for _, commandType := range *command.Subtypes {
					fmt.Println(commandType.Subtype)
					for _, example := range commandType.Examples {
						fmt.Println(example)
					}
					
				}
			}
		   }
		} 
	 }
}