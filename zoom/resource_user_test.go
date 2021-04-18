package zoom

import(
	"os"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"fmt"
	"terraform-provider-zoom/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	//"regexp"
	"testing"
)



/////////////////////////////////////////////////////////////////////////

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"zoom": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T)  {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("ZOOM_ADDRESS"); v == "" {
		t.Fatal("ZOOM_ADDRESS must be set for acceptance tests")
	}
	if v := os.Getenv("ZOOM_TOKEN"); v == "" {
		t.Fatal("ZOOM_TOKEN must be set for acceptance tests")
	}
}
















//////////////////////////////












func TestAccItem_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		//CheckDestroy: testAccCheckItemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExampleItemExists("zoom_user.user1"),
					resource.TestCheckResourceAttr("zoom_user.user1", "email", "tapendrakmr1234@gmail.com"),
					resource.TestCheckResourceAttr("zoom_user.user1", "first_name", "Ekansh"),
					resource.TestCheckResourceAttr("zoom_user.user1", "last_name", "Singh"),
				),
			},
		},
	})
}


func testAccCheckItemBasic() string {
	return fmt.Sprintf(`
resource "zoom_user" "user1" {
  email        = "tapendrakmr1234@gmail.com"
  first_name   = "Ekansh"
  last_name    = "Singh"
}
`)
}




////////////////////////////////////////TESTING FOR DELETE OPERATION//////////////////////////////////////////


func testAccCheckItemDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zoom_user" {
			continue
		}

		orderID := rs.Primary.ID

		err := c.DeleteItem(orderID)
		if err != nil {
			return err
		}
	}

	return nil
}





///////////////////////////////////////////////////////////////////////////////////////////////











//////////////////////////////////TESTING FOR CREATE OPERATION/////////////////////////////////////////////////////
func testAccCheckExampleItemExists(resource string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("Not found: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}
		name := rs.Primary.ID
		//name := "tapendrasingh66@gmail.com"
		apiClient := testAccProvider.Meta().(*client.Client)
		_, err := apiClient.GetItem(name)
		if err != nil {
			return fmt.Errorf("error fetching item with resource %s. %s", resource, err)
		}
		//fmt.Println("yesssssssss")
		//fmt.Println("body",xyz)
		return nil
	}
}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////




//////////////////////////testing multiple users///////////////////////////////////////

func TestAccItem_Multiple(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		//CheckDestroy: testAccCheckItemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemMultiple(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExampleItemExists("zoom_user.user1"),
					testAccCheckExampleItemExists("zoom_user.user2"),
				),
			},
		},
	})
}

func testAccCheckItemMultiple() string {
	return fmt.Sprintf(`
resource "zoom_user" "user1" {
	email        = "ekansh2786@gmail.com"
	first_name   = "Ekansh"
	last_name    = "Singh"
}
resource "zoom_user" "user2" {
	email        = "ekansh1786@gmail.com"
  	first_name   = "Ekansh"
  	last_name    = "Singh"
  }
`)
}


//////////////////////////////////////////////////////////////////////////////////////







//////////////////////////////////////////////TESTING FOR UPDATE OPERATION//////////////////////////////////////////////////////

func TestAccItem_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		//CheckDestroy: testAccCheckItemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemUpdatePre(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExampleItemExists("zoom_user.user1"),
					resource.TestCheckResourceAttr(
						"zoom_user.user1", "email", "ekansh1076@gmail.com"),
					resource.TestCheckResourceAttr(
						"zoom_user.user1", "first_name", "Ekansh"),
					resource.TestCheckResourceAttr(
						"zoom_user.user1", "last_name", "Singh"),
					
				),
			},
			{
				Config: testAccCheckItemUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExampleItemExists("zoom_user.user1"),
					resource.TestCheckResourceAttr(
						"zoom_user.user1", "email", "ekansh1076@gmail.com"),
					resource.TestCheckResourceAttr(
						"zoom_user.user1", "first_name", "Ekansh"),
					resource.TestCheckResourceAttr(
						"zoom_user.user1", "last_name", "kumar"),
				),
			},
		},
	})
}


func testAccCheckItemUpdatePre() string {
	return fmt.Sprintf(`
resource "zoom_user" "user1" {
	email        = "ekansh1076@gmail.com"
	first_name   = "Ekansh"
	last_name    = "Singh"
}
`)
}

func testAccCheckItemUpdatePost() string {
	return fmt.Sprintf(`
resource "zoom_user" "user1" {
	email        = "ekansh1076@gmail.com"
	first_name   = "Ekansh"
	last_name    = "kumar"
}
`)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////








////////////////////////////////////////////////////////////////////////////////
/*
var whiteSpaceRegex = regexp.MustCompile("email cannot contain whitespace")

func TestAccItem_WhitespaceName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckItemWhitespace(),
				ExpectError: whiteSpaceRegex,
			},
		},
	})
}

func testAccCheckItemWhitespace() string {
	return fmt.Sprintf(`
resource "example_item" "test_item" {
	email        = "ekansh086@gmail.com"
	first_name   = "Ekansh"
	last_name    = "kumar"
}
`)
}
*/
//////////////////////////////////////////////////////////////////////////////////