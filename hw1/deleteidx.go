package main

//第一次作业

func Yaoqiu1(tst []any, idx int) ([]any, string) {
	if idx < 0 || len(tst) <= idx {
		return tst, "下标溢出"
	}
	res := tst[:idx]
	for i := idx + 1; i < len(tst); i++ {
		res = append(res, tst[i])
	}
	return res, "ok"
}

// 不使用新的切片，实现原地删除
func Yaoqiu2(tst []any, idx int) ([]any, string) {
	if idx < 0 || len(tst) <= idx {
		return tst, "下标溢出"
	}
	for i := idx + 1; i < len(tst); i++ {
		tst[i-1] = tst[i]
	}
	tst[len(tst)-1] = ""
	return tst, "ok"
}

func Yaoqiu3[T Number](tst []T, idx int) []T {
	if idx < 0 || len(tst) <= idx {
		panic("idx不合法")
	}
	res := tst[:idx]
	for i := idx + 1; i < len(tst); i++ {
		res = append(res, tst[i])
	}
	return res
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | string
}

func Yaoqiu4[T Number](tst []T, idx int) []T {
	if idx < 0 || len(tst) <= idx {
		panic("idx不合法")
	}
	res := make([]T, 0, len(tst)-1)
	for i := 0; i < idx; i++ {
		res = append(res, tst[i])
	}
	for i := idx + 1; i < len(tst); i++ {
		res = append(res, tst[i])
	}
	return res
}
