// Run with `node --experimental-transform-types`
import { getInputStream, getFirstUnsafe, ALL_SAFE, getLevelsList } from './utils.ts';

let safeLevels = 0;

for await (const report of getInputStream()) {
    if (getFirstUnsafe(getLevelsList(report)) === ALL_SAFE) {
        safeLevels += 1;
    }
}

console.log(safeLevels);
