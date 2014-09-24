package cas

import "crypto/tls"
import "fmt"
import "io/ioutil"
import "net/http"
import "net/url"

func RequestST(cas_server_url, username, password, service string) string {
	response := postFormData(cas_server_url + "/v1/tickets",
		url.Values{"username": {username}, "password": {password}})
	defer response.Body.Close()

	tgtLocation := response.Header.Get("Location")

	stResponse := postFormData(tgtLocation, url.Values{"service": {service}})
	defer stResponse.Body.Close()

	st, err := ioutil.ReadAll(stResponse.Body)
	if err != nil {
		fmt.Print("Error reading ST response body:")
		fmt.Println(err)
	}

	if stResponse.StatusCode != 200 {
		fmt.Println("ST response should be 200 but is", stResponse.StatusCode)
	}

	return string(st)
}

func postFormData(url string, formData url.Values) *http.Response {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}
	response, err := client.PostForm(url, formData)

	if err != nil {
		fmt.Print("Error on POST ", url, "with form data", formData)
		fmt.Println(err)
	}

	return response
}
