package gracenotify

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/guoyk93/grace/gracetrack"
)

func Notify(title string, ctx *context.Context, err *error) {
	u := strings.TrimSpace(os.Getenv("NOTIFY_URL"))
	if u == "" {
		return
	}
	var items []string
	if track := gracetrack.Extract(*ctx); track != nil {
		items = track.DumpPlain()
	}
	if *err != nil {
		items = append(items, "错误: "+(*err).Error())
	}
	if len(items) == 0 {
		items = append(items, "执行完成")
	}
	text := title + "\n" + strings.Join(items, "\n")

	buf, _ := json.Marshal(map[string]interface{}{"text": text})
	if len(buf) == 0 {
		return
	}
	res, _ := http.Post(u, "application/json", bytes.NewReader(buf))
	if res != nil {
		_ = res.Body.Close()
	}
}
