// Test generated by RoostGPT for test golang-crud-api using AI Type Open AI and AI Model gpt-4-1106-preview

package allproducts_test

import (
	"testing"
	"reflect"
	"golang_e-commerce_API/allproducts"
)

// MockRepository is a mock of the Repository interface
type MockRepository struct {
	BooksByCategory map[string][]allproducts.AllProduct
}

// NewMockRepository creates a new instance of MockRepository with test data
func NewMockRepository() *MockRepository {
	return &MockRepository{
		BooksByCategory: map[string][]allproducts.AllProduct{
			"Fiction": {
				{ID: 1, Name_product: "The Great Gatsby", Category: "Fiction"},
				{ID: 2, Name_product: "1984", Category: "Fiction"},
			},
			"Non-Fiction": {
				{ID: 3, Name_product: "Sapiens", Category: "Non-Fiction"},
			},
		},
	}
}

// FindByCategory in MockRepository returns the products for the given category
func (m *MockRepository) FindByCategory(category string) ([]allproducts.AllProduct, error) {
	if books, ok := m.BooksByCategory[category]; ok {
		return books, nil
	}
	return nil, allproducts.ErrCategoryNotFound
}

// TestServiceFindByCategory tests the FindByCategory method with both positive and negative cases
func TestServiceFindByCategory(t *testing.T) {
	mockRepo := NewMockRepository()
	service := allproducts.NewService(mockRepo)
	
	tests := []struct {
		category string
		expected []allproducts.AllProduct
		err      error
	}{
		// Positive test case
		{"Fiction", mockRepo.BooksByCategory["Fiction"], nil},
		// Negative test case
		{"Poetry", nil, allproducts.ErrCategoryNotFound},
	}
	
	for _, test := range tests {
		products, err := service.FindByCategory(test.category)
		
		if !reflect.DeepEqual(products, test.expected) || (err != test.err) {
			t.Errorf("FindByCategory(%s): expected %v, got %v; expected error %v, got %v", 
					 test.category, test.expected, products, test.err, err)
		}
	}
}
