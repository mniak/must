package must

func Must0(err error) {
	fatal(err)
}

func Must1[T1 any](v1 T1, err error) T1 {
	fatal(err)
	return v1
}

func Must2[T1 any, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	fatal(err)
	return v1, v2
}

func Must3[T1 any, T2 any, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	fatal(err)
	return v1, v2, v3
}

func Must4[T1 any, T2 any, T3 any, T4 any](v1 T1, v2 T2, v3 T3, v4 T4, err error) (T1, T2, T3, T4) {
	fatal(err)
	return v1, v2, v3, v4
}

func Must5[T1 any, T2 any, T3 any, T4 any, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, err error) (T1, T2, T3, T4, T5) {
	fatal(err)
	return v1, v2, v3, v4, v5
}
