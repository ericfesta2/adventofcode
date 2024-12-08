// Run with `node --experimental-transform-types`

import * as fs from 'fs';
import * as readline from 'readline';

const inp = fs.createReadStream('input.txt');

const lineReader = readline.createInterface({
    input: inp
});

const SAFE_ADJ_DIFF_ABS_MIN = 1;
const SAFE_ADJ_DIFF_ABS_MAX = 3;

let safeLevels = 0;

for await (const report of lineReader) {
    const levels = report.split(' ').map((level: string): number => parseInt(level));

    const isIncreasing = levels[0] < levels[levels.length - 1];
    let isReportSafe = true;

    for (let i = 0; i < levels.length - 1; i++) {
        const adjacentDiff = levels[i + 1] - levels[i]; // Positive when increasing, negative when decreasing

        if (!(
            (isIncreasing && adjacentDiff >= SAFE_ADJ_DIFF_ABS_MIN && adjacentDiff <= SAFE_ADJ_DIFF_ABS_MAX) ||
            (!isIncreasing && adjacentDiff >= -SAFE_ADJ_DIFF_ABS_MAX && adjacentDiff <= -SAFE_ADJ_DIFF_ABS_MIN)
        )) {
            // No need to check the rest of this level
            isReportSafe = false;
            break;
        }
    }

    if (isReportSafe) {
        safeLevels += 1;
    }
}

console.log(safeLevels);
