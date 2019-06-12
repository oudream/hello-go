package utils1

type mainCallBack func(v int) (r int)

//var fCbs []mainCallBack = make([]mainCallBack, 0, 100)
var fCbs = make([]mainCallBack, 0, 100)

func RegCallBack(cb mainCallBack) (r int) {
	if len(fCbs) >= 100 {
		return
	}
	fCbs = append(fCbs, cb)
	r = len(fCbs)
	return
}

func DoCallBack(v int) {
	for _,cb := range fCbs {
		cb(v)
	}
}
