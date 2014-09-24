Go CAS Client
=============

How to request a Service Ticket
-------------------------------

    import (
      "fmt"
      "github.com/lucasuyezu/cas"
    )

    func main() {
      cas_server_url := "https://your-cas-server.com"
      username := "my-user"
      password := "my-password"
      service := "service.company.com"
      fmt.Println("ST is", cas.RequestST(cas_server_url, username, password, service))
    }

TODO
----

* Validate a Service Ticket
* Invalidate a Service Ticket
* Reuse a TGT (Ticket Granting Ticket) to generate more than one Service Ticket
* Invalidate a TGT (sign out from all services)
