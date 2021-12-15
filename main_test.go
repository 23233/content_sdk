package content_sdk

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	sdk := New()

	appId := os.Getenv("app")

	t.Log(sdk.TextSecCheck("日天"))
	t.Log(sdk.ImageSecCheck("./test.jpg"))
	t.Log(sdk.GetAccessToken(appId))
	t.Log(sdk.RefreshAccessToken(appId))
	t.Log(sdk.GetAccessToken(appId))
}
