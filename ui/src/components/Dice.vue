<script setup lang="ts">
import { CSSProperties, computed, ref, watch } from 'vue'

const a = ref("1")
const c = ref("10")
const d = ref("9.5")
const e = ref("10")
const f = ref("10")

const rollOver = ref(false)
let directionRange: CSSProperties = {
  direction: "ltr"
}

// a = betAmount
// b = profitOnWin
// c = rollUnder
// d = multiplier
// e = winChance

// c = e
// d = a * (95 / c) - a

function removeZero(num: number, precision: number = 4): string {
  return (+(num).toFixed(precision) / 1).toString()
}

watch([c, e], ([newC, newE], [oldC, oldE]) => {
  // if (parseFloat(newC) < 0.01 || parseFloat(newC) > 100 || parseFloat(newE) < 0.01 || parseFloat(newE) > 100 || parseFloat(newF) < 0.01 || parseFloat(newF) > 100) return
  // if (newC !== oldC) {
  //   // c.value = removeZero(parseFloat(newC))
  //   e.value = removeZero(rollOver.value ? (99.99 - parseFloat(newC)) : parseFloat(newC))
  //   f.value = removeZero(parseFloat(newC))
  //   d.value = removeZero((95 / parseFloat(e.value)))
  //   console.log("c", newC, oldC)
  //   console.log("c", c.value, d.value, e.value, f.value)
  // } else if (newE !== oldE) {
  //   // e.value = removeZero(parseFloat(newE))
  //   c.value = removeZero(rollOver.value ? (99.99 - parseFloat(newE)) : parseFloat(newE))
  //   f.value = removeZero(parseFloat(newE))
  //   d.value = removeZero((95 / parseFloat(newE)))
  //   console.log("e", newE, oldE)
  //   console.log("e", c.value, d.value, e.value, f.value)
  // } else if (newF !== oldF) {
  //   // f.value = removeZero(parseFloat(newF))
  //   c.value = removeZero(rollOver.value ? (99.99 - parseFloat(newF)) : parseFloat(newF))
  //   e.value = removeZero(parseFloat(newF))
  //   d.value = removeZero((95 / parseFloat(e.value)))
  //   console.log("f", newF, oldF)
  //   console.log("f", c.value, d.value, e.value, f.value)
  // }

  if (parseFloat(newC) < 0.01 || parseFloat(newC) > 100 || parseFloat(newE) < 0.01 || parseFloat(newE) > 100) return
  if (newC !== oldC) {
    console.log("c", newC)
    e.value = removeZero(rollOver.value ? (99.99 - parseFloat(newC)) : parseFloat(newC))
    d.value = removeZero((95 / parseFloat(e.value)), 4)
  } else if (newE !== oldE) {
    console.log("e", newE)
    c.value = removeZero(rollOver.value ? (99.99 - parseFloat(newE)) : parseFloat(newE))
    d.value = removeZero((95 / parseFloat(newE)), 4)
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
    console.log("d", newD)
    // d.value = removeZero(parseFloat(newD))
    e.value = removeZero((95 / parseFloat(newD)))
  }
})

// watch(f, (newF, oldF) => {
//   if (parseFloat(newF) < 0.01 || parseFloat(newF) > 100) return
//   if (newF !== oldF) {
//     console.log("f", newF)
//     // f.value = removeZero(parseFloat(newF))
//     e.value = newF
//   }
// })

const profit = computed(() => {
  return removeZero((parseFloat(a.value) * (95 / parseFloat(e.value)) - parseFloat(a.value)))
})

function toggleRollOver() {
  rollOver.value = !rollOver.value
  console.log(rollOver.value ? "rtl" : "ltr")
  directionRange = {
    direction: rollOver.value ? "rtl" : "ltr"
  }
  c.value = removeZero(99.99 - parseFloat(c.value))
}

function validateValue() {
  d.value = removeZero((95 / parseFloat(e.value)))
}
</script>

<template>
  <div class="card max-w-xl bg-neutral shadow-xl">
    <div class="card-body grid grid-cols-3 gap-4">
      <!-- Coins -->
      <div class="col-span-2">
        <label class="form-control w-full">
          <div class="pl-0 label">
            <span class="label-text">BET AMOUNT</span>
          </div>
          <input type="number" @blur="validateValue" v-model.number="a" class="input input-bordered input-md w-full" />
        </label>
      </div>
      <!-- Bet Amount -->
      <div class="col-span-1">
        <div class="pl-0 label">
          <span class="label-text">PROFIT ON WIN</span>
        </div>
        <button class="w-full btn btn-outline btn-disabled no-animation btn-md">{{
            parseFloat(profit).toFixed(2).toString()
          }}</button>
      </div>
      <!-- Roll Under -->
      <div class="col-span-1">
        <label class="form-control w-full">
          <div class="pl-0 label">
            <span class="label-text">ROLL {{ rollOver ? "OVER" : "UNDER" }}</span>
          </div>
          <div class="grid grid-cols-3 gap-2">
            <!-- Roll Change Button -->
            <div class="col-span-1">
              <button @click="toggleRollOver" class="w-full btn btn-outline btn-md">
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
            <div class="col-span-2">
              <input type="number" @blur="validateValue" v-model.number="c"
                class="input input-bordered input-md w-full" />
            </div>
          </div>
        </label>
      </div>
      <!-- Multiplier -->
      <div class="col-span-1">
        <label class="form-control w-full">
          <div class="pl-0 label">
            <span class="label-text">MULTIPLIER</span>
          </div>
          <input type="number" @blur="validateValue" v-model.number="d" class="input input-bordered input-md w-full" />
        </label>
      </div>
      <!-- Win Chance -->
      <div class="col-span-1">
        <label class="form-control w-full">
          <div class="pl-0 label">
            <span class="label-text">WIN CHANCE</span>
          </div>
          <input type="number" @blur="validateValue" v-model.number="e" class="input input-bordered input-md w-full" />
        </label>
      </div>
      <!-- Range -->
      <div class="pt-4 col-span-3">
        <input type="range" min="0" max="100" v-model="f" class="range" :style="directionRange" step="0.01" />
        <div class="w-full flex justify-between text-xs pt-2">
          <span>0</span>
          <span>100</span>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>