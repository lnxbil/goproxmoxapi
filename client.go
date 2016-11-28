// Package goproxmoxapi provides API for Proxmox Virtualisation Environment
package goproxmoxapi

import (
	"encoding/json"
	"net/http"
	"net/url"
  "crypto/tls"
  "fmt"
  "io/ioutil"
  "reflect"
  "strconv"
  "strings"
  "time"
)

// Token lifetime in minutes after creation
const TOKENLIFETIME = 120 // Minutes
// Refresh tokens when tokens are going to expire in less than or equal to  REFRESHBEFORE minutes
const REFRESHBEFORE = 5   // Minutes

// errorString is a trivial implementation of error.
type errorString struct {
  s string
}

// Type Client defines structure used for all API requests
// It holds following information:
// TargetNode: PVE node name to expose API on
// User, Pass, Realm: user information used for authentication
// Priviledges: json formatted user privilidges on target PVE node
// TokenDT, Ticket, CSRFPreventionToken: token information used to expose API
//
// According to documentation (https://pve.proxmox.com/wiki/Proxmox_VE_API)
// tokens are valid for 2 hours. To get a new ticket - the old ticket has to be
// used as password to the POST /access/ticket method. To be able to refresh
// tokens - ticket obtain time is being stored as TokenDT, along with Ticket
// and CSRFPreventionToken
type Client struct {
  TargetNode string // URL target node
  User, Pass, Realm string // Authentication information
  Priviledges map[string]interface{} // User priveleges
  TokenDT time.Time // Authentication token creation Time
  Ticket, CSRFPreventionToken string // Obtained Authentication tokens
  *http.Client
}

// creates new PVE client
func New(username, password, realm, tnode string) (*Client, error) {
	var csrfpreventiontoken, ticket string
  var privs map[string]interface{}

  // Prepare data for this request
  startt := time.Now()
	apiUrl := "https://" + tnode + ":8006"
	resource := "/api2/json/access/ticket"
	data := url.Values{}
  data.Set("username", username + "@" + realm)
  data.Add("password", password)

  // Prepare client
	tr := &http.Transport{ TLSClientConfig: &tls.Config{ InsecureSkipVerify: true }, }
	client := &http.Client{ Transport: tr }

  // Prepare URL
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
    return nil, err
	}
	u.Path = resource
	urlStr := fmt.Sprintf("%v", u)

  // Generate new request
	r, err := http.NewRequest("POST", urlStr, strings.NewReader( data.Encode()) )
	if err != nil {
    return nil, err
	}

  // Request ticket and authorization information
	resp, err := client.Do(r)
	if err != nil {
    return nil, err
	}
  defer resp.Body.Close()

  // Raise error if response status is not 200
  if resp.StatusCode != 200 {
    return nil, &errorString{ resp.Status, }
  }

  // Obtain tickets from returned body
  body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
    return nil, err
	} else {
    b := []byte( string(body) )
		var f interface{}
		err := json.Unmarshal(b, &f)

    if err != nil {
			return nil, err
		} else {
			jsondata := f.(map[string]interface{})["data"].(map[string]interface{})
      privs = jsondata["cap"].(map[string]interface{})
			csrfpreventiontoken = jsondata["CSRFPreventionToken"].(string)
			ticket = jsondata["ticket"].(string)
		}
	}

  // Return resulting Client structure
	return &Client{
		tnode,                               // Target node
    username, password, realm, privs,    // User information passed
    startt, ticket, csrfpreventiontoken, // Obtained ticket/token information
		client,                              // Resulting http.Client structure
	}, nil
}

// unconditionally refreshes API tokens when called
func (c *Client) RefreshToken() error {
  // This function is encapsulated within NewRequest call prior doing any operation
  c, err := New(c.User, c.Ticket, c.Realm, c.TargetNode)
  if err != nil {
		return err
	} else {
		return nil
	}
}

