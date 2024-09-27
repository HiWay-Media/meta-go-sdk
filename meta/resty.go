package meta


func (o *meta) HealthCheck() error {
	resp, err := o.restyPost("/", nil)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("healthcheck error")
	}
	return nil
}
