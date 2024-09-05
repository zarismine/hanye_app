package pb

import "strings"

func Address2string(address *Address) string {
	strList := []string{address.ProvinceName, address.CityName, address.DistrictName, address.Detail}
	resp := strings.Join(strList, ",")
	return resp
}
