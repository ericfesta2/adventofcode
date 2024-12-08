// Run with `node --experimental-transform-types`
import { ALL_SAFE, getFirstUnsafe, getInputStream, getLevelsList } from './utils.ts';

declare global {
    interface Array<T> {
        removeIndexImmutably<T>(this: Array<T>, indexToRemove: number): T[];
    }
}

Array.prototype.removeIndexImmutably = function <T>(indexToRemove: number): Array<T> {
    return this.filter((_, ind: number): boolean => ind !== indexToRemove);
};

// Keep track of the recursive call depth so that we only check subsets of this report with one level removed
const applyDampener = (levels: number[], callDepth: number = 0): boolean => {
    const firstUnsafe = getFirstUnsafe(levels);

    if (firstUnsafe === ALL_SAFE) {
        return true;
    }

    // If an unsafe index does exist, try the safety check again with 
    return callDepth === 0 ? (
        applyDampener(levels.removeIndexImmutably(firstUnsafe), callDepth + 1) ||
        applyDampener(levels.removeIndexImmutably(firstUnsafe + 1), callDepth + 1)
     ) : false;
};

let safeLevels = 0;

for await (const report of getInputStream()) {
    if (applyDampener(getLevelsList(report))) {
        safeLevels += 1;
    }
}

console.log(safeLevels);
