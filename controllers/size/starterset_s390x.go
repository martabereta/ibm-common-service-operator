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

package size

const StarterSet = `
- name: ibm-cert-manager-operator
  spec:
    certManager:
      certManagerCAInjector:
        resources:
          limits:
            cpu: 100m
            memory: 520Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 30m
            ephemeral-storage: 256Mi
            memory: 350Mi
      certManagerController:
        resources:
          limits:
            cpu: 110m
            memory: 530Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 70m
            ephemeral-storage: 256Mi
            memory: 390Mi
      certManagerWebhook:
        resources:
          limits:
            cpu: 60m
            memory: 100Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 50m
            memory: 90Mi
- name: ibm-mongodb-operator
  spec:
    mongoDB:
      replicas: 1
      resources:
        limits:
          cpu: 1000m
          memory: 700Mi
        requests:
          ephemeral-storage: 256Mi
          cpu: 500m
          memory: 700Mi
      metrics:
        resources:
          requests:
            ephemeral-storage: 256Mi
            cpu: 100m
            memory: 300Mi
          limits:
            cpu: 1000m
            memory: 350Mi
- name: ibm-im-mongodb-operator
  spec:
    mongoDB:
      replicas: 1
      resources:
        limits:
          cpu: 1000m
          memory: 700Mi
        requests:
          ephemeral-storage: 256Mi
          cpu: 500m
          memory: 700Mi
      metrics:
        resources:
          requests:
            ephemeral-storage: 256Mi
            cpu: 100m
            memory: 300Mi
          limits:
            cpu: 1000m
            memory: 350Mi
- name: ibm-im-mongodb-operator-v4.0
  spec:
    mongoDB:
      replicas: 1
      resources:
        limits:
          cpu: 1000m
          memory: 700Mi
        requests:
          ephemeral-storage: 256Mi
          cpu: 500m
          memory: 700Mi
      metrics:
        resources:
          requests:
            ephemeral-storage: 256Mi
            cpu: 100m
            memory: 300Mi
          limits:
            cpu: 1000m
            memory: 350Mi
- name: ibm-im-mongodb-operator-v4.1
  spec:
    mongoDB:
      replicas: 1
      resources:
        limits:
          cpu: 1000m
          memory: 700Mi
        requests:
          ephemeral-storage: 256Mi
          cpu: 500m
          memory: 700Mi
      metrics:
        resources:
          requests:
            ephemeral-storage: 256Mi
            cpu: 100m
            memory: 300Mi
          limits:
            cpu: 1000m
            memory: 350Mi
- name: ibm-im-mongodb-operator-v4.2
  spec:
    mongoDB:
      replicas: 1
      resources:
        limits:
          cpu: 1000m
          memory: 700Mi
        requests:
          ephemeral-storage: 256Mi
          cpu: 500m
          memory: 700Mi
      metrics:
        resources:
          requests:
            ephemeral-storage: 256Mi
            cpu: 100m
            memory: 300Mi
          limits:
            cpu: 1000m
            memory: 350Mi
- name: ibm-iam-operator
  spec:
    authentication:
      replicas: 1
      auditService:
        resources:
          limits:
            cpu: 1000m
            memory: 50Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 20m
            memory: 50Mi
      authService:
        resources:
          limits:
            cpu: 1000m
            memory: 650Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 140m
            memory: 525Mi
      clientRegistration:
        resources:
          limits:
            cpu: 1000m
            memory: 50Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 20m
            memory: 50Mi
      identityManager:
        resources:
          limits:
            cpu: 1000m
            memory: 220Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 50m
            memory: 120Mi
      identityProvider:
        resources:
          limits:
            cpu: 1000m
            memory: 230Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 80m
            memory: 130Mi
    oidcclientwatcher:
      replicas: 1
      resources:
        limits:
          cpu: 1000m
          memory: 256Mi
        requests:
          ephemeral-storage: 256Mi
          cpu: 30m
          memory: 50Mi
    pap:
      auditService:
        resources:
          limits:
            cpu: 1000m
            memory: 50Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 20m
            memory: 50Mi
      papService:
        resources:
          limits:
            cpu: 1000m
            memory: 330Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 50m
            memory: 160Mi
      replicas: 1
    policycontroller:
      replicas: 1
      resources:
        limits:
          cpu: 1000m
          memory: 50Mi
        requests:
          ephemeral-storage: 256Mi
          cpu: 20m
          memory: 50Mi
    policydecision:
      auditService:
        resources:
          limits:
            cpu: 1000m
            memory: 50Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 20m
            memory: 50Mi
      resources:
        limits:
          cpu: 1000m
          memory: 50Mi
        requests:
          ephemeral-storage: 256Mi
          cpu: 20m
          memory: 50Mi
      replicas: 1
    secretwatcher:
      resources:
        limits:
          cpu: 1000m
          memory: 145Mi
        requests:
          ephemeral-storage: 256Mi
          cpu: 30m
          memory: 120Mi
      replicas: 1
    securityonboarding:
      replicas: 1
      resources:
        limits:
          cpu: 1000m
          memory: 50Mi
        requests:
          ephemeral-storage: 256Mi
          cpu: 20m
          memory: 50Mi
      iamOnboarding:
        resources:
          limits:
            cpu: 1000m
            memory: 1024Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 20m
            memory: 64Mi
- name: ibm-im-operator
  spec:
    authentication:
      replicas: 1
      authService:
        resources:
          limits:
            cpu: 1000m
            memory: 650Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 140m
            memory: 525Mi
      clientRegistration:
        resources:
          limits:
            cpu: 1000m
            memory: 50Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 20m
            memory: 50Mi
      identityManager:
        resources:
          limits:
            cpu: 1000m
            memory: 220Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 50m
            memory: 120Mi
      identityProvider:
        resources:
          limits:
            cpu: 1000m
            memory: 230Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 80m
            memory: 130Mi
- name: ibm-im-operator-v4.0
  spec:
    authentication:
      replicas: 1
      authService:
        resources:
          limits:
            cpu: 1000m
            memory: 650Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 140m
            memory: 525Mi
      clientRegistration:
        resources:
          limits:
            cpu: 1000m
            memory: 50Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 20m
            memory: 50Mi
      identityManager:
        resources:
          limits:
            cpu: 1000m
            memory: 220Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 50m
            memory: 120Mi
      identityProvider:
        resources:
          limits:
            cpu: 1000m
            memory: 230Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 80m
            memory: 130Mi
- name: ibm-im-operator-v4.1
  spec:
    authentication:
      replicas: 1
      authService:
        resources:
          limits:
            cpu: 1000m
            memory: 650Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 140m
            memory: 525Mi
      clientRegistration:
        resources:
          limits:
            cpu: 1000m
            memory: 50Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 20m
            memory: 50Mi
      identityManager:
        resources:
          limits:
            cpu: 1000m
            memory: 220Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 50m
            memory: 120Mi
      identityProvider:
        resources:
          limits:
            cpu: 1000m
            memory: 230Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 80m
            memory: 130Mi
- name: ibm-im-operator-v4.2
  spec:
    authentication:
      replicas: 1
      authService:
        resources:
          limits:
            cpu: 1000m
            memory: 650Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 140m
            memory: 525Mi
      clientRegistration:
        resources:
          limits:
            cpu: 1000m
            memory: 50Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 20m
            memory: 50Mi
      identityManager:
        resources:
          limits:
            cpu: 1000m
            memory: 220Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 50m
            memory: 120Mi
      identityProvider:
        resources:
          limits:
            cpu: 1000m
            memory: 230Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 80m
            memory: 130Mi
- name: ibm-management-ingress-operator
  spec:
    managementIngress:
      replicas: 1
      resources:
        requests:
          ephemeral-storage: 256Mi
          cpu: 50m
          memory: 100Mi
        limits:
          cpu: 50m
          memory: 170Mi
- name: ibm-ingress-nginx-operator
  spec:
    nginxIngress:
      ingress:
        replicas: 1
        resources:
          requests:
            ephemeral-storage: 256Mi
            cpu: 100m
            memory: 140Mi
          limits:
            cpu: 100m
            memory: 350Mi
      defaultBackend:
        replicas: 1
        resources:
          requests:
            ephemeral-storage: 256Mi
            cpu: 20m
            memory: 50Mi
          limits:
            cpu: 20m
            memory: 50Mi
      kubectl:
        resources:
          requests:
            ephemeral-storage: 256Mi
            memory: 150Mi
            cpu: 30m
          limits:
            memory: 150Mi
            cpu: 30m
- name: ibm-licensing-operator
  spec:
    IBMLicensing:
      resources:
        requests:
          ephemeral-storage: 256Mi
          cpu: 100m
          memory: 220Mi
        limits:
          cpu: 200m
          memory: 320Mi
    IBMLicenseServiceReporter:
      databaseContainer:
        resources:
          requests:
            ephemeral-storage: 256Mi
            cpu: 200m
            memory: 256Mi
          limits:
            cpu: 300m
            memory: 300Mi
      receiverContainer:
        resources:
          requests:
            ephemeral-storage: 256Mi
            cpu: 200m
            memory: 256Mi
          limits:
            cpu: 300m
            memory: 384Mi
- name: ibm-commonui-operator
  spec:
    commonWebUI:
      replicas: 1
      resources:
        requests:
          ephemeral-storage: 256Mi
          memory: 256Mi
          cpu: 150m
        limits:
          memory: 440Mi
          cpu: 1000m
- name: ibm-idp-config-ui-operator
  spec:
    commonWebUI:
      replicas: 1
      resources:
        requests:
          ephemeral-storage: 256Mi
          memory: 256Mi
          cpu: 150m
        limits:
          memory: 440Mi
          cpu: 1000m
- name: ibm-idp-config-ui-operator-v4.0
  spec:
    commonWebUI:
      replicas: 1
      resources:
        requests:
          ephemeral-storage: 256Mi
          memory: 256Mi
          cpu: 150m
        limits:
          memory: 440Mi
          cpu: 1000m
- name: ibm-idp-config-ui-operator-v4.1
  spec:
    commonWebUI:
      replicas: 1
      resources:
        requests:
          ephemeral-storage: 256Mi
          memory: 256Mi
          cpu: 150m
        limits:
          memory: 440Mi
          cpu: 1000m
- name: ibm-idp-config-ui-operator-v4.2
  spec:
    commonWebUI:
      replicas: 1
      resources:
        requests:
          ephemeral-storage: 256Mi
          memory: 256Mi
          cpu: 150m
        limits:
          memory: 440Mi
          cpu: 1000m
- name: ibm-platform-api-operator
  spec:
    platformApi:
      auditService:
        resources:
          limits:
            cpu: 25m
            memory: 50Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 25m
            memory: 50Mi
      platformApi:
        resources:
          limits:
            cpu: 25m
            memory: 50Mi
          requests:
            ephemeral-storage: 256Mi
            cpu: 25m
            memory: 50Mi
      replicas: 1
- name: ibm-healthcheck-operator
  spec:
    healthService:
      memcached:
        replicas: 1
        resources:
          requests:
            ephemeral-storage: 256Mi
            memory: 50Mi
            cpu: 20m
          limits:
            memory: 100Mi
            cpu: 200m
      healthService:
        replicas: 1
        resources:
          requests:
            ephemeral-storage: 256Mi
            memory: 50Mi
            cpu: 20m
          limits:
            memory: 100Mi
            cpu: 200m
- name: ibm-auditlogging-operator
  spec:
    auditLogging:
      fluentd:
        resources:
          requests:
            ephemeral-storage: 256Mi
            cpu: 25m
            memory: 100Mi
          limits:
            cpu: 35m
            memory: 150Mi
- name: ibm-monitoring-grafana-operator
  spec:
    grafana:
      grafanaConfig:
        resources:
          requests:
            ephemeral-storage: 256Mi
            cpu: 20m
            memory: 65Mi
          limits:
            cpu: 150m
            memory: 100Mi
      dashboardConfig:
        resources:
          requests:
            ephemeral-storage: 256Mi
            cpu: 5m
            memory: 50Mi
          limits:
            cpu: 20m
            memory: 80Mi
      routerConfig:
        resources:
          requests:
            ephemeral-storage: 256Mi
            cpu: 10m
            memory: 50Mi
          limits:
            cpu: 50m
            memory: 50Mi
- name: ibm-apicatalog
  spec:
    apicatalogmanager:
      profile: small
`
