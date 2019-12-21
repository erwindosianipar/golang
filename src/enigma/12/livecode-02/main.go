package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type contact struct {
	id      int
	name    string
	gender  string
	groupID int
}

type group struct {
	id   int
	name string
}

var contacts []contact
var groups []group

// global
var menu int

// contact
var contactID, contactGroupID int
var contactName, gender string

// group
var groupID int
var groupName string

func main() {
	menu = menuUtama()

	switch menu {
	case 1:
		addContact()
		main()
	case 2:
		addGroup()
		main()
	case 3:
		viewListContact()
		main()
	case 4:
		viewListGroup()
		main()
	case 6:
		fmt.Println("Success logout from the application")
		os.Exit(0)
	}

}

func menuUtama() int {
	fmt.Println(strings.Repeat("-", 35))
	fmt.Println("Choose the menu application")
	fmt.Println(strings.Repeat("-", 35))
	fmt.Println("1. Add Contact")
	fmt.Println("2. Add Group")
	fmt.Println("3. View List Contact")
	fmt.Println("4. View List Group")
	fmt.Println("5. Show Contact with Group Name")
	fmt.Println("6. Exit the application")
	fmt.Println(strings.Repeat("-", 35))

	for {
		fmt.Printf("Enter application menu: ")
		menu, err := strconv.Atoi(input())
		if err == nil {
			if menu > 0 && menu < 7 {
				return menu
			}
			fmt.Println("Error: Menu you entered must 1-6")
		} else {
			fmt.Println("Error: Menu must only contains a number")
		}
	}
}

func input() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func addContact() {

	if len(contacts) > 0 {

		fmt.Println(strings.Repeat("-", 35))
		fmt.Println("Add a contact")
		fmt.Println(strings.Repeat("-", 35))

		for {
			fmt.Printf("Enter a contact ID: ")
			contactID, err := strconv.Atoi(input())
			if err == nil {
				if contactID > 0 {
					break
				}
				fmt.Println("Error: Contact ID cannot be empty and zero")
			} else {
				fmt.Println("Error: Menu must only contains a number")
			}
		}

		for {
			fmt.Printf("Enter a contact name: ")
			contactName = input()
			if len(contactName) > 2 {
				break
			}
			fmt.Println("Error: Contact name must be 3 character")
		}

		for {
			fmt.Printf("Enter a contact gender: ")
			gender = input()
			if len(gender) > 3 {
				break
			}
			fmt.Println("Error: Gender of contact  must be 3 character")
		}

	} else {
		fmt.Println("Error: To add contact must be have a group")
	}

}

func addGroup() {

	fmt.Println(strings.Repeat("-", 35))
	fmt.Println("Add a group")
	fmt.Println(strings.Repeat("-", 35))

	for {
		fmt.Printf("Enter a group ID: ")
		groupID, err := strconv.Atoi(input())
		if err == nil {
			if groupID > 0 {
				break
			}
			fmt.Println("Error: Group ID cannot be empty and zero")
		} else {
			fmt.Println("Error: Menu must only contains a number")
		}
	}

	for {
		fmt.Printf("Enter a group name: ")
		GroupName := input()
		if len(GroupName) > 2 {
			break
		}
		fmt.Println("Error: A group name minimal 3 character")
	}
	saveGroup(groupID, groupName)
	return
}

func saveGroup(groupID int, groupName string) {
	groups = append(groups, group{
		id:   groupID,
		name: groupName,
	})
	fmt.Println(groups)
}

func saveContact(contactID int, contactName string, gender string, groupID int) {
	contacts = append(contacts, contact{
		id:      contactID,
		name:    contactName,
		gender:  gender,
		groupID: groupID,
	})
	fmt.Println(contacts)
}

func viewListContact() {
	fmt.Println(contacts)
}

func viewListGroup() {
	fmt.Println(groups)
}
