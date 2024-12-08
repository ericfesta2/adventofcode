import * as fs from 'fs';
import * as readline from 'readline';

export const ALL_SAFE = -1;

const SAFE_ADJ_DIFF_ABS_MIN = 1;
const SAFE_ADJ_DIFF_ABS_MAX = 3;

export const getInputStream = (): readline.Interface =>
    readline.createInterface({
        input: fs.createReadStream('input.txt')
    });

export const getLevelsList = (reportLine: string): number[] =>
    reportLine.split(' ').map((level: string): number => parseInt(level));

// Returns the first index of an unsafe value candidate or -1 if this report is safe
// Numeric return instead of boolean needed for part 2
export const getFirstUnsafe = (levels: number[]): number => {
    const isIncreasing = levels[0] < levels[levels.length - 1];

    for (let i = 0; i < levels.length - 1; i++) {
        const adjacentDiff = levels[i + 1] - levels[i]; // Positive when increasing, negative when decreasing

        if (!(
            (isIncreasing && adjacentDiff >= SAFE_ADJ_DIFF_ABS_MIN && adjacentDiff <= SAFE_ADJ_DIFF_ABS_MAX) ||
            (!isIncreasing && adjacentDiff >= -SAFE_ADJ_DIFF_ABS_MAX && adjacentDiff <= -SAFE_ADJ_DIFF_ABS_MIN)
        )) {
            // No need to check the rest of this level
            return i;
        }
    }

    return ALL_SAFE;
};
