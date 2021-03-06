package validation

import (
	"regexp"

	"github.com/asaskevich/govalidator"
	"github.com/ibrahimakbar31/comment-api-go/core/model"
)

//ValidEmailFormat function to validate email
func ValidEmailFormat(input string) bool {
	checkIsEmail := govalidator.IsEmail(input)
	return checkIsEmail
}

//ValidUsernameFormat function to validate username
func ValidUsernameFormat(input string) bool {
	rgxUsername := regexp.MustCompile("^[a-zA-Z0-9_]*$")
	checkIsUsername := rgxUsername.MatchString(input)
	return checkIsUsername
}

//ValidMemberInOrganization function to validate member organization
func ValidMemberInOrganization(orgCode string, member model.Member) bool {
	for _, organizationMember := range member.OrganizationMembers {
		if organizationMember.Organization.Code == orgCode {
			return true
		}
	}
	return false
}
