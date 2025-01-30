// Run as `kotlin part_1.kts`

import kotlin.math.pow

private val input = "2 77706 5847 9258441 0 741 883933 12".split(" ").map { it.toLong() }

private val NUM_BLINKS = 75
private val DEFAULT_VALUE_MULTIPLIER = 2024

private tailrec fun numDigits(value: Long, count: Int = 0): Int =
    if (value == 0L) {
        count
    } else {
        numDigits(value / 10, count + 1)
    }

private fun changeStone(stoneValue: Long): List<Long> =
    if (stoneValue == 0L) {
        listOf(1)
    } else {
        val digitCount = numDigits(stoneValue)

        if (digitCount % 2 == 0) {
            // Mod by this value to get the right half, divide to get the left half
            val pivot = 10.0.pow(digitCount / 2).toInt()

            listOf(stoneValue / pivot, stoneValue % pivot)
        } else {
            listOf(stoneValue * DEFAULT_VALUE_MULTIPLIER)
        }
    }

@Deprecated("""
    The below method worked for part 1 but ran out of memory for part 2 since it did not cache repeated calculations
    (e.g., two stones in the list with the same value)
""")
private fun List<Long>.blinkTimes(times: Int) =
    (0..<times).fold(this) { list, _ ->
        list.flatMap { changeStone(it) }
    }

// To solve part 2, for each blink, count the frequencies of each distinct stone value in the list, then compute
// each distinct new value once; the total number of stones is summed only at the end of the computation.
private tailrec fun numStones(stoneValueFreqs: Map<Long, Long>, times: Int): Long =
    if (times == 0) {
        stoneValueFreqs.values.sum()
    } else {
        val newStoneFreqs = mutableMapOf<Long, Long>().withDefault { 0L }

        stoneValueFreqs.forEach { (origStone, origStoneFreq) ->
            changeStone(origStone).forEach { newStone ->
                newStoneFreqs.put(newStone, newStoneFreqs.getOrDefault(newStone, 0) + origStoneFreq)
            }
        }

        numStones(newStoneFreqs, times - 1)
    }

println(numStones(input.groupingBy { it }.eachCount().mapValues { it.value.toLong() }, NUM_BLINKS))
