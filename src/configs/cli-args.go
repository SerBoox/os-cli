package configs

import (
	"errors"
)

// CliArgs contain cli and other common params
type CliArgs struct {
	AuthHost string

	Login    string
	Password string

	InstName  string
	ImageRef  string
	FlavorRef int64
}

// GetArgs method set default params
func (cliArgs *CliArgs) GetArgs() *CliArgs {
	return cliArgs
}

// ValidateAuthHost method validate cli param
func (cliArgs *CliArgs) ValidateAuthHost() error {
	if cliArgs.AuthHost == "" {
		return errors.New("AuthHost not be empty")
	}

	return nil
}

// ValidateLogin method validate cli param
func (cliArgs *CliArgs) ValidateLogin() error {
	if cliArgs.Login == "" {
		return errors.New("login not be empty")
	}

	return nil
}

// ValidatePassword method validate cli param
func (cliArgs *CliArgs) ValidatePassword() error {
	if cliArgs.Password == "" {
		return errors.New("password not be empty")
	}

	return nil
}

// ValidateInstName method validate cli param
func (cliArgs *CliArgs) ValidateInstName() error {
	if cliArgs.InstName == "" {
		return errors.New("instName not be empty")
	}

	return nil
}

// ValidateImageRef method validate cli param
func (cliArgs *CliArgs) ValidateImageRef() error {
	if cliArgs.ImageRef == "" {
		return errors.New("imageRef not be empty")
	}

	return nil
}

// ValidateFlavorRef method validate cli param
func (cliArgs *CliArgs) ValidateFlavorRef() error {
	if cliArgs.FlavorRef == int64(0) {
		return errors.New("flavorRef not be empty")
	}

	return nil
}
