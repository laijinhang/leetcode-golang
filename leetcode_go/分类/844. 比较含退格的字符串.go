package main

func backspaceCompare(S string, T string) bool {
	bs, ts := []byte{}, []byte{}
	for i := 0;i < len(S);i++ {
		if S[i] == '#' {
			if len(bs) != 0 {
				bs = bs[:len(bs)-1]
			}
		} else {
			bs = append(bs, S[i])
		}
	}
	for i := 0;i < len(T);i++ {
		if T[i] == '#' {
			if len(ts) != 0 {
				ts = ts[:len(ts)-1]
			}
		} else {
			ts = append(ts, T[i])
		}
	}
	if string(bs) != string(ts) {
		return false
	}
	return true
}