package snr

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestHelp(t *testing.T) {
	//c := NewProxyClient("snr-proxy-test.qoorp.com", "qoorp", "jtgSiW9T57ggjLBWsEf1XPW4ZxNGLgQHRp7G2QrdRxZhXRskPzed4yZhQTH5voyAfCQ8RHGcPvadEHP2g7W8P1IEAW94M6tB1d0b")
	//c.Do(nil)
	x := NewInformationshuvud()
	b, err := xml.MarshalIndent(x, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
