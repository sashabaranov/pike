package pike

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
)

const outro = `

🎉 %s 🎉

You project is placed in %s
Inside it you can find:
  • sql/ — database migration files ( use with https://github.com/golang-migrate/migrate )
  • proto/ — protobuf gRPC service description
  • %s/ and cli/ – Go source files


%s

1. Run %s
   It will compile .proto to .pb.go. Use it every time you modify .proto

2. %s

   CERT_DIR=%s/certs/dev
   mkdir -p $CERT_DIR
   certstrap --depot-path $CERT_DIR init --expires "30 years" --common-name "CA"
   certstrap --depot-path $CERT_DIR request-cert --domain localhost
   certstrap --depot-path $CERT_DIR sign localhost --CA CA

3. %s will start grpc server

`

func (p Project) PrintOutro() {
	fmt.Printf(
		outro,

		Bold("All done!"),
		Bold(Green(p.AbsolutePath())),
		p.Name,
		Underline(Bold("Further steps")),
		Cyan("bin/compile_proto.sh"),
		Cyan("Generate certificates using https://github.com/square/certstrap"),
		p.AbsolutePath(),
		Cyan("bin/run.sh"),
	)
}
