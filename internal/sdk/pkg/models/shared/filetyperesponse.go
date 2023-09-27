// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type FileTypeResponseFileInfoGroupIDs string

const (
	FileTypeResponseFileInfoGroupIDsA FileTypeResponseFileInfoGroupIDs = "A"
	FileTypeResponseFileInfoGroupIDsD FileTypeResponseFileInfoGroupIDs = "D"
	FileTypeResponseFileInfoGroupIDsE FileTypeResponseFileInfoGroupIDs = "E"
	FileTypeResponseFileInfoGroupIDsG FileTypeResponseFileInfoGroupIDs = "G"
	FileTypeResponseFileInfoGroupIDsI FileTypeResponseFileInfoGroupIDs = "I"
	FileTypeResponseFileInfoGroupIDsM FileTypeResponseFileInfoGroupIDs = "M"
	FileTypeResponseFileInfoGroupIDsP FileTypeResponseFileInfoGroupIDs = "P"
	FileTypeResponseFileInfoGroupIDsT FileTypeResponseFileInfoGroupIDs = "T"
	FileTypeResponseFileInfoGroupIDsZ FileTypeResponseFileInfoGroupIDs = "Z"
	FileTypeResponseFileInfoGroupIDsO FileTypeResponseFileInfoGroupIDs = "O"
)

func (e FileTypeResponseFileInfoGroupIDs) ToPointer() *FileTypeResponseFileInfoGroupIDs {
	return &e
}

func (e *FileTypeResponseFileInfoGroupIDs) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "A":
		fallthrough
	case "D":
		fallthrough
	case "E":
		fallthrough
	case "G":
		fallthrough
	case "I":
		fallthrough
	case "M":
		fallthrough
	case "P":
		fallthrough
	case "T":
		fallthrough
	case "Z":
		fallthrough
	case "O":
		*e = FileTypeResponseFileInfoGroupIDs(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FileTypeResponseFileInfoGroupIDs: %v", v)
	}
}

type FileTypeResponseFileInfo struct {
	// File type description
	Description string `json:"description"`
	// File is password-protected or not
	Encrypted bool `json:"encrypted"`
	// File type extension
	Extensions string `json:"extensions"`
	// File type category
	GroupID string `json:"groupID"`
	// File type category.
	//   * `A` - Archive files
	//   * `AP` - Application
	//   * `D` - Document (Microsoft Office)
	//   * `D_ENC` - Encrypted Documents
	//   * `E` - Executables
	//   * `G` - Graphical format (JPG, PNG, GIF, etc.)
	//   * `I` - Disk image
	//   * `M` - Audio or video format
	//   * `OPENSSL_ENC` - OpenSSL Encrypted Files
	//   * `P` - PDF format
	//   * `T` - Text
	//   * `Z` - Mail messages
	//   * `O` - Other (anything that is not recognized as one of the above)
	//
	GroupIDs []FileTypeResponseFileInfoGroupIDs `json:"groupIDs"`
	// MIME type
	Type string `json:"type"`
	// File type ID
	TypeID string `json:"typeID"`
	// A list of file type IDs
	TypeIds []string `json:"type_ids"`
}

// FileTypeResponseFileInfoDetails - Detailed information.
type FileTypeResponseFileInfoDetails struct {
}

// FileTypeResponseFileMetadata - Metadata information.
type FileTypeResponseFileMetadata struct {
}

// FileTypeResponse - response information from FileType engine
type FileTypeResponse struct {
	FileInfo FileTypeResponseFileInfo `json:"file_info"`
	// Detailed information.
	FileInfoDetails *FileTypeResponseFileInfoDetails `json:"file_info_details,omitempty"`
	// Metadata information.
	FileMetadata *FileTypeResponseFileMetadata `json:"file_metadata,omitempty"`
	// SHA256 Hash of user-interface template. For web console only.
	ResultTemplateHash *string `json:"result_template_hash,omitempty"`
}
