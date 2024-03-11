<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { useForm } from '@vorms/core'
import { IconArrowsSort, IconCoins, IconX } from '@tabler/icons-vue'
import { state } from '../main';

const value = ref(new Map<number, number>)
const result = ref(new Map<number, string>)

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
    let response = await fetch("http://localhost:1323/dice", {
      method: 'POST',
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });
    if (response.ok) {
      state.coins -= data.betAmount
      let json = await response.json();
      let JSONvalue = json["value"]
      let JSONnonce = json["nonce"]
      value.value.set(JSONnonce, JSONvalue / 100)
      if (data.roll) {
        if (JSONvalue / 100 > data.rollValue) {
          result.value.set(JSONnonce, "win")
          state.coins = removeZero(removeZero(state.coins) + removeZero((data.betAmount * (95 / data.winChance) - data.betAmount)))
        } else {
          result.value.set(JSONnonce, "lose")
        }
      } else {
        if (JSONvalue / 100 < data.rollValue) {
          result.value.set(JSONnonce, "win")
          state.coins = removeZero(removeZero(state.coins) + removeZero((data.betAmount * (95 / data.winChance) - data.betAmount)))
        } else {
          result.value.set(JSONnonce, "lose")
        }
      }
    } else {
      alert("HTTP-Error: " + response.status);
    }
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
  <div class="grid place-items-center col-span-1">
    <div class="pt-28 pb-20">
      <form @submit="handleSubmit" @reset="handleReset">
        <div class="card max-w-xl bg-neutral shadow-xl">
          <div class="card-body grid grid-cols-3 gap-4">
            <!-- Coins -->
            <div class="col-span-2">
              <div class="pl-0 label">
                <span class="label-text">BET AMOUNT</span>
              </div>
              <label class="input input-bordered flex items-center gap-2">
                <IconCoins :size=20 :stroke-width=1.5 />
                <input type="number" step=0.01 v-model.number="a" class="grow w-full max-w-xs" inputmode="decimal" />
              </label>
              <div className="label">
                <span className="label-text-alt"
                  style="color: var(--fallback-er,oklch(var(--er)/var(--tw-bg-opacity)));">
                  {{ errors.betAmount }}
                </span>
              </div>
            </div>
            <!-- Bet Amount -->
            <div class="col-span-1">
              <div class="pl-0 label">
                <span class="label-text">PROFIT ON WIN</span>
              </div>
              <label class="input input-bordered input-success flex items-center gap-2">
                <IconCoins :size=20 :stroke-width=1.5 />
                <input type="text" v-model="profit" disabled className="grow w-full max-w-xs" />
              </label>
            </div>
            <!-- Roll Under -->
            <div class="col-span-1">
              <div class="pl-0 label">
                <span class="label-text">
                  ROLL {{ rollOver ? "OVER" : "UNDER" }}
                </span>
              </div>
              <div class="grid grid-cols-3 gap-2">
                <!-- Roll Change Button -->
                <div class="col-span-1">
                  <button @click.prevent="toggleRollOver" class="w-full btn btn-square">
                    <IconArrowsSort :size=24 :stroke-width=1.5 />
                  </button>
                </div>
                <!-- Roll Under Input -->
                <div class="col-span-2">
                  <label class="input input-bordered flex items-center gap-2">
                    <input type="number" step=0.0001 v-model.number="c" class="grow w-full max-w-xs"
                      inputmode="decimal" />
                  </label>
                </div>
              </div>
              <div className="label">
                <span className="label-text-alt"
                  style="color: var(--fallback-er,oklch(var(--er)/var(--tw-bg-opacity)));">
                  {{ errors.rollValue }}
                </span>
              </div>
            </div>
            <!-- Multiplier -->
            <div class="col-span-1">
              <div class="pl-0 label">
                <span class="label-text">MULTIPLIER</span>
              </div>
              <label class="input input-bordered flex items-center gap-2">
                <IconX :size=20 :stroke-width=1.5 />
                <input type="number" step=0.0001 v-model.number="d" class="grow w-full max-w-xs" inputmode="decimal" />
              </label>
              <div className="label">
                <span className="label-text-alt"
                  style="color: var(--fallback-er,oklch(var(--er)/var(--tw-bg-opacity)));">
                  {{ errors.multiplier }}
                </span>
              </div>
            </div>
            <!-- Win Chance -->
            <div class="col-span-1">
              <div class="pl-0 label">
                <span class="label-text">WIN CHANCE</span>
              </div>
              <label class="input input-bordered flex items-center gap-2">
                <input type="number" step=0.0001 v-model.number="e" class="grow w-full max-w-xs" inputmode="decimal" />
              </label>
              <div className="label">
                <span className="label-text-alt"
                  style="color: var(--fallback-er,oklch(var(--er)/var(--tw-bg-opacity)));">
                  {{ errors.winChance }}
                </span>
              </div>
            </div>
            <!-- Roll Dice -->
            <div class="col-span-3 pt-4">
              <button type="submit" class="btn btn-success w-full">ROLL DICE</button>
            </div>
          </div>
        </div>
      </form>
    </div>
    <div class="max-w-lg">
      <table class="table">
        <thead>
          <tr>
            <th>Nonce</th>
            <!-- <th>ServerSeed</th>
            <th>ClientSeed</th>  -->
            <th>Value</th>
            <th>Result</th>
            <th>Net</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(v, k) in Array.from(value).reverse()" :key="k">
            <th>{{ v[0] }}</th>
            <td>{{ v[1] }}</td>
            <td>{{ result.get(v[0]) }}</td>
            <td>{{ result.get(v[0]) === "win" ? profit : 0 }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style scoped></style>