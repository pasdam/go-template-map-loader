package tm

func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	if len(maps) < 1 {
		return make(map[string]interface{}, 0)
	}
	out := make(map[string]interface{}, len(maps[0]))
	for k, v := range maps[0] {
		out[k] = v
	}
	for i := 1; i < len(maps); i++ {
		for k, v := range maps[i] {
			if v, ok := v.(map[string]interface{}); ok {
				if bv, ok := out[k]; ok {
					if bv, ok := bv.(map[string]interface{}); ok {
						out[k] = MergeMaps(bv, v)
						continue
					}
				}
			}
			out[k] = v
		}
	}
	return out
}
