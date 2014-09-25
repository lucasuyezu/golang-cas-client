Go CAS Client
=============

How to request a Service Ticket
-------------------------------

    import (
      "fmt"
      "github.com/lucasuyezu/cas"
    )

    func main() {
      cas := cas.New("https://server", "user", "pass")
      ticket, _ := cas.CreateServiceTicket("service")

      fmt.Println("ST is", ticket)
    }

TODO
----

* Improve error handling
* Validate a Service Ticket
* Invalidate a Service Ticket
* Reuse a TGT (Ticket Granting Ticket) to generate more than one Service Ticket
* Invalidate a TGT (sign out from all services)