// creates request to PVE API and returns request response code and body for a given call.
// Response and body interpretation is on the caller shoulders.
func (c *Client) NewRequest(method, path string, data url.Values) (int, []byte, error) {
  // Refresh token if our token is going to expire "soon" (in 5 min)
  if time.Since( c.TokenDT ).Minutes() <= REFRESHBEFORE {
    err := c.RefreshToken()
    if err != nil {
      return 0, nil, err
    }
  }

  // Prepare URL
	apiUrl := "https://" + c.TargetNode + ":8006"
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
    return 0, nil, err
	}

  // Add requested Path to URL
	u.Path = path
	r, err := http.NewRequest(method, u.String(), strings.NewReader( data.Encode()) )
	if err != nil {
    return 0, nil, err
	}

  // Add CSRFPreventionToken to the request header if not it is not "GET"
  if method != "GET" {
    r.Header.Add("CSRFPreventionToken", c.CSRFPreventionToken)
  }

  // Add authentication cookie to the request
  token_expiration := c.TokenDT.Add( time.Duration( TOKENLIFETIME )*time.Minute )
  cookie := http.Cookie{Name: "PVEAuthCookie", Value: c.Ticket, Expires: token_expiration}
  r.AddCookie(&cookie)

  // Fetch data by requesting
	resp, err := c.Do(r)
	if err != nil {
    return resp.StatusCode, nil, err
	}
  defer resp.Body.Close()

  // Raise error if response status is not 200
  if resp.StatusCode != 200 {
    return resp.StatusCode, nil, &errorString{ resp.Status, }
  }

  // Obtain body of the response and return it, caller has to know what to do next
  rbody, err := ioutil.ReadAll(resp.Body)
  // Any Error reading body ?
	if err != nil {
    return resp.StatusCode, nil, err
	} else {
    return resp.StatusCode, rbody, err
	}
}

// convert bool to string "1" for true and "0" for false
func btoa(b bool) string {
  if b {
    return "1"
  } else {
    return "0"
  }
}

// convert string to bool, "0" to false and any other string to true
func atobf(a string) bool {
  if a=="0" {
    return false
  } else {
    return true
  }
}

// convert string to bool, "1" to true and any other string to false
func atobt(a string) bool {
  if a=="1" {
    return true
  } else {
    return false
  }
}

// Return error for a struct which contains only a string
func (e *errorString) Error() string {
  return e.s
}

// Cleans up url.Values from unwanted parameters.
// Two lists: wanted and unwanted parameter names can be passed.
// If wanted list is not empty, unwanted list will not be processed.
func structCleanUp(v url.Values, wanted []string, unwanted []string) (values url.Values) {
  vv := url.Values{}
  if len(wanted) > 0 {
    // If wanted list passed - process only wanted
    for _, wtd := range wanted {
      val := v.Get( strings.ToLower(wtd) )
      if val != "" {
        vv.Add( strings.ToLower(wtd), val )
      }
    }
  } else {
    // Delete unwanted parameters if any
    // This does a side effect of removing entries from original list of params,
    // however it suits our purpose anyway
    for _, unwtd := range unwanted {
      v.Del( strings.ToLower( unwtd ) )
    }
    vv = v
  }
/* extra cleanup entries with empty values  
  for key, val := range vv {
    if val[0] == "" {
      v.Del( key )
    }
  }
*/
  return vv
}

// Converts a struct to a map (to a url.Values string map)
// source: https://gist.github.com/tonyhb/5819315
// limitation: This func cannot deal with embedded structs.
func structToMap(i interface{}, wanted []string, unwanted []string) (values url.Values) {
//func structToMap(i interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for j := 0; j < iVal.NumField(); j++ {
		f := iVal.Field(j)
		// You can use tags here...
		// tag := typ.Field(j).Tag.Get("tagname")
		// Convert each type into a string for the url.Values string map
		var v string
		switch f.Interface().(type) {
		case []string:
      v = strings.Join(f.Interface().([]string), ",")
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}
    // Append field only if its value defined
    if v != "" {
      values.Set(strings.ToLower(typ.Field(j).Name), v)
    }
	}
  return structCleanUp(values, wanted, unwanted)
}

// Returns unmarhsled request data
func dataUnmarshal(rbody []byte, v interface{}) error {
  var f map[string]interface{}

  err := json.Unmarshal(rbody, &f)
  zz, err := json.Marshal(f["data"])
  err = json.Unmarshal( zz, &v )

  return err
}
