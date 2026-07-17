/*
Copyright © 2026 Milos Slavic <m.slavic@outlook.com>
*/
package cmd

import (
	"fmt"

	"github.com/MilosSlavic/go-playground/pkg/linkedlist"
	"github.com/spf13/cobra"
)

type product struct {
	name string
}

// linkedlistCmd represents the linkedlist command
var linkedlistCmd = &cobra.Command{
	Use:   "linkedlist",
	Short: "Execute linked list example",
	Long:  `Execute linked list example`,
	Run: func(cmd *cobra.Command, args []string) {
		linkedListExample()
	},
}

func init() {
	rootCmd.AddCommand(linkedlistCmd)
}

func linkedListExample() {
	ll := linkedlist.NewLinkedList()
	ll.Add(&product{name: "Milk"})
	item := &product{name: "Bread"}
	ll.Add(item)

	printLinkedList(ll)

	fmt.Println("Deleting...", item.name)
	ll.Delete(item)

	printLinkedList(ll)
}

func printLinkedList(ll *linkedlist.LinkedList) {
	for value := range ll.All() {
		item := value.(*product)
		fmt.Println(item.name)
	}
}
