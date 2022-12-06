package must

func Do0(err error) {
	if err != nil {
		fatal(err)
	}
}

func Do1[T1 any](v1 T1, err error) T1 {
	if err != nil {
		fatal(err)
	}
	return v1
}

func Do2[T1 any, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	if err != nil {
		fatal(err)
	}
	return v1, v2
}

func Do3[T1 any, T2 any, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	if err != nil {
		fatal(err)
	}
	return v1, v2, v3
}

func Do4[T1 any, T2 any, T3 any, T4 any](v1 T1, v2 T2, v3 T3, v4 T4, err error) (T1, T2, T3, T4) {
	if err != nil {
		fatal(err)
	}
	return v1, v2, v3, v4
}

func Do5[T1 any, T2 any, T3 any, T4 any, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, err error) (T1, T2, T3, T4, T5) {
	if err != nil {
		fatal(err)
	}
	return v1, v2, v3, v4, v5
}

func Do6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, err error) (T1, T2, T3, T4, T5, T6) {
	if err != nil {
		fatal(err)
	}
	return v1, v2, v3, v4, v5, v6
}

func Do7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, err error) (T1, T2, T3, T4, T5, T6, T7) {
	if err != nil {
		fatal(err)
	}
	return v1, v2, v3, v4, v5, v6, v7
}

func Do8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, err error) (T1, T2, T3, T4, T5, T6, T7, T8) {
	if err != nil {
		fatal(err)
	}
	return v1, v2, v3, v4, v5, v6, v7, v8
}

func Do9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, err error) (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	if err != nil {
		fatal(err)
	}
	return v1, v2, v3, v4, v5, v6, v7, v8, v9
}

func Do10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, err error) (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) {
	if err != nil {
		fatal(err)
	}
	return v1, v2, v3, v4, v5, v6, v7, v8, v9, v10
}
