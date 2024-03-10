<script setup lang="ts">
</script>

<template>
  <div>
    <div class="navbar bg-base-100">
      <div class="navbar-start">
        <div class="dropdown">
          <div tabindex="0" role="button" class="btn btn-ghost lg:hidden">
            <svg width="24px" height="24px" stroke-width="1.5" viewBox="0 0 24 24" fill="none"
              xmlns="http://www.w3.org/2000/svg" stroke="currentColor">
              <path d="M3 5H21" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
              </path>
              <path d="M3 12H21" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"
                stroke-linejoin="round">
              </path>
              <path d="M3 19H21" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"
                stroke-linejoin="round">
              </path>
            </svg>
          </div>
          <ul tabindex="0" class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52">
            <li><router-link to="/crash">Crash</router-link></li>
            <li><router-link to="/dice">Dice</router-link></li>
            <li><router-link to="/roll">Roll</router-link></li>
            <li><router-link to="/plinko">Plinko</router-link></li>
          </ul>
        </div>
        <a class="btn btn-ghost text-xl">Hobby</a>
      </div>
      <div class="navbar-center hidden lg:flex">
        <ul class="menu menu-horizontal px-1">
          <li><router-link to="/crash">Crash</router-link></li>
          <li><router-link to="/dice">Dice</router-link></li>
          <li><router-link to="/roll">Roll</router-link></li>
          <li><router-link to="/plinko">Plinko</router-link></li>
        </ul>
      </div>
      <div class="navbar-end">
        <a class="btn">Button</a>
      </div>
    </div>
    <div class="pointer-events-none fixed inset-0 flex items-center justify-center">
      <router-view class="pointer-events-auto"></router-view>
    </div>
  </div>
</template>
