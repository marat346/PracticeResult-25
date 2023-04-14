package service

import (
	cities //-------------------------
)

type CityService struct {
	repo repository.CityList
}

func NewCityService(repo repository.CityList) *CityService {
	return &CityService{repo: repo}
}

// Create sends information to the repository and returns a response
func (s *CityService) Create(city cities.CityRequest) (string, error) {
	return s.repo.Create(city)
}

// Delete sends information to the repository and returns a response
func (s *CityService) Delete(id int) error {
	return s.repo.Delete(id)
}

// SetPopulation sends information to the repository and returns a response
func (s *CityService) SetPopulation(id, population int) error {
	return s.repo.SetPopulation(id, population)
}

// GetFromRegion sends information to the repository and returns a response
func (s *CityService) GetFromRegion(region string) ([]string, error) {
	return s.repo.GetFromRegion(region)
}

// GetFromDistrict sends information to the repository and returns a response
func (s *CityService) GetFromDistrict(distinct string) ([]string, error) {
	return s.repo.GetFromDistrict(distinct)
}

// GetFromPopulation sends information to the repository and returns a response
func (s *CityService) GetFromPopulation(start, end int) ([]string, error) {
	return s.repo.GetFromPopulation(start, end)
}

// GetFromFoundation sends information to the repository and returns a response
func (s *CityService) GetFromFoundation(start, end int) ([]string, error) {
	return s.repo.GetFromFoundation(start, end)
}

// GetFull sends information to the repository and returns a response
func (s *CityService) GetFull(id int) (*cities.City, error) {
	return s.repo.GetFull(id)
}