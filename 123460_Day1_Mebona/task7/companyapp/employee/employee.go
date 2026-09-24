
package employee

import (
	"companyapp/utils"
	"fmt"
)


func HireEmployee(name string, designation string) string {
	
	return fmt.Sprintf("%s has been hired as a %s at %s.", name, designation, utils.CompanyName)
}