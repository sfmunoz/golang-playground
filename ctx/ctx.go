//
// vim: set foldmethod=marker:
//
// URL:  https://github.com/sfmunoz/golang-playground
// Date: Wed Jan 31 09:30:29 AM UTC 2024
//

// {{{ package

package ctx

// }}}
// {{{ imports

import (
	"context"
	"fmt"
	"log"
	"time"
)

// }}}
// ---- funcs (private) ----
// {{{ func theJob()

func theJob(ctx context.Context, t time.Duration, cancel func()) {
	defer cancel()
	idx := fmt.Sprintf("%s | %s", ctx, t)
	select {
	case <-time.After(t):
		log.Printf("%s -> job done!!", idx)
		return
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("%s -> err = %v", idx, err)
		return
	}
}

// }}}
// ---- funcs (public) ----
// {{{ func Main()

func Main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.LUTC)
	items := []struct{ ok, ko time.Duration }{{50, 100}, {150, 100}}
	for _, v := range items {
		ok, ko := v.ok*time.Millisecond, v.ko*time.Millisecond
		log.Printf("job=%s; timeout=%s", ok, ko)
		ctx, cancel := context.WithTimeout(context.Background(), ko)
		theJob(ctx, ok, cancel)
	}
}

// }}}
