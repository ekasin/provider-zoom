package client

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestClient_GetItem(t *testing.T) {
	testCases := []struct {
		testName     string
		itemName     string
		seedData     map[string]User
		expectErr    bool
		expectedResp *User
	}{
		{
			testName: "user exists",
			itemName: "tapendrasingh66@gmail.com",
			seedData: map[string]User{
				"ekansh0786@gmail.com": {
					Id:        "qPinZDw3TJG6_6eZfiYpJQ",
					EmailId:   "tapendrasingh66@gmail.com",
					FirstName: "tapendra",
					LastName:  "kumar",
					Type:        1,
				},
			},
			expectErr: false,
			expectedResp: &User{
				Id:        "qPinZDw3TJG6_6eZfiYpJQ",
				EmailId:   "tapendrasingh66@gmail.com",
				FirstName: "tapendra",
				LastName:  "kumar",
				Type:        1,
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
			client := NewClient(os.Getenv("ZOOM_TOKEN"))
			item, err := client.GetItem(tc.itemName)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, item)
		})
	}
}


func TestClient_NewItem(t *testing.T) {
	testCases := []struct {
		testName  string
		newItem   *User
		seedData  map[string]User
		expectErr bool
	}{
		{
			testName: "success",
			newItem: &User{
				Id:        "qPinZDw3TJG6_6eZfiYpJQ",
				EmailId:   "tapendrasingh66@gmail.com",
				FirstName: "tapendra",
				LastName:  "kumar",
				Type:        1,
			},
			seedData:  nil,
			expectErr: false,
		},
		{
			testName: "item already exists",
			newItem: &User{
				Id:        "qPinZDw3TJG6_6eZfiYpJQ",
				EmailId:   "tapendrasingh66@gmail.com",
				FirstName: "tapendra",
				LastName:  "kumar",
				Type:        1,
			},
			seedData: map[string]User{
				"item1": {
					Id:        "qPinZDw3TJG6_6eZfiYpJQ",
					EmailId:   "tapendrasingh66@gmail.com",
					FirstName: "tapendra",
					LastName:  "kumar",
					Type:        1,
				},
			},
			expectErr: true,
		},
		
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("ZOOM_TOKEN"))

			err := client.NewItem(tc.newItem)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			item, err := client.GetItem(tc.newItem.EmailId)
			assert.NoError(t, err)
			assert.Equal(t, tc.newItem, item)
		})
	}
}


func TestClient_UpdateItem(t *testing.T) {
	testCases := []struct {
		testName    string
		updatedItem *User
		seedData    map[string]User
		expectErr   bool
	}{
		{
			testName: "item exists",
			updatedItem: &User{
				Id:        "qPinZDw3TJG6_6eZfiYpJQ",
					EmailId:   "tapendrasingh66@gmail.com",
					FirstName: "tapendra",
					LastName:  "kumar",
					Type:        1,
			},
			seedData: map[string]User{
				"item1": {
					Id:        "qPinZDw3TJG6_6eZfiYpJQ",
					EmailId:   "tapendrasingh66@gmail.com",
					FirstName: "tapendra",
					LastName:  "kumar",
					Type:        1,
				},
			},
			expectErr: false,
		},
		{
			testName: "item does not exist",
			updatedItem: &User{
				Id :       "dfhjjddfjsd",
				EmailId:   "ui17ec38@iitsurat.ac.in",
				FirstName: "ekansh",
				LastName:  "rock",
				Type:        1,
			},
			seedData:  nil,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("ZOOM_TOKEN"))
			err := client.UpdateItem(tc.updatedItem)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			item, err := client.GetItem(tc.updatedItem.EmailId)
			assert.NoError(t, err)
			assert.Equal(t, tc.updatedItem, item)
		})
	}
}


func TestClient_DeleteItem(t *testing.T) {
	testCases := []struct {
		testName  string
		itemName  string
		seedData  map[string]User
		expectErr bool
	}{
		{
			testName: "user exists",
			itemName: "tapendrakmr786@gmail.com",
			seedData: map[string]User{
				"user1": {
					Id:        "qPinZDw3TJG6_6eZfiYpJQ",
					EmailId:   "tapendrasingh66@gmail.com",
					FirstName: "tapendra",
					LastName:  "kumar",
					Type:        1,
				},
			},
			expectErr: false,
		},	
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("ZOOM_TOKEN"))
			err := client.DeleteItem(tc.itemName)
			log.Println(err)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			_, err = client.GetItem(tc.itemName)
			assert.Error(t, err)
		})
	}
}
