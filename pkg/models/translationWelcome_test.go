package models_test

import (
	"testing"
	"gopher-translator/pkg/mock"
)

func TestTranslationWelcome(t *testing.T) {
	testString := "test"
	welcome := "Welcome"
	routes := map[string]string{
		testString: testString,
	}
	mockInstance := mock.CreateNewMock()
	welcomeTranslation := mockInstance.CreateNewTranslationWelcomeStructInstance(welcome, routes)

	if welcomeTranslation.Welcome != welcome {
		t.Errorf("Welcome message expected %v got %v", welcome, welcomeTranslation.Welcome)
	}

	if welcomeTranslation.Routes[testString] != routes[testString] {
		t.Errorf("Routes key %v value expected %v got %v", testString, welcomeTranslation.Routes[testString], routes[testString])
	}
}