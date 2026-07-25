package config

func (c Config) Validate() error {
	if err := c.App.Validate(); err != nil {
		return err
	}

	return nil
}
