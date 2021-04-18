package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"terraform-provider-zoom/server"
	"os"
)

// HostURL - Default Hashicups URL
//const HostURL string = "https://api.zoom.us/v2/users"

// Client -
type Client struct {
	hostname   string
	authToken  string
	httpClient *http.Client
}

// AuthStruct -

// NewClient -
func NewClient(hostname string, token string) *Client {
	return &Client{
		hostname:   hostname,
		authToken:  token,
		httpClient: &http.Client{},
	}
}

func (c *Client) NewItem(item *server.Item) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return err
	}
	_, err = c.httpRequest("POST", buf, item)
	if err != nil {
		return err
	}
	return nil
}




///////////////////////////////////////////////////Get User//////////////////////////////////////////////////////////////////////



func (c *Client) GetItem(name string) (*server.Item, error) {
	//fmt.Println("name",name)
	body, err := c.gethttpRequest(fmt.Sprintf("%v", name), "GET", bytes.Buffer{})
	//fmt.Println(err)
	if err != nil {
		return nil, err
	}
	item := &server.Item{}
	err = json.NewDecoder(body).Decode(item)
	//fmt.Println("heloo world")
	//fmt.Println(item.FirstName)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (c *Client) gethttpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.getrequestPath(path), &body)
	if err != nil {
		return nil, err
	}

	//var bearer = "Bearer " + c.authToken
	//var str1 string = os.Getenv("ZOOM_TOKEN")
	var bearer = "Bearer " + os.Getenv("ZOOM_TOKEN")

	//var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTkyOTI4NTMsImlhdCI6MTYxODY4ODA1M30.lRrdfygWH8pgGcm0l4H3MCO1Uma7NGQ-r1TnobrQL-E"

	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	//fmt.Println("new error",err)
	if err != nil {
		return nil, err
	}
	
	//fmt.Println("response",resp)

	fmt.Println("statuscode",resp.StatusCode)
	

	/*
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        fmt.Println("HTTP Status is in the 2xx range")
		return resp.Body, nil
    } else {
        fmt.Println("Argh! Broken")
		return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
    }
	*/


	
	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s", resp.StatusCode, respBody.String())
	}
	

	//fmt.Println("response body",resp.Body)

	return resp.Body, nil
}

func (c *Client) getrequestPath(path string) string {
	return fmt.Sprintf("%s/%s", os.Getenv("ZOOM_ADDRESS"), path)
}




/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////











/////////////////////////////////////////////////////////////update////////////////////////////////////////////////////////////////////////////////////////////


func (c *Client) UpdateItem(item *server.Item) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return err
	}
	_, err = c.updatehttpRequest(fmt.Sprintf("%s", item.EmailId), "PATCH", buf,item)
	if err != nil {
		return err
	}
	return nil
}


func (c *Client) updatehttpRequest(path,method string, body bytes.Buffer, item *server.Item) (closer io.ReadCloser, err error) {
	data := fmt.Sprintf("{\"first_name\":\"%s\",\"last_name\":\"%s\"}",item.FirstName,item.LastName)
	payload := strings.NewReader(data)

	req, err := http.NewRequest(method, c.updaterequestPath(path), payload)
	var str1 string
	str1 = "Bearer "
	var str2 string
	str2 = os.Getenv("ZOOM_TOKEN")
	//str2 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTkyOTI4NTMsImlhdCI6MTYxODY4ODA1M30.lRrdfygWH8pgGcm0l4H3MCO1Uma7NGQ-r1TnobrQL-E"
	req.Header.Add("Authorization", str1+str2)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        fmt.Println("HTTP Status is in the 2xx range")
		return resp.Body, nil
    } else {
        fmt.Println("Argh! Broken")
		return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
    }

	/*
	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s", resp.StatusCode, respBody.String())
	}
	*/
	return resp.Body, nil
}


func (c *Client) updaterequestPath(path string) string {
	//fmt.Println(c.hostname)
	return fmt.Sprintf("%s/%s",os.Getenv("ZOOM_ADDRESS"), path)

}


///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


///////////////////////////////////Delete/////////////////////////////////

func (c *Client) DeleteItem(userId string) error {
	_, err := c.deletehttpRequest(fmt.Sprintf("%s", userId), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}


func (c *Client) deletehttpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.deleterequestPath(path), &body)
	if err != nil {
		return nil, err
	}


	var str1 string
	str1 = "Bearer "
	var str2 string
	str2 = os.Getenv("ZOOM_TOKEN")
	//str2 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTkyOTI4NTMsImlhdCI6MTYxODY4ODA1M30.lRrdfygWH8pgGcm0l4H3MCO1Uma7NGQ-r1TnobrQL-E"
	req.Header.Add("Authorization", str1+str2)
	//client := &http.Client{}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}


	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        fmt.Println("HTTP Status is in the 2xx range")
		return resp.Body, nil
    } else {
        fmt.Println("Argh! Broken")
		return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
    }

	
	/*
	//fmt.Println("statuscode",resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v ", resp.StatusCode)
	}
	*/
	
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        fmt.Println("HTTP Status is in the 2xx range")
		return resp.Body, nil
    } else {
        fmt.Println("Argh! Broken")
		return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
    }

	return resp.Body, nil
}

func (c *Client) deleterequestPath(path string) string {
	return fmt.Sprintf("%s/%s", os.Getenv("ZOOM_ADDRESS"), path)
}



//////////////////////////////////////////////////////////////////////////////////////////////////////////


func (c *Client) httpRequest(method string, body bytes.Buffer, item *server.Item) (closer io.ReadCloser, err error) {
	data := fmt.Sprintf("{\"action\":\"create\",\"user_info\":{\"email\":\"%s\",\"type\":1,\"first_name\":\"%s\",\"last_name\":\"%s\"}}", item.EmailId, item.FirstName, item.LastName)

	payload := strings.NewReader(data)

	req, err := http.NewRequest(method, c.requestPath(), payload)
	var str1 string
	str1 = "Bearer "
	var str2 string
	str2 = os.Getenv("ZOOM_TOKEN")
	//str2 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTkyOTI4NTMsImlhdCI6MTYxODY4ODA1M30.lRrdfygWH8pgGcm0l4H3MCO1Uma7NGQ-r1TnobrQL-E"
	req.Header.Add("Authorization", str1+str2)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	
	/*
	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s", resp.StatusCode, respBody.String())
	}
	*/

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        fmt.Println("HTTP Status is in the 2xx range")
		return resp.Body, nil
    } else {
        fmt.Println("Argh! Broken")
		return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
    }

	//return resp.Body, nil
}

func (c *Client) requestPath() string {
	//fmt.Println(c.hostname)
	return fmt.Sprintf("%s?access_token=%s", os.Getenv("ZOOM_ADDRESS"), os.Getenv("ZOOM_TOKEN"))

}
