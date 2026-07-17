/*
Copyright © 2026 Milos Slavic <m.slavic@outlook.com>
*/
package cmd

import (
	"fmt"

	"github.com/MilosSlavic/go-playground/pkg/linkedlist/generic"
	"github.com/spf13/cobra"
)

type animal struct {
	name string
}

// genericLinkedlistCmd represents the genericLinkedlist command
var genericLinkedlistCmd = &cobra.Command{
	Use:   "genericLinkedlist",
	Short: "Execute generic linked list example",
	Long:  `Execute generic linked list example`,
	Run: func(cmd *cobra.Command, args []string) {
		genericLinkedListExample()
	},
}

func init() {
	rootCmd.AddCommand(genericLinkedlistCmd)
}

func genericLinkedListExample() {
	ll := generic.NewLinkedList[*animal]()
	ll.Add(&animal{name: "Cow"})
	item := &animal{name: "Chicken"}
	ll.Add(item)

	printAnimalList(ll)

	fmt.Println("Deleting...", item.name)
	ll.Delete(item)

	printAnimalList(ll)
}

func printAnimalList(ll *generic.LinkedList[*animal]) {
	for value := range ll.All() {
		fmt.Println(value.name)
	}
}
