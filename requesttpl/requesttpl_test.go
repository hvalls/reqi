package requesttpl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	tpl := New("slack/send", "Send slack message", "https://www.slack.com/xxx/xxx/xxx", "post", "{ \"text\": \"Test\" }")

	str, _ := tpl.String()

	expected := "name: slack/send\ndescription: Send slack message\nurl: https://www.slack.com/xxx/xxx/xxx\nmethod: post\nbody: '{ \"text\": \"Test\" }'\n"

	assert.Equal(t, expected, str, "The two contents should be the same")
}
