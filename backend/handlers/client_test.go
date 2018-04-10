package handlers

import (
	"testing"

	"github.com/NaturalFractals/CapstoneProject/backend/model"
)

func (mdb *mockDB) GetClients(c *model.ClientSummary) error {
	return nil
}

// Test the get client endpoint
func TestGetClient(t *testing.T) {
}

func (mdb *mockDB) AddNewClient(c *model.Client) error {
	return nil
}

// Test the create client endpoint
func TestCreateClient(t *testing.T) {
}

func (mdb *mockDB) DeleteClient(c *model.Client) error {
	return nil
}

// Test the delete client endpoint
func TestDeleteClient(t *testing.T) {
}
