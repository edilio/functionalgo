package funct

type Formula func (n int) int
type Recursive func(recur Formula, n int) int


func Memoizer(memo map[int]int, formula Recursive) Formula {
	var recur Formula
    recur = func (n int) int {
		result, ok := memo[n]
        if !ok {
            result = formula(recur, n);
			memo[n] = result;
        }
        return result;
    };
    return recur;
};

func Fib (recur Formula, n int) int {
	return recur(n - 1) + recur(n - 2)
}

func Fact(recur Formula, n int) int {
	return n * recur( n-1 )
}

