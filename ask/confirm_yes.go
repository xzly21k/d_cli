package ask

import "github.com/AlecAivazis/survey/v2"

const (
	Yes = "YES"
	No  = "NO"
)

// ConfirmYes 询问YES/NO，yes返回true
func ConfirmYes(tips string) (bool, error) {
	result := ""
	prompt := &survey.Select{
		Message: tips,
		Options: []string{
			Yes,
			No,
		},
	}
	err := survey.AskOne(prompt, &result)
	if err != nil {
		return false, err
	}
	return result == Yes, nil
}
