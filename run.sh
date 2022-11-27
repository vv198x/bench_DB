 #/bin/bash
 go test -bench=. -benchmem 2>/dev/null | column -t