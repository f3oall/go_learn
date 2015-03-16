package main

import "testing"

func ExmplOrder(t *testing.T) Order {
	order := Order{Customer: "Test", Items: Products{Product{Name: "Test"}}, Bill: "1", ItemsAmount: 1}
	return order
}
func ExmplOrder1(t *testing.T) Order {
	order := Order{Customer: "Test1", Items: Products{Product{Name: "Test1"}}, Bill: "2", ItemsAmount: 2}
	return order
}
func ExmplOrders(t *testing.T) Orders {
	var ords Orders
	ords = append(ords, ExmplOrder(t))
	ords = append(ords, ExmplOrder1(t))
	return ords
}

func (ords Orders) TGetNewItem(t *testing.T) {
	o := ords.GetNewItem()
	assertEqual(t, "", o.(Order).Customer)
}

func (ords Orders) TGetItem(t *testing.T) {
	o := ords.GetItem(0)
	assertEqual(t, ords[0].Customer, o.(Order).Customer)
}

func (ords Orders) TGetName(t *testing.T) {
	assertEqual(t, "order", ords.GetName())
}

func (o Order) TShow(t *testing.T) {
	assertEqual(t, "\r\nCustomer: Test\r\nItems: \r\nName: Test\r\nNumber: \r\nPrice: \r\nAmount: \r\n____________________________________________\r\nBill: 1", o.Show())
}

func (ords Orders) TFindByName(t *testing.T) {
	assertEqual(t, 0, ords.FindByName("Test"))
}

func (ords Orders) TAppend(t *testing.T) {
	i := ords.GetNewItem()
	ords.Append(i)
	assertEqual(t, ords[2].Customer, "")
}

func (ords Orders) TEdit(t *testing.T) {
	i := ords.GetItem(1)
	ords.Edit(0, i)
	assertEqual(t, ords[0].Customer, ords[1].Customer)
}

func (ords Orders) TRemove(t *testing.T) {
	ords.Remove(1)
	i := ords.GetItem(0)
	ords.Append(i)
	assertEqual(t, ords[0].Customer, ords[1].Customer)
}
func TestOrders(t *testing.T) {
	ords := ExmplOrders(t)
	ords.TGetItem(t)
	ords.TGetNewItem(t)
	ords.TGetName(t)
	ords.TAppend(t)
	ords.TEdit(t)
	ords.TRemove(t)
	o := ExmplOrder(t)
	o.TShow(t)
}
