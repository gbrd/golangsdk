package users

import (
	"fmt"

	"github.com/chnsz/golangsdk"
)

type CreateOpts struct {
	// Enterprise user name.
	// maxLength: 64
	// minLength: 1
	Name string `json:"name" required:"true"`
	// HUAWEI CLOUD meeting user account, if it is carried, it shall prevail; otherwise, it will be automatically generated in the background. The account is unique in the whole system
	// The account number can only contain uppercase and lowercase letters, numbers, _, -, ., and @ symbols, and cannot be pure numbers and a . sign after @.
	// maxLength: 64
	// minLength: 0
	// Description: used for account/password authentication
	Account string `json:"account,omitempty"`
	// Third-party User ID
	// Description: used in App ID authentication mode
	ThirdAccount string `json:"thirdAccount,omitempty"`
	// The country the phone number belongs to.
	// Default: chinaPR.
	// maxLength: 255
	// minLength: 0
	Country string `json:"country,omitempty"`
	// Department ID, if not carried, the default root department.
	// Default: 1
	// maxLength: 32
	// minLength: 0
	DeptCode string `json:"deptCode,omitempty"`
	// Description.
	// maxLength: 128
	// minLength: 0
	Description string `json:"desc,omitempty"`
	// Email.
	// maxLength: 255
	// minLength: 0
	// Unified email format, if the enterprise turns off the SMS function, the email is required,
	// otherwise the mobile phone and email are required.
	Email string `json:"email,omitempty"`
	// The English name of the enterprise user.
	// maxLength: 64
	// minLength: 0
	EnglishName string `json:"englishName,omitempty"`
	// User function bits.
	Function *UserFunction `json:"function,omitempty"`
	// Mobile phone number, country code must be added.
	// For example, the mobile phone in mainland China is "+86xxxxxxxxxxxx".
	// When filling in the mobile phone number, the "country" parameter is required.
	// Only pure numbers are allowed for mobile phone numbers.
	// Description: Fill in at least one mobile phone number or email address.
	// maxLength: 32
	// minLength: 0
	Phone string `json:"phone,omitempty"`
	// Whether to hide the phone number.
	// When set to true, the mobile phone number will not be displayed in the address book and conference.
	// Default: false
	HidePhone *bool `json:"hidePhone,omitempty"`
	// The password of the enterprise user account. If it is carried, the actual carrying shall prevail.
	// Otherwise, it will be generated by default in the background, and the password must meet the following
	// requirements:
	//   1, 8-32 bits
	//   2. It cannot be consistent with the positive and reverse order of the account
	//   3. Contains at least two character types: lowercase letters, uppercase letters, numbers,
	//      special characters (` ~ ! @ # $ % ^ & * ( ) - _ = + | [ { } ] ; : " ,' < . > / ?)
	Password string `json:"pwd,omitempty"`
	// Whether to send email and SMS notifications for account opening.
	// 0: do not send
	// Send without filling in or other values, and send by default
	// maxLength: 32
	// minLength: 0
	SendNotify string `json:"sendNotify,omitempty"`
	// Signature.
	// maxLength: 512
	// minLength: 0
	Signature string `json:"signature,omitempty"`
	// Address book sorting level, the lower the serial number, the higher the priority.
	// Default: 10000
	// maximum: 10000
	//minimum: 1
	SortLevel int `json:"sortLevel,omitempty"`
	// user status.
	//   0: normal
	//   1: disable
	// default: 0
	Status *int `json:"status,omitempty"`
	// Position (title).
	// maxLength: 32
	// minLength: 0
	Title string `json:"title,omitempty"`
	// Authorization token.
	Token string `json:"-" required:"true"`
}

type UserFunction struct {
	// Whether to enable the intelligent collaborative whiteboard function.
	// If it is enabled, it means that the account is used for the intelligent collaborative whiteboard,
	// which occupies the resources of the enterprise intelligent collaborative whiteboard.
	// If the resources are insufficient, it cannot be opened.
	// Default: false.
	EnableRoom bool `json:"enableRoom,omitempty"`
}

// Create is a method to create a new (enterprise) user using given parameters.
func Create(c *golangsdk.ServiceClient, opts CreateOpts) (*User, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	var r User
	_, err = c.Post(rootURL(c), b, &r, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"Content-Type":   "application/json;charset=UTF-8",
			"X-Access-Token": opts.Token,
		},
	})
	return &r, err
}

