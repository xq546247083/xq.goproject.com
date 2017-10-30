package spider

import (
	"errors"
	"testing"
)

func TestSpider(t *testing.T) {
	Start()
	t.Error(errors.New("1"))
}
