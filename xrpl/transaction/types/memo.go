package types

// MemoWrapper wraps a Memo for optional inclusion in a transaction.
type MemoWrapper struct {
	Memo Memo
}

// Memo represents a transaction memo, including optional data, format, and type fields.
type Memo struct {
	MemoData   string `json:",omitempty"`
	MemoFormat string `json:",omitempty"`
	MemoType   string `json:",omitempty"`
}

// Flatten returns a map containing the Memo if it is set, or nil otherwise.
func (mw *MemoWrapper) Flatten() map[string]interface{} {
	if mw.Memo != (Memo{}) {
		flattened := make(map[string]interface{})
		flattened["Memo"] = mw.Memo.Flatten()
		return flattened
	}
	return nil
}

// Flatten returns a map of the Memo fields that are non-empty.
func (m *Memo) Flatten() map[string]interface{} {
	flattened := make(map[string]interface{})

	if m.MemoData != "" {
		flattened["MemoData"] = m.MemoData
	}

	if m.MemoFormat != "" {
		flattened["MemoFormat"] = m.MemoFormat
	}

	if m.MemoType != "" {
		flattened["MemoType"] = m.MemoType
	}

	return flattened
}
