package cohesity

import (
  "encoding/json"
  "errors"
  "fmt"
  "log"
  "strconv"

  "github.com/cohesity/management-sdk-go/apihelper"
  CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
  "github.com/cohesity/management-sdk-go/unirest-go"
  "github.com/hashicorp/terraform/helper/schema"
)

type LicenceClusterParams struct {
  LicenseString string `json:"licenseString"`
}

type ClassifiedAccountInfo struct {
  //Timestamp at which license key was validated
  ActivationTimestamp int64 `json:activationTimestamp`
  // If cluster is claimed
  IsClusterActivated bool `json:isClusterActivated`
}

func resourceCohesityClusterLicense() *schema.Resource {
  return &schema.Resource{
    Create: resourceCohesityClusterLicenseCreate,
    Read:   resourceCohesityClusterLicenseRead,
    Delete: resourceCohesityClusterLicenseDelete,
    Update: resourceCohesityClusterLicenseUpdate,
    Schema: map[string]*schema.Schema{
      "license_key": {
        Type:        schema.TypeString,
        Required:    true,
        Sensitive:   true,
        DefaultFunc: schema.EnvDefaultFunc("VIRTUAL_COHESITY_CLUSTER_LICENSE_KEY", ""),
        Description: "Cohesity license key to apply after cluster creation",
      },
    },
  }
}

func resourceCohesityClusterLicenseCreate(resourceData *schema.ResourceData, configMetaData interface{}) error {
  var cohesityConfig = configMetaData.(Config)
  client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
    cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
  if err != nil {
    log.Printf(err.Error())
    return errors.New("Failed to authenticate with Cohesity")
  }

  // Check if license is already applied on the cluster.
  currLicenseInfo, err2 := getCurrentLicenseInfo(client)
  if err2 == nil {
    if currLicenseInfo.IsClusterActivated {
      log.Printf("[INFO] Cluster is already licensed.")
      return errors.New("cluster is already licensed")
    }
  } else {
    log.Printf("[WARNING] Failed to verify if license is already applied. Proceeding further.")
  }

  var licenseKey = resourceData.Get("license_key").(string)

  // the endpoint path uri for license endpoint
  _pathUrl := "irisservices/api/v1/nexus/license/upload_key"

  // API Path
  _apiPath := fmt.Sprintf("https://%s/%s", *client.Configuration().ClusterVip(), _pathUrl)

  //prepare headers for the outgoing request
  headers := map[string]interface{}{
    "user-agent":    "cohesity-Go-sdk-1.1.0",
    "accept":        "application/json",
    "content-type":  "application/json; charset=utf-8",
    "Authorization": fmt.Sprintf("%s %s", *client.Configuration().AccessToken().TokenType, *client.Configuration().AccessToken().AccessToken),
  }

  //prepare API request
  var requestBodyLicence LicenceClusterParams
  requestBodyLicence.LicenseString = licenseKey

  _request := unirest.Post(_apiPath, headers, requestBodyLicence)
  //and invoke the API call request to fetch the response
  _response, err := unirest.AsString(_request, client.Configuration().SkipSSL())
  if err != nil {
    log.Printf("[FAILURE] Failed to apply license for the cluster. %s", err.Error())
    return err
  }

  //error handling using HTTP status codes
  if _response.Code == 0 {
    err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
  } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
    err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
  }
  if err != nil {
    log.Printf("[FAILURE] Failed to apply license for the cluster %s", err.Error())
    return err
  }

  resourceData.SetId(strconv.FormatInt(currLicenseInfo.ActivationTimestamp, 10))
  return nil
}

func resourceCohesityClusterLicenseRead(resourceData *schema.ResourceData, configMetaData interface{}) error {
  log.Printf("[DEBUG] resourceCohesityClusterLicenseRead() called.")
  var cohesityConfig = configMetaData.(Config)
  client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
    cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
  if err != nil {
    log.Printf(err.Error())
    return errors.New("failed to authenticate with Cohesity")
  }

  // Check if license is already applied on the cluster.
  currLicenseInfo, err2 := getCurrentLicenseInfo(client)
  if err2 == nil {
    if currLicenseInfo.IsClusterActivated {
      log.Printf("[INFO] Cluster is already licensed.")
      if strconv.FormatInt(currLicenseInfo.ActivationTimestamp, 10) == resourceData.Id() {
        return nil
      } else {
        log.Printf("[INFO] Different instance of license information found.")
        resourceData.SetId("")
        return nil
      }
    } else {
      log.Printf("[INFO] Cluster isn't licensed.")
      resourceData.SetId("")
      return nil
    }
  } else {
    log.Printf("[WARNING] Failed to verify if license is already applied. Unsetting the current state.")
    resourceData.SetId("")
    return nil
  }

  return nil
}

func resourceCohesityClusterLicenseDelete(resourceData *schema.ResourceData, configMetaData interface{}) error {
  log.Printf("[DEBUG] resourceCohesityClusterLicenseDelete() called.")
  resourceData.SetId("")
  return nil
}

func resourceCohesityClusterLicenseUpdate(resourceData *schema.ResourceData, configMetaData interface{}) error {
  return nil
}

func getCurrentLicenseInfo(client CohesityManagementSdk.COHESITYMANAGEMENTSDK) (*ClassifiedAccountInfo, error) {
  log.Printf("[DEBUG] getCurrentLicenseInfo() called.")
  // the endpoint path uri for license endpoint
  _pathUrl := "irisservices/api/v1/nexus/account/is_classified"

  // API Path
  _apiPath := fmt.Sprintf("https://%s/%s", *client.Configuration().ClusterVip(), _pathUrl)

  //prepare headers for the outgoing request
  headers := map[string]interface{}{
    "user-agent":    "cohesity-Go-sdk-1.1.0",
    "accept":        "application/json",
    "content-type":  "application/json; charset=utf-8",
    "Authorization": fmt.Sprintf("%s %s", *client.Configuration().AccessToken().TokenType, *client.Configuration().AccessToken().AccessToken),
  }

  log.Printf("[DEBUG] Calling API %s", _apiPath)
  //prepare API request
  _request := unirest.Get(_apiPath, headers)
  //and invoke the API call request to fetch the response
  _response, err := unirest.AsString(_request, client.Configuration().SkipSSL())
  if err != nil {
    log.Printf("[Error] Failed to check if license is applied on the cluster. Error %s", err.Error())
    return nil, err
  }

  //error handling using HTTP status codes
  if _response.Code == 0 {
    err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
  } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
    err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
  }
  if err != nil {
    log.Printf("[Error] Failed to check if license is applied on the cluster. Error %s", err.Error())
    return nil, err
  }

  //returning the response
  var retVal *ClassifiedAccountInfo = &ClassifiedAccountInfo{}
  err = json.Unmarshal(_response.RawBody, &retVal)

  if err != nil {
    log.Printf("[Error] Failed to parse the response and verify if license is applied for cluster. Error %s", err.Error())
    return nil, nil
  }

  log.Printf("[INFO] Is cluster licensed : %t", retVal.IsClusterActivated)
  return retVal, nil
}
