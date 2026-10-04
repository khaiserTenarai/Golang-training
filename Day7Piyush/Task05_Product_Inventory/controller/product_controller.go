package controller

import (
	"fmt"
	"task05_product_inventory/repository"
	"task05_product_inventory/view"
)

type ProductController struct {
	repo *repository.ProductRepository
	view *view.ProductView
}

func NewProductController(repo *repository.ProductRepository, v *view.ProductView) *ProductController {
	return &ProductController{repo: repo, view: v}
}

func (c *ProductController) Run() {
	for {
		choice := c.view.ShowMenu()
		switch choice {
		case 1:
			c.AddProduct()
		case 2:
			c.ListProducts()
		case 3:
			c.UpdateProduct()
		case 4:
			c.DeleteProduct()
		case 5:
			c.IncreaseStock()
		case 6:
			c.DecreaseStock()
		case 7:
			c.LowStockProducts()
		case 8:
			c.SearchProduct()
		case 9:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func (c *ProductController) AddProduct() {
	p := c.view.GetProductInput()
	id, err := c.repo.Create(p)
	if err != nil {
		c.view.ShowError("adding product", err)
		return
	}
	c.view.ShowSuccess(fmt.Sprintf("Product added with ID: %d", id))
}

func (c *ProductController) ListProducts() {
	products, err := c.repo.GetAll()
	if err != nil {
		c.view.ShowError("listing products", err)
		return
	}
	c.view.ShowProducts(products)
}

func (c *ProductController) UpdateProduct() {
	id := c.view.GetID()
	existing, err := c.repo.GetByID(id)
	if err != nil {
		c.view.ShowError("fetching product", err)
		return
	}
	updated := c.view.GetUpdateInput(existing)
	if err := c.repo.Update(updated); err != nil {
		c.view.ShowError("updating product", err)
		return
	}
	c.view.ShowSuccess("Product updated.")
}

func (c *ProductController) DeleteProduct() {
	id := c.view.GetID()
	if err := c.repo.Delete(id); err != nil {
		c.view.ShowError("deleting product", err)
		return
	}
	c.view.ShowSuccess("Product deleted.")
}

func (c *ProductController) IncreaseStock() {
	id := c.view.GetID()
	qty := c.view.GetQuantity()
	if err := c.repo.IncreaseStock(id, qty); err != nil {
		c.view.ShowError("increasing stock", err)
		return
	}
	c.view.ShowSuccess(fmt.Sprintf("Stock increased by %d.", qty))
}

func (c *ProductController) DecreaseStock() {
	id := c.view.GetID()
	qty := c.view.GetQuantity()
	if err := c.repo.DecreaseStock(id, qty); err != nil {
		c.view.ShowError("decreasing stock", err)
		return
	}
	c.view.ShowSuccess(fmt.Sprintf("Stock decreased by %d.", qty))
}

func (c *ProductController) LowStockProducts() {
	threshold := c.view.GetThreshold()
	products, err := c.repo.GetLowStock(threshold)
	if err != nil {
		c.view.ShowError("fetching low-stock products", err)
		return
	}
	c.view.ShowProducts(products)
}

func (c *ProductController) SearchProduct() {
	name := c.view.GetSearchName()
	products, err := c.repo.SearchByName(name)
	if err != nil {
		c.view.ShowError("searching products", err)
		return
	}
	c.view.ShowProducts(products)
}
