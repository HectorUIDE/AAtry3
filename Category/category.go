package Category

import (
	"fmt"
)

type Category struct {
	Name string
}

var categories []Category

//---------Setter-----------------------------------------------------

// Makes sure categories don't repeat themselves
func (c *Category) SetCategory(name string) error {
	for _, category := range categories {
		if category.Name == name {
			return fmt.Errorf("Category already exists")
		}
	}
	c.Name = name
	return fmt.Errorf("Category is adequate!")
}

//---------Getter-----------------------------------------------------

func (c *Category) GetCategory() string {
	return c.Name
}

//---------Functions-------------

func (newCategory Category) AddCategory() error {
	if newCategory.Name == "" {
		return fmt.Errorf("Failed to add Category")
	}
	categories = append(categories, newCategory)
	return fmt.Errorf("Category added successfully!")
}

func ViewCategories() {
	fmt.Println("Categories:")
	for _, category := range categories {
		fmt.Println(category.Name)
	}
}
