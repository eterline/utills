package utilla

import (
	"log"
	"os/exec"
	"os/user"
	"unicode"
)

// Checks root permissions of user
// If isn't, throw error.
func MustPriveges() {
	current, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	if current.Uid != "0" {
		log.Fatal("Requires superuser privileges.")
	}
}

// Do commands with arguments.
// 'Command' 'args' -> 'command output'.
func ExecCmd(cmd string, arg ...string) string {
	out, err := exec.Command(cmd, arg...).Output()
	if err != nil {
		return err.Error()
	}
	return string(out)
}

// String contains only int.
func StrIsInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// Clean string from letters: "%@^&*123ewwedw" -> "123".
// If None -> "0".
// func StringOnlyInt(str string) string {
// 	var buffer bytes.Buffer
// 	nums := []string{
// 		"0",
// 		"1",
// 		"2",
// 		"3",
// 		"4",
// 		"5",
// 		"6",
// 		"7",
// 		"8",
// 		"9",
// 	}
// 	for _, i := range str {
// 		if slices.Contains(nums, string(i)) {
// 			buffer.WriteString(string(i))
// 		}
// 	}
// 	res := buffer.String()
// 	if res != "" {
// 		return res
// 	} else {
// 		return "0"
// 	}
// }
