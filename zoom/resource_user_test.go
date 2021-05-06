package zoom

import(
	"fmt"
	"terraform-provider-zoom/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
	"log"
)


func TestAccItem_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckItemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zoom_user.user1", "email", "tapendrakmr39@gmail.com"),
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
  email        = "tapendrakmr39@gmail.com"
  first_name   = "Ekansh"
  last_name    = "Singh"
  type         =  1
}
`)
}


func testAccCheckItemDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zoom_user" {
			continue
		}

		orderID := rs.Primary.ID

		err := c.DeleteItem(orderID)
		if err != nil {
			log.Println("[ERROR]: ",err)
			return err
		}
	}

	return nil
}


func TestAccItem_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckItemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemUpdatePre(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"zoom_user.user1", "email", "ekansh336@gmail.com"),
					resource.TestCheckResourceAttr(
						"zoom_user.user1", "first_name", "Ekansh"),
					resource.TestCheckResourceAttr(
						"zoom_user.user1", "last_name", "Singh"),
					
				),
			},
			{
				Config: testAccCheckItemUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"zoom_user.user1", "email", "ekansh336@gmail.com"),
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
	email        = "ekansh336@gmail.com"
	first_name   = "Ekansh"
	last_name    = "Singh"
	active       = "activate"
	type         =  1
}
`)
}

func testAccCheckItemUpdatePost() string {
	return fmt.Sprintf(`
resource "zoom_user" "user1" {
	email        = "ekansh336@gmail.com"
	first_name   = "Ekansh"
	last_name    = "kumar"
	active       = "activate"
	type         =  1
}
`)
}
