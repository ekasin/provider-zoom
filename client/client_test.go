package client

import(
	"terraform-provider-zoom/server"
	"github.com/stretchr/testify/assert"
	"testing"
	"log"
)


func TestClient_GetItem(t *testing.T) {
	testCases := []struct {
		testName     string
		itemName     string
		seedData     map[string]server.Item
		expectErr    bool
		expectedResp *server.Item
	}{
		{
			testName: "user exists",
			itemName: "ekansh0786@gmail.com",
			seedData: map[string]server.Item{
				"ekansh0786@gmail.com": {
					EmailId:   "ekansh0786@gmail.com",
					FirstName: "ekansh",
					LastName:  "rock",
				},
			},
			expectErr: false,
			expectedResp: &server.Item{
				EmailId:   "ekansh0786@gmail.com",
				FirstName: "ekansh",
				LastName:  "rock",
			},
		},
		
		{
			testName:     "user does not exist",
			itemName:     "ui17co14@iitsurat.ac.in",
			seedData:     nil,
			expectErr:    true,
			expectedResp: nil,
		},
		
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient("https://api.zoom.us/v2/users", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTkyOTI4NTMsImlhdCI6MTYxODY4ODA1M30.lRrdfygWH8pgGcm0l4H3MCO1Uma7NGQ-r1TnobrQL-E")

			item, err := client.GetItem(tc.itemName)
			if tc.expectErr {
				log.Println("[READ ERROR]: ", err)
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			log.Println("[READ ERROR]: ", err)
			assert.Equal(t, tc.expectedResp, item)
		})
	}
}



/////////////////////////////////////////////////////////////////////



func TestClient_NewItem(t *testing.T) {
	testCases := []struct {
		testName  string
		newItem   *server.Item
		seedData  map[string]server.Item
		expectErr bool
	}{
		{
			testName: "success",
			newItem: &server.Item{
				EmailId:   "ekansh0786@gmail.com",
				FirstName: "ekansh",
				LastName:  "rock",
			},
			seedData:  nil,
			expectErr: false,
		},
		{
			testName: "item already exists",
			newItem: &server.Item{
				EmailId:   "ekansh0786@gmail.com",
				FirstName: "ekansh",
				LastName:  "rock",
			},
			seedData: map[string]server.Item{
				"item1": {
				EmailId:   "ekansh0786@gmail.com",
				FirstName: "ekansh",
				LastName:  "rock",
				},
			},
			expectErr: true,
		},
		
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient("https://api.zoom.us/v2/users", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTkyOTI4NTMsImlhdCI6MTYxODY4ODA1M30.lRrdfygWH8pgGcm0l4H3MCO1Uma7NGQ-r1TnobrQL-E")


			err := client.NewItem(tc.newItem)
			if tc.expectErr {
				log.Println("[CREATE ERROR]: ", err)
				assert.Error(t, err)
				return
			}
			item, err := client.GetItem(tc.newItem.EmailId)
			assert.NoError(t, err)
			log.Println("[CREATE ERROR]: ", err)
			assert.Equal(t, tc.newItem, item)
		})
	}
}



/////////////////////////////////////////////////////////////////////


///////////////////////////////////////////////update//////////////////////////////////////


func TestClient_UpdateItem(t *testing.T) {
	testCases := []struct {
		testName    string
		updatedItem *server.Item
		seedData    map[string]server.Item
		expectErr   bool
	}{
		{
			testName: "item exists",
			updatedItem: &server.Item{
				EmailId:   "ekansh0786@gmail.com",
					FirstName: "ekansh",
					LastName:  "rock",
			},
			seedData: map[string]server.Item{
				"item1": {
					EmailId:   "ekansh0786@gmail.com",
					FirstName: "ekansh",
					LastName:  "rock",
				},
			},
			expectErr: false,
		},
		{
			testName: "item does not exist",
			updatedItem: &server.Item{
				EmailId:   "ui17ec38@iitsurat.ac.in",
				FirstName: "ekansh",
				LastName:  "rock",
			},
			seedData:  nil,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient("https://api.zoom.us/v2/users", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTkyOTI4NTMsImlhdCI6MTYxODY4ODA1M30.lRrdfygWH8pgGcm0l4H3MCO1Uma7NGQ-r1TnobrQL-E")
			err := client.UpdateItem(tc.updatedItem)
			if tc.expectErr {
				log.Println("[UPDATE ERROR]: ", err)
				assert.Error(t, err)
				return
			}
			item, err := client.GetItem(tc.updatedItem.EmailId)
			assert.NoError(t, err)
			log.Println("[UPDATE ERROR]: ", err)
			assert.Equal(t, tc.updatedItem, item)
		})
	}
}





////////////////////////////////////////////////////////////////////////////////////////////


////////////////////////////Delete Testing/////////////////////////

/*

func TestClient_DeleteItem(t *testing.T) {
	testCases := []struct {
		testName  string
		itemName  string
		seedData  map[string]server.User
		expectErr bool
	}{
		{
			testName: "user exists",
			itemName: "ekansh0786@gmail.com",
			seedData: map[string]server.User{
				"user1": {
					EmailId:   "ekansh0786@gmail.com",
					FirstName: "ekansh",
					LastName:  "rock",
				},
			},
			expectErr: false,
		},
		
		
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient("https://api.zoom.us/v2/users", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTkyOTI4NTMsImlhdCI6MTYxODY4ODA1M30.lRrdfygWH8pgGcm0l4H3MCO1Uma7NGQ-r1TnobrQL-E")
			
			err := client.DeleteItem(tc.itemName)
			fmt.Println(err)
			if tc.expectErr {
				log.Println("[DELETE ERROR]: ", err)
				assert.Error(t, err)
				return
			}
			_, err = client.GetItem(tc.itemName)
			log.Println("[DELETE ERROR]: ", err)
			assert.Error(t, err)
		})
	}
}

*/

//////////////////////////////////////////////////////////////////
