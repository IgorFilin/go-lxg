package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pterm/pterm"
)

func GetThemeInfo(chosedId int) {
	switch (chosedId) {
	case 0:
	  createdStructureForReading()
	}
}

func createdStructureForReading() {
	 var commandsStructure struct {
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
	//  var resultOutput string
	 jsonData, err := os.ReadFile("/data/permission.json")
	 if err != nil {
		fmt.Println(err)
	 } else {
		json.Unmarshal(jsonData, &commandsStructure)
		fmt.Println()
		fmt.Println(commandsStructure.Metadata.Title)
		for _, dataCommand := range commandsStructure.Data {
		   fmt.Println()
           fmt.Println(dataCommand.Title)
		   fmt.Println()
		   for _, command := range dataCommand.Commands {
			bulletListItems := []pterm.BulletListItem{
				{
					Level:       0,                            // Level 0 (top level)
					Text:        command.Name + " - " + command.Description,  
					Bullet:      ">",                       // Text to display
					TextStyle:   pterm.NewStyle(pterm.BgGreen), // Text color
					BulletStyle: pterm.NewStyle(pterm.FgYellow),
				},
				// {
				// 	Level:       1,                                  // Level 1 (sub-item)
			    //     Text:        ,                            // Text to display
			    //     TextStyle:   pterm.NewStyle(pterm.FgGreen),      // Text color
			    //     Bullet:      "-",                                // Custom bullet symbol
			    //     BulletStyle: pterm.NewStyle(pterm.FgLightWhite), // Bullet color
				// },
			}
			// fmt.Print(" - " + command.Name)
			// fmt.Print(" - " +command.Description)
			if command.Usage != nil {
				// fmt.Print(" - " + *command.Usage)
				appendedItem := pterm.BulletListItem{
					Level:       2,                            // Level 0 (top level)
					Text:        *command.Usage,     
					Bullet: "$",                  // Text to display
					TextStyle:   pterm.NewStyle(pterm.FgGreen), // Text color
					BulletStyle: pterm.NewStyle(pterm.FgGreen),  // Bullet color
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