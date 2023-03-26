/**
 * @param {number[][]} trips
 * @param {number} capacity
 * @return {boolean}
 */
var carPooling = function(trips, capacity) {
    trips.sort((a, b) => a[0] - b[0])
    let maxCapacity = 0;

    for (let i = 1; i < trips.length; i++) {
        let previous = trips[i - 1];
        let current = trips[i];

        if (current[1] < previous[2]) {
            trips[i] = [current[0] + previous[0], previous[1], current[2]];
            trips.splice(i - 1, 1)
            i--
        }
    }

    //set maximum load
    for(let i = 0; i < trips.length; i++) {
        maxCapacity = Math.max(maxCapacity, trips[i][0]) 
    }

    return maxCapacity <= capacity
};

carPooling([[2,2,6],[2,4,7],[8,6,7]], 11)//7