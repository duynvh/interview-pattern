// function findPermutations(nums) {
//     const result = [];
//     const permutations = [[]];
//     const length = nums.length;

//     for (let i = 0; i < length; i++) {
//         const current = nums[i];

//         const n = permutations.length;

//         for (let p = 0; p < n; p++) {
//             const oldPermutations = permutations.shift();

//             for (let j = 0; j < oldPermutations.length + 1; j++) {
//                 //clone the permutation
//                 const newPermutations = oldPermutations.slice(0);

//                 newPermutations.splice(j, 0, current);

//                 if (newPermutations.length === length) {
//                     result.push(newPermutations);
//                 } else {
//                     permutations.push(newPermutations);
//                 }
//             }
//         }
//     }

//     return result;
// }

function findPermutations(nums) {
    let subsets = [];

    generatePermutationRecursive(nums, 0, [], subsets);

    return subsets;
}

function generatePermutationRecursive(nums, index, currentPermutation, subsets) {
    if (index === nums.length) {
        subsets.push(currentPermutation);
    } else {
        for (let i = 0; i < currentPermutation.length + 1; i++) {
            let newPermutation = currentPermutation.slice(0);

            newPermutation.splice(i, 0, nums[index]);

            generatePermutationRecursive(nums, index + 1, newPermutation, subsets);
        }
    }
}

var letterCasePermutation = function(s) {
    const permutations = [s];

    for (let i = 0; i < s.length; i++) {
        if (isNaN(s[i], 10)) {
            const n = permutations.length;

            for (let j = 0; j < n; j++) {
                const chs = permutations[j].split('');

                if (chs[i] === chs[i].toLowerCase()) {
                    chs[i] = chs[i].toUpperCase();
                } else {
                    chs[i] = chs[i].toLowerCase();
                }

                permutations.push(chs.join(''));
            }
        }
    }

    return permutations;
};

console.log(letterCasePermutation('ab7c'));