package util

import (
	"fmt"

	"github.com/free5gc/openapi/models"
)

// SearchNFServiceUri returns NF Uri derived from NfProfile with corresponding service
func SearchNFServiceUri(nfProfile models.NfProfile, serviceName models.ServiceName,
	nfServiceStatus models.NfServiceStatus) (nfUri string) {
	if nfProfile.NfServices != nil {
		for _, service := range *nfProfile.NfServices {
			if service.ServiceName == serviceName && service.NfServiceStatus == nfServiceStatus {
				if nfProfile.Fqdn != "" {
					nfUri = nfProfile.Fqdn
					//nfUri = "http://127.0.0.1:9999"
				} else if service.Fqdn != "" {
					nfUri = service.Fqdn
					//nfUri = "http://127.0.0.1:9999"
				} else if service.ApiPrefix != "" {
					nfUri = service.ApiPrefix
					//nfUri = "http://127.0.0.1:9999"
				} else if service.IpEndPoints != nil {
					point := (*service.IpEndPoints)[0]
					if point.Ipv4Address != "" {
						nfUri = getSbiUri(service.Scheme, point.Ipv4Address, point.Port)
						//nfUri = "http://127.0.0.1:9999"
					} else if len(nfProfile.Ipv4Addresses) != 0 {
						nfUri = getSbiUri(service.Scheme, nfProfile.Ipv4Addresses[0], point.Port)
						//nfUri = "http://127.0.0.1:9999"
					}
				}
			}
			if nfUri != "" {
				break
			}
		}
	}

	return
}

func getSbiUri(scheme models.UriScheme, ipv4Address string, port int32) (uri string) {
	if port != 0 {
		uri = fmt.Sprintf("%s://%s:%d", scheme, ipv4Address, port)
	} else {
		switch scheme {
		case models.UriScheme_HTTP:
			uri = fmt.Sprintf("%s://%s:80", scheme, ipv4Address)
		case models.UriScheme_HTTPS:
			uri = fmt.Sprintf("%s://%s:443", scheme, ipv4Address)
		}
	}
	//uri = fmt.Sprintf("http://localhost:9999")
	return
}
