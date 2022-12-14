{{/*
Expand the name of the chart.
*/}}
{{- define "banklo-authorizer-adapter.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "banklo-authorizer-adapter.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "banklo-authorizer-adapter.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "banklo-authorizer-adapter.labels" -}}
helm.sh/chart: {{ include "banklo-authorizer-adapter.chart" . }}
{{ include "banklo-authorizer-adapter.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "banklo-authorizer-adapter.selectorLabels" -}}
app.kubernetes.io/name: {{ include "banklo-authorizer-adapter.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app: banklo-authorizer-adapter
version: v1
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "banklo-authorizer-adapter.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "banklo-authorizer-adapter.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}
