package usecases

import "fmt"

type ServiceIntegrationUseCases struct {
}

func (uc *ServiceIntegrationUseCases) ValidateServiceKey(serviceKey string) error {
	if serviceKey != "c5d7c94f" {
		return fmt.Errorf("invalid service key: %v", serviceKey)
	}
	return nil
}
