<script setup lang="ts">
import { computed, watch } from 'vue'
import { useForm } from '@vorms/core'

const { errors, register, handleSubmit, handleReset } = useForm({
  initialValues: {
    betAmount: 1,
    rollValue: 10,
    multiplier: 9.5,
    winChance: 10,
    roll: false
  },
  validate(data) {
    const errors: Record<string, any> = {}
    // if roll is under, rollValue should be between 0.01 and 94
    // if roll is over, rollValue should be between 5.99 and 99.98
    // if (!data.roll) {
    //   if (data.rollValue <= 0.01 || data.rollValue >= 94) {
    //     errors.rollValue = "Roll value should be between 0.01 and 94"
    //   }
    // } else {
    //   if (data.rollValue <= 5.99 || data.rollValue >= 99.98) {
    //     errors.rollValue = "Roll value should be between 5.99 and 99.98"
    //   }
    // }
    // // betAmount should be greater than 0.1
    // if (data.betAmount < 0.1) {
    //   errors.betAmount = "Bet amount should be greater than 0.1"
    // }
    if (!data.betAmount) {
      errors.betAmount = "Bet amount is required"
    }
    if (!data.rollValue) {
      errors.rollValue = "Roll value is required"
    }
    if (!data.multiplier) {
      errors.multiplier = "Multiplier is required"
    }
    if (!data.winChance) {
      errors.winChance = "Win chance is required"
    }

    return errors
  },
  async onSubmit(data) {
    console.log(JSON.stringify(data))
  }
})

const { value: a } = register('betAmount', {
  validate(value) {
    if (value < 0.1) {
      return "Bet amount should be greater than 0.1"
    }
  }
})
const { value: c } = register('rollValue', {
  validate(value) {
    if (!rollOver.value) {
      if (value <= 0.01 || value >= 94) {
        return "Roll value should be between 0.01 and 94"
      }
    } else {
      if (value <= 5.99 || value >= 99.98) {
        return "Roll value should be between 5.99 and 99.98"
      }
    }
  }
})
const { value: d } = register('multiplier', {
  validate(value) {
    if (value <= 0.01 || value >= 94) {
      return "Multiplier should be between 0.01 and 94"
    }
  }
})
const { value: e } = register('winChance')
const { value: rollOver } = register('roll')

// const rollOver = ref(false)

// a = betAmount
// b = profitOnWin
// c = rollUnder
// d = multiplier
// e = winChance

function removeZero(num: number, precision: number = 4): number {
  return (+(num).toFixed(precision) / 1)
}

watch([c, e], ([newC, newE], [oldC, oldE]) => {
  if (!rollOver.value) {
    if (newC <= 0.01 || newC >= 94) return
    if (newE <= 0.01 || newE >= 94) return
  } else {
    if (99.99 - newC <= 5.99 || 99.99 - newC >= 99.98) return
    if (newE <= 5.99 || newE >= 99.98) return
  }
  if (newC !== oldC) {
    e.value = removeZero(rollOver.value ? (99.99 - newC) : newC)
    d.value = removeZero((95 / e.value))
  } else if (newE !== oldE) {
    c.value = removeZero(rollOver.value ? (99.99 - newE) : newE)
    d.value = removeZero((95 / newE))
  }
})


watch(a, (newA, oldA) => {
  if (newA < 0.1) return
  if (newA !== oldA) {
    d.value = removeZero((95 / e.value))
  }
})

watch(d, (newD, oldD) => {
  if (newD < 1.01) return
  if (newD !== oldD) {
    e.value = removeZero((95 / newD))
  }
})

const profit = computed(() => {
  let USD = new Intl.NumberFormat('en-US', {})
  return USD.format(removeZero((a.value * (95 / e.value) - a.value)))
})

function toggleRollOver() {
  rollOver.value = !rollOver.value
  c.value = removeZero(99.99 - c.value)
}
</script>

<template>
  <form @submit="handleSubmit" @reset="handleReset">
    <div class="card max-w-xl bg-neutral shadow-xl">
      <div class="card-body grid grid-cols-3 gap-4">
        <!-- Coins -->
        <div class="col-span-2">
          <div class="pl-0 label">
            <span class="label-text">BET AMOUNT</span>
          </div>
          <input type="number" step=0.01 v-model.number="a" class="input input-bordered w-full max-w-xs" />
          <div className="label">
            <span className="label-text-alt" style="color: var(--fallback-er,oklch(var(--er)/var(--tw-bg-opacity)));">{{
    errors.betAmount }}</span>
          </div>
        </div>
        <!-- Bet Amount -->
        <div class="col-span-1">
          <div class="pl-0 label">
            <span class="label-text">PROFIT ON WIN</span>
          </div>
          <button class="w-full max-w-xs btn btn-outline btn-disabled no-animation text-sm">{{
    profit
  }}</button>
        </div>
        <!-- Roll Under -->
        <div class="col-span-1">
          <div class="pl-0 label">
            <span class="label-text">ROLL {{ rollOver ? "OVER" : "UNDER" }}</span>
          </div>
          <div class="grid grid-cols-3 gap-2">
            <!-- Roll Change Button -->
            <div class="col-span-1">
              <button @click.prevent="toggleRollOver" class="w-full btn btn-square">
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
              <input type="number" step=0.0001 v-model.number="c" class="input input-bordered w-full max-w-xs" />
            </div>
          </div>
          <div className="label">
            <span className="label-text-alt" style="color: var(--fallback-er,oklch(var(--er)/var(--tw-bg-opacity)));">{{
    errors.rollValue }}</span>
          </div>
        </div>
        <!-- Multiplier -->
        <div class="col-span-1">
          <div class="pl-0 label">
            <span class="label-text">MULTIPLIER</span>
          </div>
          <input type="number" step=0.0001 v-model.number="d" class="input input-bordered w-full max-w-xs" />
          <div className="label">
            <span className="label-text-alt" style="color: var(--fallback-er,oklch(var(--er)/var(--tw-bg-opacity)));">{{
    errors.multiplier }}</span>
          </div>
        </div>
        <!-- Win Chance -->
        <div class="col-span-1">
          <div class="pl-0 label">
            <span class="label-text">WIN CHANCE</span>
          </div>
          <input type="number" step=0.0001 v-model.number="e" class="input input-bordered w-full max-w-xs" />
          <div className="label">
            <span className="label-text-alt" style="color: var(--fallback-er,oklch(var(--er)/var(--tw-bg-opacity)));">{{
              errors.winChance }}</span>
          </div>
        </div>
        <!-- Roll Dice -->
        <div class="col-span-3 pt-4">
          <button type="submit" class="btn btn-success w-full">Roll Dice</button>
        </div>
      </div>
    </div>
  </form>
</template>

<style scoped></style>