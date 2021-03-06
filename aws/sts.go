/*
Copyright © 2021 Cloud Posse, LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	common "github.com/cloudposse/turf/common/error"
)

func getStsClient(sess *session.Session) *sts.STS {
	return sts.New(sess)
}

func getStsClientWithCreds(sess *session.Session, creds *credentials.Credentials) *sts.STS {
	return sts.New(sess, &aws.Config{Credentials: creds})
}

// GetSession return a new AWS Session
func GetSession() *session.Session {
	session := session.Must(session.NewSession())
	return session
}

// GetCreds return credentials that can be used on a session
func GetCreds(sess *session.Session, role string) *credentials.Credentials {
	creds := stscreds.NewCredentials(sess, role)
	return creds
}

// GetAccountID returns the AWS Account ID of the session
func GetAccountID(sess *session.Session) string {
	client := getStsClient(sess)

	input := sts.GetCallerIdentityInput{}
	ident, err := client.GetCallerIdentity(&input)

	common.AssertErrorNil(err)
	return *ident.Account
}

// GetAccountIDWithRole returns the AWS Account ID of the session after assuming a role
func GetAccountIDWithRole(sess *session.Session, role string) string {
	creds := GetCreds(sess, role)
	client := getStsClientWithCreds(sess, creds)

	input := sts.GetCallerIdentityInput{}
	ident, err := client.GetCallerIdentity(&input)

	common.AssertErrorNil(err)
	return *ident.Account
}
