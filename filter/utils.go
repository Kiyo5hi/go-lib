package filter

func (v *Value) Primitive() any {
	if v.String != nil {
		return string(*v.String)
	}
	if v.Int != nil {
		return int64(*v.Int)
	}
	if v.Float != nil {
		return float64(*v.Float)
	}
	if v.Boolean != nil {
		return bool(*v.Boolean)
	}
	return nil
}
