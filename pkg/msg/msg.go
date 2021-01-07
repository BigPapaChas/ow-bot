package msg

import (
	"encoding/json"
	"fmt"
)

// Arguments accepted by the function to generate the message json payload.
// 0-3 roles can be provided
// Example:
// 	{
// 		"usrn": "Chris",
// 		"roles": [
// 			{
// 				"role": "Damage",
// 				"sr": "3000"
// 			},
// 			{
// 				"role": "Support",
// 				"sr": "3000"
// 			},
// 			{
// 				"role": "Tank",
// 				"sr": "3000"
// 			}
// 		]
// 	}
type MsgArgs struct {
	Usrn  string
	Roles [3]Role
}

//
type Role struct {
	Role string
	SR   string
}

// A message can have one content string (the actual discord message)
// and embeded objects that can have images, links, etc.
// Discord can have up to 10 embeded objects. Each object can have a different set of keys.
// May want to standardize what we want embeded or create structs for each type of key that
// could be added (Image, Link, etc.)
type Message struct {
	Content string                   `json:"content"`
	Embeds  []map[string]interface{} `json:"embeds"`
}

// Testing:
// func main() {
// 	damage := Role{
// 		"Damage",
// 		"3000",
// 	}
// 	support := Role{
// 		"Support",
// 		"3000",
// 	}
// 	roles := [3]Role{
// 		damage, support,
// 	}
// 	msg, err := msgBuild(MsgArgs{"Chris", roles})
// 	if err != nil {
// 		return
// 	}
// 	fmt.Println(string(msg))
// }

// Builds the mesage to be sent to discord
func msgBuild(args MsgArgs) ([]byte, error) {
	content := cntBuild(args)
	// Possibly want to build the urls for images or use the username to get the bytes for an image here
	// Example dmgUrl := "https://MyURL/damage_" + args.Usrn
	embs := make([]map[string]interface{}, 0, 10)
	msg := Message{
		content,
		embs,
	}
	return json.Marshal(msg)
}

func cntBuild(args MsgArgs) string {
	content := "OW report for: " + args.Usrn
	for i := 0; i < len(args.Roles); i++ {
		if args.Roles[i].Role != "" {
			content += fmt.Sprintf("\n%s: %s", args.Roles[i].Role, args.Roles[i].SR)
		}
	}
	return content
}
