package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you?\n")

	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"name": "Utkarsh",
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name :")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	templates := map[string]string{
		"welcome": "Welcome, {{.name}}! How are you?\n",
		"notification": "{{.name}} you have a new notification {{.notification}}\n",
		"error": "Error occured : {{.errorMessage}}\n",
	}

	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		fmt.Println("Menu:")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)


	}
	
}