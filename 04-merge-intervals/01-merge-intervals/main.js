function merge(intervals) {
    if (intervals.length < 2) {
        return intervals
    }

    intervals.sort((a, b) => a[0] - b[0]);
    let result = [intervals[0]];

    for (let i = 1; i < intervals.length; i++) {
        let lastInterval = result[result.length - 1];
        let [curStart, curEnd] = intervals[i];

        if (lastInterval[1] < curStart) {
            result.push(intervals[i]);
        } else {
            if (lastInterval[1] < curEnd) {
                lastInterval[1] = curEnd;
            }
        }
    }

    console.log(result);
}

merge([[1,4], [2,5], [7,9]])//[[1,5], [7,9]], Since the first two intervals [1,4] and [2,5] overlap, we merged them into one [1,5].
merge([[6,7], [2,4], [5,9]])//[[2,4], [5,9]], Since the intervals [6,7] and [5,9] overlap, we merged them into one [5,9].
merge([[1,4], [2,6], [3,5]])//[[1,6]], Since all the given intervals overlap, we merged them into one.
merge([[2,5]])