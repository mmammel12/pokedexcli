package main

import (
	"fmt"
)

func commandMap(c *config, _ []string) error {
	locationsResp, err := c.pokeapiClient.ListLocations(c.Next)
	if err != nil {
		return err
	}

	c.Next = locationsResp.Next
	c.Previous = locationsResp.Previous

	fmt.Println()
	for _, location := range locationsResp.Results {
		fmt.Println("\t- " + location.Name)
	}
	fmt.Println()

	return nil
}

func commandMapb(c *config, _ []string) error {
	locationsResp, err := c.pokeapiClient.ListLocations(c.Previous)
	if err != nil {
		return err
	}

	c.Next = locationsResp.Next
	c.Previous = locationsResp.Previous

	fmt.Println()
	for _, location := range locationsResp.Results {
		fmt.Println("\t- " + location.Name)
	}
	fmt.Println()

	return nil
}
