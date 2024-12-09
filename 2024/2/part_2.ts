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

const applyDampener = (levels: number[]): boolean => {
    const firstUnsafe = getFirstUnsafe(levels);

    if (firstUnsafe === ALL_SAFE) {
        return true;
    }

    // If an unsafe index does exist, try the safety check again with (a) that index removed and
    // (b) the next index removed to see if either can be safe
    return (
        getFirstUnsafe(levels.removeIndexImmutably(firstUnsafe)) === ALL_SAFE ||
        getFirstUnsafe(levels.removeIndexImmutably(firstUnsafe + 1)) === ALL_SAFE
    );
};

let safeLevels = 0;

for await (const report of getInputStream()) {
    if (applyDampener(getLevelsList(report))) {
        safeLevels += 1;
    }
}

console.log(safeLevels);
