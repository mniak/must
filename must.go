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

func Must6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, err error) (T1, T2, T3, T4, T5, T6) {
	fatal(err)
	return v1, v2, v3, v4, v5, v6
}

func Must7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, err error) (T1, T2, T3, T4, T5, T6, T7) {
	fatal(err)
	return v1, v2, v3, v4, v5, v6, v7
}

func Must8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, err error) (T1, T2, T3, T4, T5, T6, T7, T8) {
	fatal(err)
	return v1, v2, v3, v4, v5, v6, v7, v8
}

func Must9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, err error) (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	fatal(err)
	return v1, v2, v3, v4, v5, v6, v7, v8, v9
}

func Must10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, err error) (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) {
	fatal(err)
	return v1, v2, v3, v4, v5, v6, v7, v8, v9, v10
}
