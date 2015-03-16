package main

import "testing"

func ExmplProduct(t *testing.T) Product {
	product := Product{Name: "Test", Number: "1", Price: "1", Amount: "1"}
	return product
}
func ExmplProduct1(t *testing.T) Product {
	product := Product{Name: "Test1", Number: "2", Price: "2", Amount: "2"}
	return product
}
func ExmplProducts(t *testing.T) Products {
	var prds Products
	prds = append(prds, ExmplProduct(t))
	prds = append(prds, ExmplProduct1(t))
	return prds
}

func (prds Products) TGetNewItem(t *testing.T) {
	p := prds.GetNewItem()
	assertEqual(t, "", p.(Product).Name)
}

func (prds Products) TGetItem(t *testing.T) {
	p := prds.GetItem(0)
	assertEqual(t, prds[0].Name, p.(Product).Name)
}

func (prds Products) TGetName(t *testing.T) {
	assertEqual(t, "product", prds.GetName())
}

func (p Product) TShow(t *testing.T) {
	assertEqual(t, "\r\nName: Test\r\nNumber: 1\r\nPrice: 1\r\nAmount: 1", p.Show())
}

func (prds Products) TFindByName(t *testing.T) {
	assertEqual(t, 0, prds.FindByName("Test"))
}

func (prds Products) TAppend(t *testing.T) {
	i := prds.GetNewItem()
	prds.Append(i)
	assertEqual(t, prds[2].Name, "")
}

func (prds Products) TEdit(t *testing.T) {
	i := prds.GetItem(1)
	prds.Edit(0, i)
	assertEqual(t, prds[0].Name, prds[1].Name)
}

func (prds Products) TRemove(t *testing.T) {
	prds.Remove(1)
	i := prds.GetItem(0)
	prds.Append(i)
	assertEqual(t, prds[0].Name, prds[1].Name)
}
func TestProducts(t *testing.T) {
	prds := ExmplProducts(t)
	prds.TGetItem(t)
	prds.TGetNewItem(t)
	prds.TGetName(t)
	prds.TAppend(t)
	prds.TEdit(t)
	prds.TRemove(t)
	p := ExmplProduct(t)
	p.TShow(t)
}
