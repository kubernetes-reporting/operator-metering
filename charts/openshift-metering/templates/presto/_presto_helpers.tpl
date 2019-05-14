{{- de***REMOVED***ne "presto-hive-catalog-properties" -}}
connector.name=hive-hadoop2
hive.allow-drop-table=true
hive.allow-rename-table=true
hive.storage-format={{ .Values.presto.spec.hive.con***REMOVED***g.defaultFileFormat | upper }}
hive.compression-codec=SNAPPY
hive.hdfs.authentication.type=NONE
hive.metastore.authentication.type=NONE
hive.metastore.uri={{ .Values.presto.spec.hive.con***REMOVED***g.metastoreURIs }}
{{- if .Values.presto.spec.presto.con***REMOVED***g.metastoreTimeout }}
hive.metastore-timeout={{ .Values.presto.spec.presto.con***REMOVED***g.metastoreTimeout }}
{{- end }}
{{ end }}

{{- de***REMOVED***ne "presto-jmx-catalog-properties" -}}
connector.name=jmx
{{ end }}

{{- de***REMOVED***ne "presto-common-env" }}
- name: MY_NODE_ID
  valueFrom:
    ***REMOVED***eldRef:
      ***REMOVED***eldPath: metadata.uid
- name: MY_NODE_NAME
  valueFrom:
    ***REMOVED***eldRef:
      ***REMOVED***eldPath: spec.nodeName
- name: MY_POD_NAME
  valueFrom:
    ***REMOVED***eldRef:
      ***REMOVED***eldPath: metadata.name
- name: MY_POD_NAMESPACE
  valueFrom:
    ***REMOVED***eldRef:
      ***REMOVED***eldPath: metadata.namespace
- name: MY_MEM_REQUEST
  valueFrom:
    resourceFieldRef:
      containerName: presto
      resource: requests.memory
- name: MY_MEM_LIMIT
  valueFrom:
    resourceFieldRef:
      containerName: presto
      resource: limits.memory
- name: AWS_ACCESS_KEY_ID
  valueFrom:
    secretKeyRef:
      name: "{{ .Values.presto.spec.con***REMOVED***g.awsCredentialsSecretName }}"
      key: aws-access-key-id
      optional: true
- name: AWS_SECRET_ACCESS_KEY
  valueFrom:
    secretKeyRef:
      name: "{{ .Values.presto.spec.con***REMOVED***g.awsCredentialsSecretName }}"
      key: aws-secret-access-key
      optional: true
{{- end }}
