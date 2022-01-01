package zap

type Automation struct {
	c *Client
}

// Runs an automation plan
//
// This component is optional and therefore the API will only work if it is installed
func (a Automation) RunPlan(file string) (map[string]interface{}, error) {
	m := map[string]string{
		"filePath": file,
	}
	return a.c.Request("automation/action/runPlan/", m)
}

// Returns Automation Plan Progress
//
// This component is optional and therefore the API will only work if it is installed
func (a Automation) PlanProgress(planId string) (map[string]interface{}, error) {
	m := map[string]string{
		"planId": planId,
	}
	return a.c.Zap.Request("automation/view/planProgress/", m)
}
