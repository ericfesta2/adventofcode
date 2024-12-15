// Run as `kotlin part_1.kts`

import kotlin.math.pow

private val input = "2 77706 5847 9258441 0 741 883933 12".split(" ").map { it.toLong() }

private val NUM_BLINKS = 25
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

private fun List<Long>.blinkTimes(times: Int) =
    (0..<times).fold(this) { list, _ ->
        list.flatMap { changeStone(it) }
    }

// Part 1 - works well enough for 25 blinks but won't for 75
input.blinkTimes(NUM_BLINKS).let { println(it.size) }
