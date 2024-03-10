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

watch([c, e], ([newC, newE], [oldC, oldE]) => {
  if (parseFloat(newC) < 0.01 || parseFloat(newE) < 0.01) return
  if (newC !== oldC) {
    console.log("c", newC)
    e.value = rollOver.value ? (99.99 - parseFloat(newC)).toFixed(2).toString() : parseFloat(newC).toFixed(2).toString()
    d.value = (95 / parseFloat(e.value)).toFixed(4).toString()
  } else if (newE !== oldE) {
    console.log("e", newE)
    c.value = rollOver.value ? (99.99 - parseFloat(newE)).toFixed(2).toString() : parseFloat(newE).toFixed(2).toString()
    d.value = (95 / parseFloat(newE)).toFixed(4).toString()
  }
})

watch(a, (newA, oldA) => {
  if (parseFloat(newA) < 0.1) return
  if (newA !== oldA) {
    console.log("a", newA)
    d.value = (95 / parseFloat(e.value)).toFixed(4)
  }
})

watch(d, (newD, oldD) => {
  if (parseFloat(newD) < 1.01) return
  if (newD !== oldD) {
    console.log("d", newD)
    e.value = (95 / parseFloat(newD)).toFixed(2).toString()
  }
})

const profit = computed(() => {
  return (parseFloat(a.value) * (95 / parseFloat(e.value)) - parseFloat(a.value)).toFixed(2).toString()
})

function toggleRollOver() {
  rollOver.value = !rollOver.value
  c.value = (99.99 - parseFloat(c.value)).toFixed(2).toString()
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