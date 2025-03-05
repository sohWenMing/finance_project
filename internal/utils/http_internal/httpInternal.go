package httpInternal

import (
	"net/http"
	"time"
)

var DefaultClient = &http.Client{
	Timeout: 5 * time.Second,
}
