package enum

import "golang.org/x/exp/slices"

const (
	TmsCabangApp      string = "tms_cabang"
	TmsPusatApp       string = "tms_pusat"
	TmsKurirApp       string = "tms_kurir"
	CustomerPortalApp string = "customer_portal"
)

func IsValidAppDomainType(s string) bool {
	stringArr := []string{TmsCabangApp, TmsPusatApp, TmsKurirApp}
	return slices.Contains(stringArr, s)
}
