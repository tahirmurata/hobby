<script setup lang="ts">
import { computed, ref, watch } from 'vue'

const a = ref("1")
const c = ref("10")
const d = ref("9.5")
const e = ref("10")

const rollOver = ref(false)

// a = betAmount
// b = profitOnWin
// c = rollUnder
// d = multiplier
// e = winChance

// c = e
// d = a * (95 / c) - a

function removeZero(num: number, precision: number = 4): string {
  return (+(+(num)).toFixed(precision) / 1).toString()
}

watch([c, e], ([newC, newE], [oldC, oldE]) => {
  if (parseFloat(newC) < 0.01 || parseFloat(newC) > 100 || parseFloat(newE) < 0.01 || parseFloat(newE) > 100) return
  if (newC !== oldC) {
    console.log("c", newC)
    // c.value = removeZero(parseFloat(newC))
    e.value = removeZero(rollOver.value ? (99.99 - parseFloat(newC)) : parseFloat(newC))
    d.value = removeZero((95 / parseFloat(e.value)))
  } else if (newE !== oldE) {
    console.log("e", newE)
    // e.value = removeZero(parseFloat(newE))
    c.value = removeZero(rollOver.value ? (99.99 - parseFloat(newE)) : parseFloat(newE))
    d.value = removeZero((95 / parseFloat(newE)))
  }
})

watch(a, (newA, oldA) => {
  if (parseFloat(newA) < 0.1) return
  if (newA !== oldA) {
    console.log("a", newA)
    // a.value = removeZero(parseFloat(newA))
    d.value = removeZero((95 / parseFloat(e.value)))
  }
})

watch(d, (newD, oldD) => {
  if (parseFloat(newD) < 1.01) return
  if (newD !== oldD) {
    // d.value = removeZero(parseFloat(newD))
    e.value = removeZero((95 / parseFloat(newD)))
  }
})

const profit = computed(() => {
  return removeZero((parseFloat(a.value) * (95 / parseFloat(e.value)) - parseFloat(a.value)))
})

function toggleRollOver() {
  rollOver.value = !rollOver.value
  c.value = removeZero(99.99 - parseFloat(c.value))
}
</script>

<template>
  <div className="card max-w-xl bg-base-100 shadow-xl">
    <div className="card-body grid grid-cols-3 gap-4">
      <!-- Coins -->
      <div className="col-span-2">
        <label className="form-control w-full">
          <div className="pl-0 label">
            <span className="label-text">BET AMOUNT</span>
          </div>
          <input type="number" v-model.number="a" className="input input-bordered input-md w-full" />
        </label>
      </div>
      <!-- Bet Amount -->
      <div className="col-span-1">
        <div className="pl-0 label">
          <span className="label-text">PROFIT ON WIN</span>
        </div>
        <button className="w-full btn btn-outline btn-disabled no-animation btn-md">{{ parseFloat(profit).toFixed(2)
          }}</button>
      </div>
      <!-- Roll Under -->
      <div className="col-span-1">
        <label className="form-control w-full">
          <div className="pl-0 label">
            <span className="label-text">ROLL {{ rollOver ? "OVER" : "UNDER" }}</span>
          </div>
          <div className="grid grid-cols-3 gap-2">
            <!-- Roll Change Button -->
            <div className="col-span-1">
              <button @click="toggleRollOver" className="w-full btn btn-outline btn-md">
                <svg width="24px" height="24px" stroke-width="1.5" viewBox="0 0 24 24" fill="none"
                  xmlns="http://www.w3.org/2000/svg" stroke="currentColor">
                  <path d="M7 4L7 5" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"
                    stroke-linejoin="round">
                  </path>
                  <path d="M7 9L7 10" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"
                    stroke-linejoin="round"></path>
                  <path d="M17 20V4M17 4L20 7M17 4L14 7" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"
                    stroke-linejoin="round"></path>
                  <path d="M7 14V20M7 20L10 17M7 20L4 17" stroke="currentColor" stroke-width="1.5"
                    stroke-linecap="round" stroke-linejoin="round"></path>
                </svg>
              </button>
            </div>
            <!-- Roll Under Input -->
            <div className="col-span-2">
              <input type="number" v-model.number="c" className="input input-bordered input-md w-full" />
            </div>
          </div>
        </label>
      </div>
      <!-- Multiplier -->
      <div className="col-span-1">
        <label className="form-control w-full">
          <div className="pl-0 label">
            <span className="label-text">MULTIPLIER</span>
          </div>
          <input type="number" v-model.number="d" className="input input-bordered input-md w-full" />
        </label>
      </div>
      <!-- Win Chance -->
      <div className="col-span-1">
        <label className="form-control w-full">
          <div className="pl-0 label">
            <span className="label-text">WIN CHANCE</span>
          </div>
          <input type="number" v-model.number="e" className="input input-bordered input-md w-full" />
        </label>
      </div>
    </div>
  </div>
</template>

<style scoped></style>