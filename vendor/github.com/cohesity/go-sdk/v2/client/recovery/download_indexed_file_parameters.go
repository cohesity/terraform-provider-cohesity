// Code generated by go-swagger; DO NOT EDIT.

package recovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDownloadIndexedFileParams creates a new DownloadIndexedFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDownloadIndexedFileParams() *DownloadIndexedFileParams {
	return &DownloadIndexedFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadIndexedFileParamsWithTimeout creates a new DownloadIndexedFileParams object
// with the ability to set a timeout on a request.
func NewDownloadIndexedFileParamsWithTimeout(timeout time.Duration) *DownloadIndexedFileParams {
	return &DownloadIndexedFileParams{
		timeout: timeout,
	}
}

// NewDownloadIndexedFileParamsWithContext creates a new DownloadIndexedFileParams object
// with the ability to set a context for a request.
func NewDownloadIndexedFileParamsWithContext(ctx context.Context) *DownloadIndexedFileParams {
	return &DownloadIndexedFileParams{
		Context: ctx,
	}
}

// NewDownloadIndexedFileParamsWithHTTPClient creates a new DownloadIndexedFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewDownloadIndexedFileParamsWithHTTPClient(client *http.Client) *DownloadIndexedFileParams {
	return &DownloadIndexedFileParams{
		HTTPClient: client,
	}
}

/*
DownloadIndexedFileParams contains all the parameters to send to the API endpoint

	for the download indexed file operation.

	Typically these are written to a http.Request.
*/
type DownloadIndexedFileParams struct {

	/* FilePath.

	   Specifies the path to the file to download. If no path is specified and snapshot environment is kVMWare, VMX file for VMware will be downloaded. For other snapshot environments, this field must be specified.
	*/
	FilePath *string

	/* Length.

	   Specifies the length of bytes to download. This can not be greater than 8MB (8388608 byets)

	   Format: int64
	*/
	Length *int64

	/* NvramFile.

	   Specifies if NVRAM file for VMware should be downloaded.
	*/
	NvramFile *bool

	/* RetryAttempt.

	   Specifies the number of attempts the protection run took to create this file.

	   Format: int64
	*/
	RetryAttempt *int64

	/* SnapshotsID.

	   Specifies the snapshot id to download from.
	*/
	SnapshotsID string

	/* StartOffset.

	   Specifies the start offset of file chunk to be downloaded.

	   Format: int64
	*/
	StartOffset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the download indexed file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadIndexedFileParams) WithDefaults() *DownloadIndexedFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the download indexed file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadIndexedFileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the download indexed file params
func (o *DownloadIndexedFileParams) WithTimeout(timeout time.Duration) *DownloadIndexedFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download indexed file params
func (o *DownloadIndexedFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download indexed file params
func (o *DownloadIndexedFileParams) WithContext(ctx context.Context) *DownloadIndexedFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download indexed file params
func (o *DownloadIndexedFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download indexed file params
func (o *DownloadIndexedFileParams) WithHTTPClient(client *http.Client) *DownloadIndexedFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download indexed file params
func (o *DownloadIndexedFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilePath adds the filePath to the download indexed file params
func (o *DownloadIndexedFileParams) WithFilePath(filePath *string) *DownloadIndexedFileParams {
	o.SetFilePath(filePath)
	return o
}

// SetFilePath adds the filePath to the download indexed file params
func (o *DownloadIndexedFileParams) SetFilePath(filePath *string) {
	o.FilePath = filePath
}

// WithLength adds the length to the download indexed file params
func (o *DownloadIndexedFileParams) WithLength(length *int64) *DownloadIndexedFileParams {
	o.SetLength(length)
	return o
}

// SetLength adds the length to the download indexed file params
func (o *DownloadIndexedFileParams) SetLength(length *int64) {
	o.Length = length
}

// WithNvramFile adds the nvramFile to the download indexed file params
func (o *DownloadIndexedFileParams) WithNvramFile(nvramFile *bool) *DownloadIndexedFileParams {
	o.SetNvramFile(nvramFile)
	return o
}

// SetNvramFile adds the nvramFile to the download indexed file params
func (o *DownloadIndexedFileParams) SetNvramFile(nvramFile *bool) {
	o.NvramFile = nvramFile
}

// WithRetryAttempt adds the retryAttempt to the download indexed file params
func (o *DownloadIndexedFileParams) WithRetryAttempt(retryAttempt *int64) *DownloadIndexedFileParams {
	o.SetRetryAttempt(retryAttempt)
	return o
}

// SetRetryAttempt adds the retryAttempt to the download indexed file params
func (o *DownloadIndexedFileParams) SetRetryAttempt(retryAttempt *int64) {
	o.RetryAttempt = retryAttempt
}

// WithSnapshotsID adds the snapshotsID to the download indexed file params
func (o *DownloadIndexedFileParams) WithSnapshotsID(snapshotsID string) *DownloadIndexedFileParams {
	o.SetSnapshotsID(snapshotsID)
	return o
}

// SetSnapshotsID adds the snapshotsId to the download indexed file params
func (o *DownloadIndexedFileParams) SetSnapshotsID(snapshotsID string) {
	o.SnapshotsID = snapshotsID
}

// WithStartOffset adds the startOffset to the download indexed file params
func (o *DownloadIndexedFileParams) WithStartOffset(startOffset *int64) *DownloadIndexedFileParams {
	o.SetStartOffset(startOffset)
	return o
}

// SetStartOffset adds the startOffset to the download indexed file params
func (o *DownloadIndexedFileParams) SetStartOffset(startOffset *int64) {
	o.StartOffset = startOffset
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadIndexedFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilePath != nil {

		// query param filePath
		var qrFilePath string

		if o.FilePath != nil {
			qrFilePath = *o.FilePath
		}
		qFilePath := qrFilePath
		if qFilePath != "" {

			if err := r.SetQueryParam("filePath", qFilePath); err != nil {
				return err
			}
		}
	}

	if o.Length != nil {

		// query param length
		var qrLength int64

		if o.Length != nil {
			qrLength = *o.Length
		}
		qLength := swag.FormatInt64(qrLength)
		if qLength != "" {

			if err := r.SetQueryParam("length", qLength); err != nil {
				return err
			}
		}
	}

	if o.NvramFile != nil {

		// query param nvramFile
		var qrNvramFile bool

		if o.NvramFile != nil {
			qrNvramFile = *o.NvramFile
		}
		qNvramFile := swag.FormatBool(qrNvramFile)
		if qNvramFile != "" {

			if err := r.SetQueryParam("nvramFile", qNvramFile); err != nil {
				return err
			}
		}
	}

	if o.RetryAttempt != nil {

		// query param retryAttempt
		var qrRetryAttempt int64

		if o.RetryAttempt != nil {
			qrRetryAttempt = *o.RetryAttempt
		}
		qRetryAttempt := swag.FormatInt64(qrRetryAttempt)
		if qRetryAttempt != "" {

			if err := r.SetQueryParam("retryAttempt", qRetryAttempt); err != nil {
				return err
			}
		}
	}

	// path param snapshotsId
	if err := r.SetPathParam("snapshotsId", o.SnapshotsID); err != nil {
		return err
	}

	if o.StartOffset != nil {

		// query param startOffset
		var qrStartOffset int64

		if o.StartOffset != nil {
			qrStartOffset = *o.StartOffset
		}
		qStartOffset := swag.FormatInt64(qrStartOffset)
		if qStartOffset != "" {

			if err := r.SetQueryParam("startOffset", qStartOffset); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
