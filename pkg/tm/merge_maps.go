package tm

func MergeMaps[K comparable](maps ...map[K]interface{}) map[K]interface{} {
	if len(maps) < 1 {
		return make(map[K]interface{})
	}
	out := make(map[K]interface{}, len(maps[0]))
	for k, v := range maps[0] {
		out[k] = v
	}
	for i := 1; i < len(maps); i++ {
		for k, v := range maps[i] {
			if vs, ok := v.(map[string]interface{}); ok {
				if bv, ok := out[k]; ok {
					if bv, ok := bv.(map[string]interface{}); ok {
						out[k] = MergeMaps(bv, vs)
						continue
					}
				}
			} else if vi, ok := v.(map[interface{}]interface{}); ok {
				if bv, ok := out[k]; ok {
					if bv, ok := bv.(map[interface{}]interface{}); ok {
						out[k] = MergeMaps(bv, vi)
						continue
					}
				}
			} else if vi, ok := v.(map[int]interface{}); ok {
				if bv, ok := out[k]; ok {
					if bv, ok := bv.(map[int]interface{}); ok {
						out[k] = MergeMaps(bv, vi)
						continue
					}
				}
			}
			out[k] = v
		}
	}
	return out
}
