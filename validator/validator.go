package validator

import (
	"errors"
	"log"
	"os/exec"
	"strings"
)

// ErrUnsupportedScriptType is given when a script goes through validation and
// it's not possible to validate
var ErrUnsupportedScriptType = errors.New("unsupported script type")

// Validators represents supported validators
var Validators = []string{"sh", "py", "rb", "lua", "pl", "php"}

// ValidateScript will detect the given script type and perform validation on it
func ValidateScript(script string) bool {
	switch strings.Split(script, ".")[1] {
	case "sh":
		_, err := exec.Command("bash", "-n", script).CombinedOutput()
		if err != nil {
			log.Println(err)
			return false
		}
		return true
	case "py":
		_, err := exec.Command("python", "-m", "py_compile", script).CombinedOutput()
		if err != nil {
			log.Println(err)
			return false
		}
		return true
	case "rb":
		_, err := exec.Command("ruby", "-c", script).CombinedOutput()
		if err != nil {
			log.Println(err)
			return false
		}
		return true
	case "lua":
		_, err := exec.Command("lua", "-p", script).CombinedOutput()
		if err != nil {
			log.Println(err)
			return false
		}
		return true
	case "pl":
		_, err := exec.Command("perl", "-c", script).CombinedOutput()
		if err != nil {
			log.Println(err)
			return false
		}
		return true
	case "php":
		_, err := exec.Command("php", "-l", script).CombinedOutput()
		if err != nil {
			log.Println(err)
			return false
		}
		return true
	default:
		log.Println(ErrUnsupportedScriptType)
		return false
	}
}
