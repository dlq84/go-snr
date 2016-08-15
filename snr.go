package snr

import (
	"encoding/base64"
	"encoding/xml"
)

// Client struct
type Client struct {
	xmlProdukt *XMLProdukt
	userID     string
	certID     string
	url        string
}

// NewProxyClient creates a new client without certificate, to use with proxy that uses client cert
func NewProxyClient(url string, userID string, certID string, basicAuth *BasicAuth) *Client {
	return &Client{
		xmlProdukt: NewXMLProdukt(url, false, basicAuth),
		userID:     userID,
		certID:     certID,
	}
}

// GetProdukt executes a request
func (c *Client) GetProdukt(xmlFraga *XMLFraga) (*GetProduktResponse, error) {
	data, err := xml.Marshal(xmlFraga)
	if err != nil {
		return nil, err
	}
	getProdukt := &GetProdukt{
		UserID:   c.userID,
		CertID:   c.certID,
		XMLFraga: base64.StdEncoding.EncodeToString(data),
	}
	return c.xmlProdukt.GetProdukt(getProdukt)
}

/*


	var (
		certFile = flag.String("cert", "steria.crt.pem", "A PEM eoncoded certificate file.")
		keyFile  = flag.String("key", "steria.key.pem", "A PEM encoded private key file.")
	)
		// NewTLSClient creates new client instance and uses client key and cert
		func NewTLSClient(userID string, certID string, keyPath string, certPath string) *Client {
			jar, _ := cookiejar.New(nil)
			cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
			if err != nil {
				log.Fatal(err)
			}
			caCertPool := x509.NewCertPool()
			tlsConfig := &tls.Config{
				Certificates: []tls.Certificate{cert},
				RootCAs:      caCertPool,
			}
			transport := &http.Transport{TLSClientConfig: tlsConfig}
			client := &http.Client{Transport: transport, Jar: jar}
			return nil
		}
*/
