/*
 * A Bit of Everything
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Contact: none@example.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package abe

type ExamplepbUpdateV2Request struct {

	Abe *ExamplepbABitOfEverything `json:"abe,omitempty"`

	UpdateMask *ProtobufFieldMask `json:"update_mask,omitempty"`
}
