// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codegurureviewer

type JobState string

// Enum values for JobState
const (
	JobStateCompleted JobState = "Completed"
	JobStatePending   JobState = "Pending"
	JobStateFailed    JobState = "Failed"
	JobStateDeleting  JobState = "Deleting"
)

func (enum JobState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum JobState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProviderType string

// Enum values for ProviderType
const (
	ProviderTypeCodeCommit             ProviderType = "CodeCommit"
	ProviderTypeGitHub                 ProviderType = "GitHub"
	ProviderTypeBitbucket              ProviderType = "Bitbucket"
	ProviderTypeGitHubEnterpriseServer ProviderType = "GitHubEnterpriseServer"
)

func (enum ProviderType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProviderType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Reaction string

// Enum values for Reaction
const (
	ReactionThumbsUp   Reaction = "ThumbsUp"
	ReactionThumbsDown Reaction = "ThumbsDown"
)

func (enum Reaction) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Reaction) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RepositoryAssociationState string

// Enum values for RepositoryAssociationState
const (
	RepositoryAssociationStateAssociated     RepositoryAssociationState = "Associated"
	RepositoryAssociationStateAssociating    RepositoryAssociationState = "Associating"
	RepositoryAssociationStateFailed         RepositoryAssociationState = "Failed"
	RepositoryAssociationStateDisassociating RepositoryAssociationState = "Disassociating"
)

func (enum RepositoryAssociationState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RepositoryAssociationState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Type string

// Enum values for Type
const (
	TypePullRequest Type = "PullRequest"
)

func (enum Type) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Type) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}