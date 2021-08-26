package helper

import (
	"os/exec"

	"github.com/unchainio/pkg/xlogger"
)

func SetMailServerIP(log *xlogger.Logger) {
	cmd := exec.Command("./setMailServerIP.sh")
	cmd.Wait()
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorln("Unable to set mail server ip")
		log.Errorln(string(output))
		log.Fatal(err)
	}
	log.Print("Mail server ip set successfully")
}
