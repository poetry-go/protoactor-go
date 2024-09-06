package actor

type messageHeader map[string]interface{}

func (header messageHeader) Get(key string) string {
	str, ok := header[key].(string)
	if !ok {
		return ""
	}

	return str
}

func (header messageHeader) Set(key string, value interface{}) {
	header[key] = value
}

func (header messageHeader) Keys() []string {
	keys := make([]string, 0, len(header))
	for k := range header {
		keys = append(keys, k)
	}
	return keys
}

func (header messageHeader) Length() int {
	return len(header)
}

func (header messageHeader) ToMap() map[string]string {
	mp := make(map[string]string)
	for k, v := range header {
		str, ok := v.(string)
		if !ok {
			continue
		}

		mp[k] = str
	}

	return mp
}

type ReadonlyMessageHeader interface {
	Get(key string) string
	Keys() []string
	Length() int
	ToMap() map[string]string
}

type MessageEnvelope struct {
	Header  messageHeader
	Message interface{}
	Sender  *PID
}

func (envelope *MessageEnvelope) GetHeader(key string) interface{} {
	if envelope.Header == nil {
		return ""
	}
	return envelope.Header.Get(key)
}

func (envelope *MessageEnvelope) SetHeader(key string, value interface{}) {
	if envelope.Header == nil {
		envelope.Header = make(map[string]interface{})
	}
	envelope.Header.Set(key, value)
}

var EmptyMessageHeader = make(messageHeader)

func WrapEnvelope(message interface{}) *MessageEnvelope {
	if e, ok := message.(*MessageEnvelope); ok {
		return e
	}
	return &MessageEnvelope{nil, message, nil}
}

func UnwrapEnvelope(message interface{}) (ReadonlyMessageHeader, interface{}, *PID) {
	if env, ok := message.(*MessageEnvelope); ok {
		return env.Header, env.Message, env.Sender
	}
	return nil, message, nil
}

func UnwrapEnvelopeHeader(message interface{}) ReadonlyMessageHeader {
	if env, ok := message.(*MessageEnvelope); ok {
		return env.Header
	}
	return nil
}

func UnwrapEnvelopeMessage(message interface{}) interface{} {
	if env, ok := message.(*MessageEnvelope); ok {
		return env.Message
	}
	return message
}

func UnwrapEnvelopeSender(message interface{}) *PID {
	if env, ok := message.(*MessageEnvelope); ok {
		return env.Sender
	}
	return nil
}

// ConvertToMessageHeader converts a map[string]string to a messageHeader
func ConvertToMessageHeader(header map[string]string) map[string]interface{} {
	headers := make(messageHeader, len(header))
	for k, v := range header {
		headers[k] = v
	}

	return headers
}
