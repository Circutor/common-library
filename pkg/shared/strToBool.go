// Copyright (c) 2021 Circutor S.A. All rights reserved.

package shared

import "strconv"

// StrToBool parses string to boolean.
func StrToBool(strBool string) bool {
	strToBool, _ := strconv.ParseBool(strBool)

	return strToBool
}
