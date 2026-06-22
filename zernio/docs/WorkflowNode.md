# WorkflowNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Stable node id referenced by edges | 
**Type** | **string** | Node kind. The 16 supported types break into four groups:   messaging (send_message),   control flow (trigger, condition, delay, wait_for_reply, a_b_split, end),   data ops (set_variable, set_field, add_tag, remove_tag, enroll_sequence),   integrations (webhook, ai, handoff, start_call).  | 
**Config** | Pointer to **map[string]interface{}** | Type-specific settings. All string fields support &#x60;{{variable}}&#x60; interpolation against the run&#39;s variable bag (resolved at execution time).  **trigger**: &#x60;{ triggerType: inbound_message|api_call|whatsapp_event, keywords:[string], matchType: any|contains|exact|regex, onlyFirstMessage:boolean, eventType: message_sent|message_delivered|message_read|message_failed|reaction }&#x60;. Default &#x60;triggerType&#x60; is &#x60;inbound_message&#x60; for legacy nodes. &#x60;eventType&#x60; is only honored when &#x60;triggerType&#x60; is &#x60;whatsapp_event&#x60; (WhatsApp-only).  **send_message**: &#x60;{ messageType: text|template|media|interactive, text, template:{name,language,variableMapping}, media:{mediaType:image|video|audio|document, url,caption}, interactive }&#x60;. &#x60;template&#x60; and &#x60;interactive&#x60; are WhatsApp-only.  **wait_for_reply**: &#x60;{ timeoutMinutes:int (max 43200), saveAs:string }&#x60;. Resume via the &#x60;&#39;reply&#39;&#x60; edge on inbound, or &#x60;&#39;timeout&#39;&#x60; edge after &#x60;timeoutMinutes&#x60; of silence.  **condition**: &#x60;{ rules:[{ id, variable, operator: equals|not_equals|contains|not_contains|starts_with|ends_with|exists|not_exists|matches, value }] }&#x60;. First matching rule takes its &#x60;id&#x60; as the sourceHandle; otherwise &#x60;&#39;default&#39;&#x60;.  **set_variable**: &#x60;{ assignments:[{ name, value }] }&#x60;. Run-scoped (lives only for this execution; use &#x60;set_field&#x60; for persistent values).  **delay**: &#x60;{ delayMinutes:int (max 43200) }&#x60;. Suspends the run, resumes via timer.  **webhook**: &#x60;{ url, method: GET|POST|PUT|PATCH|DELETE, headers, bodyTemplate, saveAs }&#x60;. SSRF-guarded (private/loopback/metadata IPs rejected). Response saved as &#x60;{ status, ok, body }&#x60; to &#x60;vars[saveAs]&#x60;. Edge: &#x60;&#39;success&#39;&#x60; on 2xx, &#x60;&#39;error&#39;&#x60; otherwise.  **ai**: &#x60;{ provider: anthropic|openai|google|mistral|groq, model, preset: smart|tools|cheap, systemPrompt, userPromptTemplate, saveAs, temperature, maxTokens, outputType: text|json, tools:[{ name, description, parameters }] }&#x60;. Set &#x60;provider&#x60; + &#x60;model&#x60; for BYOK (uses your stored API key); omit &#x60;provider&#x60; for the legacy Telnyx path. Edges: &#x60;&#39;success&#39;&#x60;, &#x60;&#39;tool:&lt;name&gt;&#39;&#x60; (model picked a tool), &#x60;&#39;error&#39;&#x60;.  **handoff**: &#x60;{ note, assignTo }&#x60;. Terminates the run as &#x60;exited&#x60;, flags the conversation for a human operator.  **start_call**: &#x60;{ to, forwardTo, requirePermissionFirst, recordingEnabled, saveAs }&#x60;. WhatsApp-only. &#x60;forwardTo&#x60; can be &#x60;tel:+E164&#x60;, &#x60;sip:user@host&#x60;, or &#x60;wss://…&#x60; (AI voice agent). Edges: &#x60;&#39;success&#39;&#x60;, &#x60;&#39;permission_required&#39;&#x60;, &#x60;&#39;failed&#39;&#x60;.  **a_b_split**: &#x60;{ percentage: number 0-100 (default 50) }&#x60;. Random branch picker. Edges: &#x60;&#39;a&#39;&#x60; (with probability &#x60;percentage/100&#x60;), &#x60;&#39;b&#39;&#x60;.  **set_field**: &#x60;{ field, value }&#x60;. Persistent custom field on the Contact (vs &#x60;set_variable&#x60; which is run-scoped). Field name is sanitized to &#x60;[A-Za-z0-9_]&#x60;. No-op on &#x60;api_call&#x60; runs (no contact).  **enroll_sequence**: &#x60;{ sequenceId, saveAs }&#x60;. Enrolls the run&#39;s contact into a Sequence. Edges: &#x60;&#39;success&#39;&#x60;, &#x60;&#39;error&#39;&#x60;.  **add_tag** / **remove_tag**: &#x60;{ tag }&#x60;. Push or pull a tag on the Contact. No-op on &#x60;api_call&#x60; runs.  **end**: no config. Terminates the run as &#x60;completed&#x60;.  | [optional] 
**Position** | Pointer to [**WorkflowNodePosition**](WorkflowNodePosition.md) |  | [optional] 

## Methods

### NewWorkflowNode

`func NewWorkflowNode(id string, type_ string, ) *WorkflowNode`

NewWorkflowNode instantiates a new WorkflowNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowNodeWithDefaults

`func NewWorkflowNodeWithDefaults() *WorkflowNode`

NewWorkflowNodeWithDefaults instantiates a new WorkflowNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowNode) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *WorkflowNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowNode) SetType(v string)`

SetType sets Type field to given value.


### GetConfig

`func (o *WorkflowNode) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *WorkflowNode) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *WorkflowNode) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *WorkflowNode) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetPosition

`func (o *WorkflowNode) GetPosition() WorkflowNodePosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *WorkflowNode) GetPositionOk() (*WorkflowNodePosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *WorkflowNode) SetPosition(v WorkflowNodePosition)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *WorkflowNode) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


