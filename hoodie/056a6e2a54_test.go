// Test generated by RoostGPT for test golang-crud-api using AI Type Open AI and AI Model gpt-4-1106-preview

func (r *repository) Update(hoodie Hoodie) (Hoodie, error) {
	err := r.db.Save(&hoodie).Error
	return hoodie, err
}
