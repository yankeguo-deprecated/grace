package gracenotify

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/guoyk93/grace"
	"net/http"
	"os"
	"strings"

	"github.com/guoyk93/grace/gracetrack"
)

const (
	FormatQYWX = "qywx"
)

func Notify(title string, ctx *context.Context, err *error) {
	u := strings.TrimSpace(os.Getenv("NOTIFY_URL"))
	if u == "" {
		return
	}
	f := strings.ToLower(strings.TrimSpace(os.Getenv("NOTIFY_FORMAT")))
	var items []string
	if track := gracetrack.Extract(*ctx); track != nil {
		items = track.DumpPlain()
	}
	if *err != nil {
		items = append(items, "ERROR: "+(*err).Error())
	}
	if len(items) == 0 {
		items = append(items, "DONE")
	}
	text := title + "\n" + strings.Join(items, "\n")
	var buf []byte
	switch f {
	case FormatQYWX:
		buf, _ = json.Marshal(grace.M{"msgtype": "text", "text": grace.M{"content": text}})
	default:
		buf, _ = json.Marshal(map[string]interface{}{"text": text})
	}
	if len(buf) == 0 {
		return
	}
	res, _ := http.Post(u, "application/json", bytes.NewReader(buf))
	if res != nil {
		_ = res.Body.Close()
	}
}
