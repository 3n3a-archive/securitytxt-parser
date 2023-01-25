package SecurityTxtParser

import (
    "testing"
    _ "regexp"
)

func TestParsingRequired(t *testing.T) {
    txtFile := `
Contact: mailto:test@example.com
Expires: 2023-03-14T00:00:00.000Z
`
    msg, err := ParseTxt(txtFile)
    if err != nil {
        t.Fatalf(`Msg: %v, Err: %v`, msg, err)
    }
}
