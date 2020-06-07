package pike

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
)

const outro = `

🎉 %s 🎉

Your project is placed in %s

%s

1. %s

   CERT_DIR=%s/certs/dev
   mkdir -p $CERT_DIR
   certstrap --depot-path $CERT_DIR init --expires "30 years" --common-name "CA"
   certstrap --depot-path $CERT_DIR request-cert --domain localhost
   certstrap --depot-path $CERT_DIR sign localhost --CA CA

2. %s will start grpc server

`

func (p Project) PrintOutro() {
	fmt.Printf(
		outro,

		Bold("All done!"),
		Bold(Green(p.AbsolutePath())),
		Underline(Bold("Further steps")),
		Cyan("Generate certificates using https://github.com/square/certstrap"),
		p.AbsolutePath(),
		Cyan("bin/run.sh"),
	)
}
