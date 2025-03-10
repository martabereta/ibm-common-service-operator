//
// Copyright 2022 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package constant

const FipsEnabledTemplate = `
- name: ibm-iam-operator
  spec:
    authentication:
      config:
        fipsEnabled: placeholder
- name: ibm-ingress-nginx-operator
  spec:
    nginxIngress:
      fips_enabled: placeholder
- name: ibm-management-ingress-operator
  spec:
    managementIngress:
      fipsEnabled: placeholder
- name: ibm-im-operator
  spec:
    authentication:
      config:
        fipsEnabled: placeholder
- name: ibm-im-operator-v4.0
  spec:
    authentication:
      config:
        fipsEnabled: placeholder
- name: ibm-im-operator-v4.1
  spec:
    authentication:
      config:
        fipsEnabled: placeholder
- name: ibm-im-operator-v4.2
  spec:
    authentication:
      config:
        fipsEnabled: placeholder
`
