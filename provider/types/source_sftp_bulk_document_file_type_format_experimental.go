// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceSftpBulkDocumentFileTypeFormatExperimental struct {
	Processing             *SourceSftpBulkProcessing `tfsdk:"processing"`
	SkipUnprocessableFiles types.Bool                `tfsdk:"skip_unprocessable_files"`
	Strategy               types.String              `tfsdk:"strategy"`
}
