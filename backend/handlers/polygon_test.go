package handlers

import (
	"testing"

	"github.com/NaturalFractals/CapstoneProject/backend/model"
)

// Interface implementation for mockDB testing
func (mdb *mockDB) GetPolygons(p *model.PolygonSummary) error {
	return nil
}

// Test get Polygon endpoing
func TestGetPolygons(t *testing.T) {

}

// Interface implementation for mockDB testing
func (mdb *mockDB) SavePolygon(p *model.Polygon, c *model.Client) error {
	return nil
}

// Test saving polygon endpoint
func TestSavePolygon(t *testing.T) {

}

// Interface implementation for mockDB testing
func (mdb *mockDB) DeletePolygon(p *model.Polygon, c *model.Client) error {
	return nil
}

// Test the delete polygon endpoint
func TestDeletePolygon(t *testing.T) {

}
