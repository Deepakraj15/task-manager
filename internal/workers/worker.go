package workers

import "fmt"

func CmdWorker(option string) string {
	switch option {
	case "1":
		return "Option 1 has been selected\n"
	case "2":
		return "Option 2 has been selected\n"
	default:
		return fmt.Sprintf("Incorrect Option: %s. Please select correct option\n", option)
	}
}