type GetOpts struct {
	// account account.
	// If it is an account/password authentication method, it refers to the HUAWEI CLOUD conference account
	// If it is the App ID authentication method, it refers to the third-party User ID
	Account string `json:"-"`
	// Account type.
	//   0: HUAWEI CLOUD conference account. Used for account/password authentication.
	//   1: Third-party User ID, used for App ID authentication.
	// default 0
	AccountType int `q:"accountType"`
	// Authorization token.
	Token string `json:"-"`
}

// Get is a method to create a new (enterprise) user using given parameters.
func Get(c *golangsdk.ServiceClient, opts GetOpts) (*User, error) {
	if opts.Account == "" || opts.Token == "" {
		return nil, fmt.Errorf("The account and authorization token must be supported.")
	}

	url := resourceURL(c, opts.Account)
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}
	url += query.String()

	var r User
	_, err = c.Get(url, &r, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"Content-Type":   "application/json;charset=UTF-8",
			"X-Access-Token": opts.Token,
		},
	})
	return &r, err
}

type UpdateOpts struct {
	// account account.
	// If it is an account/password authentication method, it refers to the HUAWEI CLOUD conference account
	// If it is the App ID authentication method, it refers to the third-party User ID
	Account string `json:"-" required:"true"`
	// Account type.
	//   0: HUAWEI CLOUD conference account. Used for account/password authentication.
	//   1: Third-party User ID, used for App ID authentication.
	// default 0
	AccountType *int `q:"accountType"`
	// The country the phone number belongs to.
	// Default: chinaPR.
	// maxLength: 255
	// minLength: 0
	Country string `json:"country,omitempty"`
	// Department ID, if not carried, the default root department.
	// Default: 1
	// maxLength: 32
	// minLength: 0
	DeptCode *string `json:"deptCode,omitempty"`
	// Description.
	// maxLength: 128
	// minLength: 0
	Description *string `json:"desc,omitempty"`
	// Email.
	// maxLength：255
	// minLength：0
	Email *string `json:"email,omitempty"`
	// The English name of the enterprise user.
	// maxLength: 64
	// minLength: 0
	EnglishName *string `json:"englishName,omitempty"`
	// Whether to hide the phone number.
	// Default: false
	HidePhone *bool `json:"hidePhone,omitempty"`
	// Enterprise user name.
	// maxLength: 64
	// minLength: 1
	Name string `json:"name,omitempty"`
	// Mobile phone number, country code must be added.
	// For example, the mobile phone in mainland China is "+86xxxxxxxxxxxx".
	// When filling in the mobile phone number, the "country" parameter is required.
	// Only pure numbers are allowed for mobile phone numbers.
	// Description: Fill in at least one mobile phone number or email address.
	// maxLength: 32
	// minLength: 0
	Phone string `json:"phone,omitempty"`
	// Signature.
	// maxLength：512
	// minLength：0
	Signature *string `json:"signature,omitempty"`
	// Address book sorting level, the lower the serial number, the higher the priority.
	// Default: 10000
	// maximum: 10000
	//minimum: 1
	SortLevel int `json:"sortLevel,omitempty"`
	// user status.
	// 0: normal
	// 1: disable
	// default: 0
	Status *int `json:"status,omitempty"`
	// Position (Title).
	// maxLength: 32
	// minLength: 0
	Title *string `json:"title,omitempty"`
	// Personal meeting ID, if not carried, it will be generated by default in the background.
	// maxLength: 32
	// minLength: 0
	VmrId *string `json:"vmrId,omitempty"`
	// Authorization token.
	Token string `json:"-" required:"true"`
}

// Update is a method to create a new (enterprise) user using given parameters.
func Update(c *golangsdk.ServiceClient, opts UpdateOpts) (*User, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	var r User
	_, err = c.Put(resourceURL(c, opts.Account), b, &r, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"Content-Type":   "application/json;charset=UTF-8",
			"X-Access-Token": opts.Token,
		},
	})
	return &r, err
}

type DeleteOpts struct {
	// Account type.
	//   0: HUAWEI CLOUD conference account. Used for account/password authentication.
	//   1: Third-party User ID, used for App ID authentication.
	// default 0
	AccountType *int `q:"accountType"`
	// Authorization token.
	Token string `json:"-"`
}

// BatchDelete is a method to delete all (enterprise) users using given list.
func BatchDelete(c *golangsdk.ServiceClient, opts DeleteOpts, accounts []string) error {
	url := deleteURL(c)
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return err
	}
	url += query.String()

	if opts.Token == "" {
		return fmt.Errorf("The authorization token must be supported.")
	}

	_, err = c.Post(url, accounts, nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"Content-Type":   "application/json;charset=UTF-8",
			"X-Access-Token": opts.Token,
		},
	})
	return err
}
