package main

import "testing"

func ExmplClient(t *testing.T) Client {
	client := Client{Name: "Test", Surname: "Test", Login: "Test", Password: "Test", CreditCard: "Test", Street: "Test", City: "Test", State: "Test", Zip: "Test", OrdersAmount: 1}
	return client
}
func ExmplClient1(t *testing.T) Client {
	client := Client{Name: "Test1", Surname: "Test1", Login: "Test1", Password: "Test1", CreditCard: "Test1", Street: "Test1", City: "Test1", State: "Test1", Zip: "Test1", OrdersAmount: 2}
	return client
}
func ExmplClients(t *testing.T) Clients {
	var cls Clients
	cls.Cls = append(cls.Cls, ExmplClient(t))
	cls.Cls = append(cls.Cls, ExmplClient1(t))
	return cls
}

func (cls Clients) TGetNewItem(t *testing.T) {
	c := cls.GetNewItem()
	assertEqual(t, "", c.(Client).Name)
}

func (cls Clients) TGetItem(t *testing.T) {
	c := cls.GetItem(0)
	assertEqual(t, cls.Cls[0].Name, c.(Client).Name)
}

func (cls Clients) TGetName(t *testing.T) {
	assertEqual(t, "client", cls.GetName())
}

func (c Client) TShow(t *testing.T) {
	assertEqual(t, "\r\nName: Test\r\nSurname: Test\r\nLogin: Test\r\nPassword: Test\r\nCredit card: Test\r\nStreet: Test\r\nCity: Test\r\nState: Test\r\nZip: Test", c.Show())
}

func (cls Clients) TFindByName(t *testing.T) {
	assertEqual(t, 0, cls.FindByName("Test"))
}

func (cls Clients) TAppend(t *testing.T) {
	i := cls.GetNewItem()
	cls.Append(i)
	assertEqual(t, cls.Cls[2].Name, "")
}

func (cls Clients) TEdit(t *testing.T) {
	i := cls.GetItem(1)
	cls.Edit(0, i)
	assertEqual(t, cls.Cls[0].Name, cls.Cls[1].Name)
}

func (cls Clients) TRemove(t *testing.T) {
	cls.Remove(1)
	i := cls.GetItem(0)
	cls.Append(i)
	assertEqual(t, cls.Cls[0].Name, cls.Cls[1].Name)
}
func TestClients(t *testing.T) {
	cls := ExmplClients(t)
	cls.TGetItem(t)
	cls.TGetNewItem(t)
	cls.TGetName(t)
	cls.TAppend(t)
	cls.TEdit(t)
	cls.TRemove(t)
	c := ExmplClient(t)
	c.TShow(t)
}
