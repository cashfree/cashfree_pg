/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2023-08-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// Execute executes the request
//  @return ReconEntity
func PGFetchRecon(xApiVersion *string, fetchReconRequest *FetchReconRequest,  contentType *string, xRequestId *string, xIdempotencyKey *string, accept *string, httpClient *http.Client) (*ReconEntity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReconEntity
	)

	if XEnableErrorAnalytics {
		SetupSentry(XEnvironment)
		defer CaptureError("PGFetchRecon")
	}

	ctx := context.Background()

	client := NewAPIClient(NewConfiguration())
	if httpClient != nil {
		client.cfg.HTTPClient = httpClient
	}

	localBasePath := client.cfg.Servers[int(XEnvironment)].URL

	localVarPath := localBasePath + "/recon"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if xApiVersion == nil {
		return localVarReturnValue, nil, reportError("xApiVersion is required and must be specified")
	}
	if fetchReconRequest == nil {
		return localVarReturnValue, nil, reportError("fetchReconRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if contentType != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", contentType, "")
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-version", xApiVersion, "")
	if xRequestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-request-id", xRequestId, "")
	}
	if xIdempotencyKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-idempotency-key", xIdempotencyKey, "")
	}
	if accept != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", accept, "")
	}
	// body params
	localVarPostBody = fetchReconRequest

if XPartnerMerchantId != nil {
	localVarHeaderParams["x-partner-merchantid"] = *XPartnerMerchantId
}

if XClientId != nil {
	localVarHeaderParams["x-client-id"] = *XClientId
}

if XClientSignature != nil {
	localVarHeaderParams["x-client-signature"] = *XClientSignature
}

if XClientSecret != nil {
	localVarHeaderParams["x-client-secret"] = *XClientSecret
}

if XPartnerApiKey != nil {
	localVarHeaderParams["x-partner-apikey"] = *XPartnerApiKey
}
	req, err := client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequestError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError404
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiError409
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v IdempotencyError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RateLimitError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// With Context
// Execute executes the request
//  @return ReconEntity
func PGFetchReconWithContext(ctx context.Context, xApiVersion *string, fetchReconRequest *FetchReconRequest,  contentType *string, xRequestId *string, xIdempotencyKey *string, accept *string, httpClient *http.Client) (*ReconEntity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReconEntity
	)

	if XEnableErrorAnalytics {
		SetupSentry(XEnvironment)
		defer CaptureError("PGFetchRecon")
	}

	client := NewAPIClient(NewConfiguration())
	if httpClient != nil {
		client.cfg.HTTPClient = httpClient
	}

	localBasePath := client.cfg.Servers[int(XEnvironment)].URL

	localVarPath := localBasePath + "/recon"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if xApiVersion == nil {
		return localVarReturnValue, nil, reportError("xApiVersion is required and must be specified")
	}
	if fetchReconRequest == nil {
		return localVarReturnValue, nil, reportError("fetchReconRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if contentType != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", contentType, "")
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-version", xApiVersion, "")
	if xRequestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-request-id", xRequestId, "")
	}
	if xIdempotencyKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-idempotency-key", xIdempotencyKey, "")
	}
	if accept != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", accept, "")
	}
	// body params
	localVarPostBody = fetchReconRequest

if XPartnerMerchantId != nil {
	localVarHeaderParams["x-partner-merchantid"] = *XPartnerMerchantId
}

if XClientId != nil {
	localVarHeaderParams["x-client-id"] = *XClientId
}

if XClientSignature != nil {
	localVarHeaderParams["x-client-signature"] = *XClientSignature
}

if XClientSecret != nil {
	localVarHeaderParams["x-client-secret"] = *XClientSecret
}

if XPartnerApiKey != nil {
	localVarHeaderParams["x-partner-apikey"] = *XPartnerApiKey
}

	localVarHeaderParams["x-sdk-platform"] = "gosdk-4.2.0"
	req, err := client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequestError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError404
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiError409
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v IdempotencyError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RateLimitError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
// With Context


